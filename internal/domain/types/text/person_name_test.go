package text_test

import (
	"github.com/pauloRohling/locknote/internal/domain/types/text"
	"github.com/pauloRohling/throw"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPersonName(t *testing.T) {
	testCases := map[string]struct {
		value     string
		expected  text.PersonName
		expectErr bool
		errType   throw.ErrorType
	}{
		"should create a valid person name": {
			value:    "valid person name",
			expected: text.PersonName("valid person name"),
		},
		"should trim spaces from the beginning and end of the person name": {
			value:    "    value with spaces    ",
			expected: text.PersonName("value with spaces"),
		},
		"should not create an empty person name": {
			value:     "",
			expected:  "",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
		"should not create a person name with more than 50 characters": {
			value:     strings.Repeat("a", 51),
			expected:  "",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
		"should not create a person name with spaces": {
			value:     "   ",
			expected:  "",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			result, err := text.NewPersonName(testCase.value)
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
