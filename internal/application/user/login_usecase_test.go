package user_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	userApplication "github.com/pauloRohling/locknote/internal/application/user"
	"github.com/pauloRohling/locknote/internal/domain/types/email"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/locknote/internal/domain/types/password"
	"github.com/pauloRohling/locknote/internal/domain/user"
	mocktoken "github.com/pauloRohling/locknote/internal/mocks/token"
	mockuser "github.com/pauloRohling/locknote/internal/mocks/user"
	"github.com/pauloRohling/locknote/pkg/testinstance"
	"github.com/pauloRohling/throw"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestLoginUsecase(t *testing.T) {
	userId, _ := id.NewID()
	userEmail, _ := email.NewEmail("test@test.com")
	userPassword, _ := password.New("test123456")

	mockUser := testinstance.Set(&user.User{}, map[string]any{
		"id":       userId,
		"email":    userEmail,
		"password": userPassword,
	})

	testCases := map[string]struct {
		input     *userApplication.LoginInput
		setup     func(userRepository *mockuser.MockRepository, tokenIssuer *mocktoken.MockIssuer)
		expectErr bool
		errType   throw.ErrorType
	}{
		"should login a user": {
			input: &userApplication.LoginInput{
				Email:    "test@test.com",
				Password: "test123456",
			},
			setup: func(userRepository *mockuser.MockRepository, tokenIssuer *mocktoken.MockIssuer) {
				userRepository.EXPECT().FindByEmail(mock.Anything, mock.Anything).Return(mockUser, nil)
				tokenIssuer.EXPECT().Issue(mock.Anything).Return("token", id.ID(uuid.New()), nil)
			},
		},
		"should return an error if the user repository fails to find the user": {
			input: &userApplication.LoginInput{
				Email:    "test@test.com",
				Password: "test123456",
			},
			setup: func(userRepository *mockuser.MockRepository, tokenIssuer *mocktoken.MockIssuer) {
				userRepository.EXPECT().FindByEmail(mock.Anything, mock.Anything).Return(nil, errors.New("user not found"))
			},
			expectErr: true,
			errType:   throw.UnauthorizedErrorType,
		},
		"should return an error if the token issuer fails to issue the token": {
			input: &userApplication.LoginInput{
				Email:    "test@test.com",
				Password: "test123456",
			},
			setup: func(userRepository *mockuser.MockRepository, tokenIssuer *mocktoken.MockIssuer) {
				userRepository.EXPECT().FindByEmail(mock.Anything, mock.Anything).Return(mockUser, nil)
				tokenIssuer.EXPECT().Issue(mock.Anything).Return("", id.ID(uuid.New()), errors.New("failed to issue token"))
			},
			expectErr: true,
			errType:   throw.UnauthorizedErrorType,
		},
		"should return an error if the user email is invalid": {
			input: &userApplication.LoginInput{
				Email:    "test",
				Password: "test123456",
			},
			setup: func(userRepository *mockuser.MockRepository, tokenIssuer *mocktoken.MockIssuer) {
			},
			expectErr: true,
			errType:   throw.UnauthorizedErrorType,
		},
		"should return and error if the user password does not match": {
			input: &userApplication.LoginInput{
				Email:    "test@test.com",
				Password: "some-other-password",
			},
			setup: func(userRepository *mockuser.MockRepository, tokenIssuer *mocktoken.MockIssuer) {
				userRepository.EXPECT().FindByEmail(mock.Anything, mock.Anything).Return(mockUser, nil)
			},
			expectErr: true,
			errType:   throw.UnauthorizedErrorType,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			userRepository := mockuser.NewMockRepository(t)
			tokenIssuer := mocktoken.NewMockIssuer(t)

			testCase.setup(userRepository, tokenIssuer)

			loginUseCase := userApplication.NewLoginUseCase(userApplication.LoginUsecaseParams{
				UserRepository: userRepository,
				TokenIssuer:    tokenIssuer,
			})

			output, err := loginUseCase.Execute(context.Background(), testCase.input)

			if testCase.expectErr {
				assert.Error(t, err)
				throw.AssertType(t, err, testCase.errType.String())
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, output)
			assert.Equal(t, output.AccessToken, "token")
		})
	}
}
