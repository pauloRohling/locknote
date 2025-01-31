package audit_test

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/throw"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserId(t *testing.T) {
	newID, _ := id.NewID()
	contextWithUserId := context.WithValue(context.Background(), audit.UserIdContextKey, newID)
	contextWithWrongUserId := context.WithValue(context.Background(), audit.UserIdContextKey, "wrong user id")

	testCases := map[string]struct {
		ctx       context.Context
		expected  id.ID
		expectErr bool
		errType   throw.ErrorType
	}{
		"should get user id from context": {
			ctx:      contextWithUserId,
			expected: newID,
		},
		"should return error if user id is not available": {
			ctx:       context.Background(),
			expected:  id.Nil,
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
		"should return error if user id is not a valid UUID": {
			ctx:       contextWithWrongUserId,
			expected:  id.Nil,
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			result, err := audit.GetUserId(testCase.ctx)

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
