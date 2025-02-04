package note

import (
	"context"
	"github.com/pauloRohling/locknote/internal/application"
	"github.com/pauloRohling/locknote/internal/domain/note"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
)

type GetNoteInput struct {
	NoteID string
}

type GetNoteOutput struct {
	Note *note.Note
}

type GetNoteParams struct {
	NoteRepository note.Repository
}

type GetNoteUseCase struct {
	GetNoteParams
}

func NewGetNoteUseCase(params GetNoteParams) *GetNoteUseCase {
	return &GetNoteUseCase{GetNoteParams: params}
}

func (usecase *GetNoteUseCase) Execute(ctx context.Context, input *GetNoteInput) (*GetNoteOutput, error) {
	noteId, err := id.FromString(input.NoteID)
	if err != nil {
		return nil, err
	}

	matchedNote, err := usecase.NoteRepository.FindByID(ctx, noteId)
	if err != nil {
		return nil, err
	}

	return &GetNoteOutput{Note: matchedNote}, nil
}

// Ensure the usecase implements the interface
var _ application.UseCase[GetNoteInput, GetNoteOutput] = (*GetNoteUseCase)(nil)
