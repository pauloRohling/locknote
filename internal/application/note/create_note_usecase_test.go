package note_test

import (
	"context"
	noteApplication "github.com/pauloRohling/locknote/internal/application/note"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/note"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	mocknote "github.com/pauloRohling/locknote/internal/mocks/note"
	"github.com/pauloRohling/throw"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateNoteUseCase(t *testing.T) {
	userId, _ := id.NewID()
	ctx := context.WithValue(context.Background(), audit.UserIdContextKey, userId)

	testCases := map[string]struct {
		input     *noteApplication.CreateNoteInput
		setup     func(noteRepository *mocknote.MockRepository)
		expectErr bool
		errType   throw.ErrorType
	}{
		"should create a new note": {
			input: &noteApplication.CreateNoteInput{
				Title:   "Test Note",
				Content: "Test Note Content",
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.EXPECT().Save(ctx, mock.Anything).Return(&note.Note{}, nil)
			},
		},
		"should return an error if the note repository fails": {
			input: &noteApplication.CreateNoteInput{
				Title:   "Test Note",
				Content: "Test Note Content",
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.EXPECT().Save(ctx, mock.Anything).Return(nil, throw.Internal().Msg("failed to save note"))
			},
			expectErr: true,
			errType:   throw.InternalErrorType,
		},
		"should return an error if the factory fails": {
			input: &noteApplication.CreateNoteInput{
				Title:   "",
				Content: "",
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.AssertNotCalled(t, "Save")
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

			createNoteUseCase := noteApplication.NewCreateNoteUseCase(noteApplication.CreateNoteParams{
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
