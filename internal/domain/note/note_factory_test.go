package note_test

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/note"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/throw"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ExpectedNote struct {
	Title   string
	Content string
}

func TestFactory(t *testing.T) {
	userId, _ := id.NewID()
	ctx := context.WithValue(context.Background(), audit.UserIdContextKey, userId)

	testCases := map[string]struct {
		params    note.NewParams
		context   context.Context
		expected  ExpectedNote
		expectErr bool
		errType   throw.ErrorType
	}{
		"should create a new note": {
			context: ctx,
			params: note.NewParams{
				Title:   "Test Note",
				Content: "Test Note Content",
			},
			expected: ExpectedNote{
				Title:   "Test Note",
				Content: "Test Note Content",
			},
		},
		"should return an error if the note title is invalid": {
			context: ctx,
			params: note.NewParams{
				Title:   "",
				Content: "Test Note Content",
			},
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
		"should return an error if the note content is invalid": {
			context: ctx,
			params: note.NewParams{
				Title:   "Test Note",
				Content: "",
			},
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
		"should return an error if the user id cannot be retrieved from the context": {
			context: context.Background(),
			params: note.NewParams{
				Title:   "Test Note",
				Content: "Test Note Content",
			},
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			factory := note.NewFactory()
			newNote, err := factory.New(testCase.context, testCase.params)

			if testCase.expectErr {
				assert.Error(t, err)
				throw.AssertType(t, err, testCase.errType.String())
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, newNote.ID())
			assert.Equal(t, newNote.Title().String(), testCase.expected.Title)
			assert.Equal(t, newNote.Content(), testCase.expected.Content)
			assert.NotEmpty(t, newNote.Audit().CreatedAt())
			assert.Equal(t, newNote.Audit().CreatedBy().String(), userId.String())
		})
	}
}
