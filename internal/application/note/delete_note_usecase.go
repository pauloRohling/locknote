package note

import (
	"context"
	"github.com/pauloRohling/locknote/internal/application"
	"github.com/pauloRohling/locknote/internal/domain/note"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
)

type DeleteNoteInput struct {
	NoteID string
}

type DeleteNoteOutput struct {
	NoteID id.ID
}

type DeleteNoteParams struct {
	NoteRepository note.Repository
}

type DeleteNoteUseCase struct {
	DeleteNoteParams
}

func NewDeleteNoteUseCase(params DeleteNoteParams) *DeleteNoteUseCase {
	return &DeleteNoteUseCase{DeleteNoteParams: params}
}

func (usecase *DeleteNoteUseCase) Execute(ctx context.Context, input *DeleteNoteInput) (*DeleteNoteOutput, error) {
	noteId, err := id.FromString(input.NoteID)
	if err != nil {
		return nil, err
	}

	err = usecase.NoteRepository.DeleteById(ctx, noteId)
	if err != nil {
		return nil, err
	}

	return &DeleteNoteOutput{NoteID: noteId}, nil
}

// Ensure the usecase implements the interface
var _ application.UseCase[DeleteNoteInput, DeleteNoteOutput] = (*DeleteNoteUseCase)(nil)
