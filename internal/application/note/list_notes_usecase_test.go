package note

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/note"
	"github.com/pauloRohling/locknote/internal/domain/pagination"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/locknote/internal/domain/types/text"
	mocknote "github.com/pauloRohling/locknote/internal/mocks/note"
	"github.com/pauloRohling/locknote/pkg/testinstance"
	"github.com/pauloRohling/throw"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestListNotesUseCase(t *testing.T) {
	userId, _ := id.NewID()
	noteId, _ := id.NewID()
	ctx := context.WithValue(context.Background(), audit.UserIdContextKey, userId)

	mockNote := testinstance.Set(&note.Note{}, map[string]any{
		"id":      noteId,
		"title":   text.Title("Test Note"),
		"content": "Test Note Content",
	})

	testCases := map[string]struct {
		input     *ListNotesInput
		setup     func(noteRepository *mocknote.MockRepository)
		expectErr bool
		errType   throw.ErrorType
	}{
		"should find all notes": {
			input: &ListNotesInput{
				Pagination: pagination.Pagination{Page: 1, Size: 10},
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.EXPECT().FindAllNotes(mock.Anything, mock.Anything).Return([]*note.Note{mockNote}, nil)
			},
		},
		"should return an error if the note repository fails to find all notes": {
			input: &ListNotesInput{
				Pagination: pagination.Pagination{Page: 1, Size: 10},
			},
			setup: func(noteRepository *mocknote.MockRepository) {
				noteRepository.EXPECT().FindAllNotes(mock.Anything, mock.Anything).Return(nil, throw.Internal().Msg("failed to find notes"))
			},
			expectErr: true,
			errType:   throw.InternalErrorType,
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			noteRepository := mocknote.NewMockRepository(t)

			if testCase.setup != nil {
				testCase.setup(noteRepository)
			}

			listNotesUseCase := NewListNotesUseCase(ListNotesParams{
				NoteRepository: noteRepository,
			})

			notes, err := listNotesUseCase.Execute(ctx, testCase.input)

			if testCase.expectErr {
				assert.Error(t, err)
				throw.AssertType(t, err, testCase.errType.String())
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, notes)
		})
	}
}
