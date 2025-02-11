package note

import (
	"context"
	"github.com/pauloRohling/locknote/internal/application"
	"github.com/pauloRohling/locknote/internal/domain/note"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
)

type UpdateNoteInput struct {
	ID      string
	Title   string
	Content string
}

type UpdateNoteOutput struct {
	Note *note.Note
}

type UpdateNoteParams struct {
	NoteFactory    note.Factory
	NoteRepository note.Repository
}

type UpdateNoteUseCase struct {
	UpdateNoteParams
}

func NewUpdateNoteUseCase(params UpdateNoteParams) *UpdateNoteUseCase {
	return &UpdateNoteUseCase{UpdateNoteParams: params}
}

func (usecase *UpdateNoteUseCase) Execute(ctx context.Context, input *UpdateNoteInput) (*UpdateNoteOutput, error) {
	noteId, err := id.FromString(input.ID)
	if err != nil {
		return nil, err
	}

	matchedNote, err := usecase.NoteRepository.FindByID(ctx, noteId)
	if err != nil {
		return nil, err
	}

	if err = matchedNote.Update(input.Title, input.Content); err != nil {
		return nil, err
	}

	matchedNote, err = usecase.NoteRepository.UpdateById(ctx, matchedNote)
	if err != nil {
		return nil, err
	}

	return &UpdateNoteOutput{Note: matchedNote}, nil
}

// Ensure the usecase implements the interface
var _ application.UseCase[UpdateNoteInput, UpdateNoteOutput] = (*UpdateNoteUseCase)(nil)
