package email_test

import (
	"github.com/pauloRohling/locknote/internal/domain/types/email"
	"github.com/pauloRohling/throw"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewEmail(t *testing.T) {
	testcases := map[string]struct {
		email     string
		expected  email.Email
		expectErr bool
		errType   throw.ErrorType
	}{
		"should create a valid email": {
			email:    "valid@email.com",
			expected: email.Email("valid@email.com"),
		},
		"should not create an email with at sign in the end": {
			email:     "email.com@",
			expected:  "",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
		"should not create an email with at sign in the beginning": {
			email:     "@email.com",
			expected:  "",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
		"should not create an email with more than one at sign": {
			email:     "valid@email.com@valid.com",
			expected:  "",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
		"should not create an empty email": {
			email:     "",
			expected:  "",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
		"should not create an email with more than 255 characters": {
			email:     strings.Repeat("a", 256),
			expected:  "",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
		"should not create an email with spaces": {
			email:     "   ",
			expected:  "",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
	}

	for testName, testCase := range testcases {
		t.Run(testName, func(t *testing.T) {
			result, err := email.NewEmail(testCase.email)
			if testCase.expectErr {
				assert.Error(t, err)
				throw.AssertType(t, err, testCase.errType.String())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
