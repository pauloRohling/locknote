package note_test

import (
	"context"
	noteApplication "github.com/pauloRohling/locknote/internal/application/note"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	mocknote "github.com/pauloRohling/locknote/internal/mocks/note"
	"github.com/pauloRohling/throw"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDeleteNoteUseCase(t *testing.T) {
	userId, _ := id.NewID()
	noteId, _ := id.NewID()
	ctx := context.WithValue(context.Background(), audit.UserIdContextKey, userId)

	testCases := map[string]struct {
		input     *noteApplication.DeleteNoteInput
		setup     func(noteRepository *mocknote.MockRepository)
		expectErr bool
		errType   throw.ErrorType
	}{
		"should delete a note": {
			input: &noteApplication.DeleteNoteInput{
				NoteID: noteId.String(),
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.EXPECT().DeleteById(mock.Anything, noteId).Return(nil)
			},
		},
		"should return an error if the note repository fails": {
			input: &noteApplication.DeleteNoteInput{
				NoteID: noteId.String(),
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.EXPECT().DeleteById(mock.Anything, noteId).Return(throw.Internal().Msg("failed to find note"))
			},
			expectErr: true,
			errType:   throw.InternalErrorType,
		},
		"should return an error if the note id is invalid": {
			input: &noteApplication.DeleteNoteInput{
				NoteID: "invalid-note-id",
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.AssertNotCalled(t, "DeleteById")
			},
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			noteRepository := mocknote.NewMockRepository(t)

			if testCase.setup != nil {
				testCase.setup(noteRepository)
			}

			DeleteNoteUseCase := noteApplication.NewDeleteNoteUseCase(noteApplication.DeleteNoteParams{
				NoteRepository: noteRepository,
			})

			matchedNote, err := DeleteNoteUseCase.Execute(ctx, testCase.input)

			if testCase.expectErr {
				assert.Error(t, err)
				throw.AssertType(t, err, testCase.errType.String())
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, matchedNote)
		})
	}
}
