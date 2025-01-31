package id_test

import (
	"github.com/google/uuid"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/throw"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewID(t *testing.T) {
	uuidV7, err := id.NewID()
	assert.NoError(t, err)

	parsedUuid, err := uuid.Parse(uuidV7.String())

	assert.NoError(t, err)
	assert.NotEqual(t, id.Nil, uuidV7)
	assert.EqualValues(t, parsedUuid.Version(), 7)
}

func TestFromString(t *testing.T) {
	uuidV7, _ := id.NewID()

	testCases := map[string]struct {
		id                string
		expected          id.ID
		expectedErrorType throw.ErrorType
	}{
		"should create an ID from string": {
			id:       uuidV7.String(),
			expected: uuidV7,
		},
		"should not create an ID from V4 UUID": {
			id:                uuid.New().String(),
			expected:          id.Nil,
			expectedErrorType: throw.ValidationErrorType,
		},
		"should not create an ID from empty string": {
			id:                "",
			expected:          id.Nil,
			expectedErrorType: throw.ValidationErrorType,
		},
		"should not create an ID from nil UUID": {
			id:                uuid.Nil.String(),
			expected:          id.Nil,
			expectedErrorType: throw.ValidationErrorType,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			result, err := id.FromString(testCase.id)
			throw.AssertType(t, err, testCase.expectedErrorType.String())
			assert.Equal(t, testCase.expected, result)
		})
	}
}
