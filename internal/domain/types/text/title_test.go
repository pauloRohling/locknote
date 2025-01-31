package text_test

import (
	"github.com/pauloRohling/locknote/internal/domain/types/text"
	"github.com/pauloRohling/throw"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResourceName(t *testing.T) {
	testCases := map[string]struct {
		value     string
		expected  text.Title
		expectErr bool
		errType   throw.ErrorType
	}{
		"should create a valid title": {
			value:    "valid title",
			expected: text.Title("valid title"),
		},
		"should trim spaces from the beginning and end of the title": {
			value:    "    value with spaces    ",
			expected: text.Title("value with spaces"),
		},
		"should not create an empty title": {
			value:     "",
			expected:  "",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
		"should not create a title with more than 255 characters": {
			value:     strings.Repeat("a", 256),
			expected:  "",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
		"should not create a title with spaces": {
			value:     "   ",
			expected:  "",
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			result, err := text.NewTitle(testCase.value)
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
