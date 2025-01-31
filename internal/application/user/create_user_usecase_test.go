package user_test

import (
	"context"
	userApplication "github.com/pauloRohling/locknote/internal/application/user"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/locknote/internal/domain/user"
	mockuser "github.com/pauloRohling/locknote/internal/mocks/user"
	"github.com/pauloRohling/throw"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateUserUseCase(t *testing.T) {
	userId, _ := id.NewID()
	ctx := context.WithValue(context.Background(), audit.UserIdContextKey, userId)

	testCases := map[string]struct {
		input     *userApplication.CreateUserInput
		setup     func(userRepository *mockuser.MockRepository)
		expectErr bool
		errType   throw.ErrorType
	}{
		"should create a new user": {
			input: &userApplication.CreateUserInput{
				Name:     "Test User",
				Email:    "test@user.com",
				Password: "test123456",
			},
			setup: func(userRepository *mockuser.MockRepository) {
				userRepository.EXPECT().Save(ctx, mock.Anything).Return(&user.User{}, nil)
			},
		},
		"should return an error if the user repository fails": {
			input: &userApplication.CreateUserInput{
				Name:     "Test User",
				Email:    "test@user.com",
				Password: "test123456",
			},
			setup: func(userRepository *mockuser.MockRepository) {
				userRepository.EXPECT().Save(ctx, mock.Anything).Return(nil, throw.Internal().Msg("failed to save user"))
			},
			expectErr: true,
			errType:   throw.InternalErrorType,
		},
		"should return an error if the factory fails": {
			input: &userApplication.CreateUserInput{
				Name:     "",
				Email:    "test@user.com",
				Password: "test123456",
			},
			setup: func(userRepository *mockuser.MockRepository) {
				userRepository.AssertNotCalled(t, "Save")
			},
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			userFactory := user.NewFactory()
			userRepository := mockuser.NewMockRepository(t)

			if testCase.setup != nil {
				testCase.setup(userRepository)
			}

			createUserUseCase := userApplication.NewCreateUserUseCase(userApplication.CreateUserParams{
				UserFactory:    userFactory,
				UserRepository: userRepository,
			})

			newUser, err := createUserUseCase.Execute(ctx, testCase.input)

			if testCase.expectErr {
				assert.Error(t, err)
				throw.AssertType(t, err, testCase.errType.String())
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, newUser)
		})
	}
}
