package user_test

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/user"
	"github.com/pauloRohling/throw"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ExpectedUser struct {
	Name  string
	Email string
}

func TestFactory(t *testing.T) {
	testCases := map[string]struct {
		params    user.NewParams
		expected  ExpectedUser
		expectErr bool
		errType   throw.ErrorType
	}{
		"should create a new user": {
			params: user.NewParams{
				Name:     "Test User",
				Email:    "test@user.com",
				Password: "test123456",
			},
			expected: ExpectedUser{
				Name:  "Test User",
				Email: "test@user.com",
			},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			factory := user.NewFactory()
			newUser, err := factory.New(context.Background(), testCase.params)

			if testCase.expectErr {
				throw.AssertType(t, err, testCase.errType.String())
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, newUser.ID())
			assert.Equal(t, newUser.Name().String(), testCase.expected.Name)
			assert.Equal(t, newUser.Email().String(), testCase.expected.Email)
			assert.NotEmpty(t, newUser.Password())
			assert.NotEmpty(t, newUser.Audit().CreatedAt())
			assert.Equal(t, newUser.Audit().CreatedBy(), newUser.ID())
		})
	}
}
