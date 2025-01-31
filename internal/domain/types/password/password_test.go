package password_test

import (
	"github.com/pauloRohling/locknote/internal/domain/types/password"
	"github.com/pauloRohling/throw"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	testCases := map[string]struct {
		password  string
		expectErr bool
		errType   throw.ErrorType
	}{
		"should create a valid password": {
			password: "password123",
		},
		"should not create a password with less than 8 characters": {
			password:  "pass",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
		"should not create a password with more than 70 characters": {
			password:  "password1234567890123456789012345678901234567890123456789012345678901234567890123456789012",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			_, err := password.New(testCase.password)
			if testCase.expectErr {
				assert.Error(t, err)
				throw.AssertType(t, err, testCase.errType.String())
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestFromEncrypted(t *testing.T) {
	testCases := map[string]struct {
		password  string
		expectErr bool
		errType   throw.ErrorType
	}{
		"should create a valid password": {
			password: "$2a$12$o4UZxjtTV/r9kN/d2dwupu04EPGmuQ.GZ9zEnNu0euRsjhWEFrGO.",
		},
		"should not create a password with a non-encrypted string": {
			password:  "teste",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
		"should not create a password with hash less than 56 characters": {
			password:  "$2a$10$teste",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			result, err := password.FromEncrypted(testCase.password)
			if testCase.expectErr {
				assert.Error(t, err)
				throw.AssertType(t, err, testCase.errType.String())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, testCase.password, result.String())
		})
	}
}
