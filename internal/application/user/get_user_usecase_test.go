package user

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/types/email"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/locknote/internal/domain/user"
	mockuser "github.com/pauloRohling/locknote/internal/mocks/user"
	"github.com/pauloRohling/locknote/pkg/testinstance"
	"github.com/pauloRohling/throw"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserUseCase(t *testing.T) {
	userId, _ := id.NewID()
	userEmail, _ := email.NewEmail("test@test.com")
	ctx := context.WithValue(context.Background(), audit.UserIdContextKey, userId)

	mockUser := testinstance.Set(&user.User{}, map[string]any{
		"id":    userId,
		"email": userEmail,
	})

	testCases := map[string]struct {
		setup     func(userRepository *mockuser.MockRepository)
		expectErr bool
		errType   throw.ErrorType
	}{
		"should get a user": {
			setup: func(userRepository *mockuser.MockRepository) {
				userRepository.EXPECT().Find(ctx).Return(mockUser, nil)
			},
		},
		"should return an error if the user repository fails to find the user": {
			setup: func(userRepository *mockuser.MockRepository) {
				userRepository.EXPECT().Find(ctx).Return(nil, throw.Internal().Msg("failed to find user"))
			},
			expectErr: true,
			errType:   throw.InternalErrorType,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			userRepository := mockuser.NewMockRepository(t)

			if testCase.setup != nil {
				testCase.setup(userRepository)
			}

			getUserUseCase := NewGetUserUseCase(GetUserParams{
				UserRepository: userRepository,
			})

			matchedUser, err := getUserUseCase.Execute(ctx)

			if testCase.expectErr {
				throw.AssertType(t, err, testCase.errType.String())
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, matchedUser)
		})
	}
}
