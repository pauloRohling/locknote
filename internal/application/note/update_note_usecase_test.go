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

func TestUpdateNoteUseCase(t *testing.T) {
	userId, _ := id.NewID()
	noteId, _ := id.NewID()
	ctx := context.WithValue(context.Background(), audit.UserIdContextKey, userId)

	mockNote := testinstance.Set(&note.Note{}, map[string]any{
		"id":      noteId,
		"title":   text.Title("Test Note"),
		"content": "Test Note Content",
	})

	testCases := map[string]struct {
		input     *noteApplication.UpdateNoteInput
		setup     func(noteRepository *mocknote.MockRepository)
		expectErr bool
		errType   throw.ErrorType
	}{
		"should update a note": {
			input: &noteApplication.UpdateNoteInput{
				ID:      noteId.String(),
				Title:   "Test Note Updated",
				Content: "Test Note Content Updated",
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.EXPECT().FindByID(ctx, mock.Anything).Return(mockNote, nil)
				noteRepository.EXPECT().UpdateById(ctx, mock.Anything).Return(mockNote, nil)
			},
		},
		"should return an error if the note repository fails to find the note": {
			input: &noteApplication.UpdateNoteInput{
				ID:      noteId.String(),
				Title:   "Test Note Updated",
				Content: "Test Note Content Updated",
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.EXPECT().FindByID(ctx, mock.Anything).Return(nil, throw.Internal().Msg("failed to find note"))
				noteRepository.AssertNotCalled(t, "UpdateById")
			},
			expectErr: true,
			errType:   throw.InternalErrorType,
		},
		"should return an error if the note repository fails to update the note": {
			input: &noteApplication.UpdateNoteInput{
				ID:      noteId.String(),
				Title:   "Test Note Updated",
				Content: "Test Note Content Updated",
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.EXPECT().FindByID(ctx, mock.Anything).Return(mockNote, nil)
				noteRepository.EXPECT().UpdateById(ctx, mock.Anything).Return(nil, throw.Internal().Msg("failed to update note"))
			},
			expectErr: true,
			errType:   throw.InternalErrorType,
		},
		"should return an error if the id is invalid": {
			input: &noteApplication.UpdateNoteInput{
				ID:      "",
				Title:   "Test Note Updated",
				Content: "Test Note Content Updated",
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.AssertNotCalled(t, "FindByID")
				noteRepository.AssertNotCalled(t, "UpdateById")
			},
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
		"should return an error if the input is invalid": {
			input: &noteApplication.UpdateNoteInput{
				ID:      noteId.String(),
				Title:   "",
				Content: "",
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.EXPECT().FindByID(ctx, mock.Anything).Return(mockNote, nil)
				noteRepository.AssertNotCalled(t, "UpdateById")
			},
			expectErr: true,
			errType:   throw.ValidationErrorType,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			noteFactory := note.NewFactory()
			noteRepository := mocknote.NewMockRepository(t)

			if testCase.setup != nil {
				testCase.setup(noteRepository)
			}

			createNoteUseCase := noteApplication.NewUpdateNoteUseCase(noteApplication.UpdateNoteParams{
				NoteFactory:    noteFactory,
				NoteRepository: noteRepository,
			})

			newNote, err := createNoteUseCase.Execute(ctx, testCase.input)

			if testCase.expectErr {
				assert.Error(t, err)
				throw.AssertType(t, err, testCase.errType.String())
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, newNote)
		})
	}
}
