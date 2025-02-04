package note_test

import (
	"context"
	noteApplication "github.com/pauloRohling/locknote/internal/application/note"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/note"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/locknote/internal/domain/types/text"
	mocknote "github.com/pauloRohling/locknote/internal/mocks/note"
	"github.com/pauloRohling/locknote/pkg/testinstance"
	"github.com/pauloRohling/throw"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetNoteUseCase(t *testing.T) {
	userId, _ := id.NewID()
	noteId, _ := id.NewID()
	ctx := context.WithValue(context.Background(), audit.UserIdContextKey, userId)

	mockNote := testinstance.Set(&note.Note{}, map[string]any{
		"id":      noteId,
		"title":   text.Title("Test Note"),
		"content": "Test Note Content",
	})

	testCases := map[string]struct {
		input     *noteApplication.GetNoteInput
		setup     func(noteRepository *mocknote.MockRepository)
		expectErr bool
		errType   throw.ErrorType
	}{
		"should find a note": {
			input: &noteApplication.GetNoteInput{
				NoteID: mockNote.ID().String(),
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.EXPECT().FindByID(mock.Anything, mockNote.ID()).Return(mockNote, nil)
			},
		},
		"should return an error if the note repository fails": {
			input: &noteApplication.GetNoteInput{
				NoteID: mockNote.ID().String(),
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.EXPECT().FindByID(mock.Anything, mockNote.ID()).Return(nil, throw.Internal().Msg("failed to find note"))
			},
			expectErr: true,
			errType:   throw.InternalErrorType,
		},
		"should return an error if the note id is invalid": {
			input: &noteApplication.GetNoteInput{
				NoteID: "invalid-note-id",
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.AssertNotCalled(t, "FindByID")
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

			getNoteUseCase := noteApplication.NewGetNoteUseCase(noteApplication.GetNoteParams{
				NoteRepository: noteRepository,
			})

			matchedNote, err := getNoteUseCase.Execute(ctx, testCase.input)

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
