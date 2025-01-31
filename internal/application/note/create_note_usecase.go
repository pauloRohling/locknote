package note

import (
	"context"
	"github.com/pauloRohling/locknote/internal/application"
	"github.com/pauloRohling/locknote/internal/domain/note"
)

type CreateNoteInput struct {
	Title   string
	Content string
}

type CreateNoteOutput struct {
	Note *note.Note
}

type CreateNoteParams struct {
	NoteFactory    note.Factory
	NoteRepository note.Repository
}

type CreateNoteUseCase struct {
	CreateNoteParams
}

func NewCreateNoteUseCase(params CreateNoteParams) *CreateNoteUseCase {
	return &CreateNoteUseCase{CreateNoteParams: params}
}

func (usecase *CreateNoteUseCase) Execute(ctx context.Context, input *CreateNoteInput) (*CreateNoteOutput, error) {
	createNoteParams := note.NewParams{
		Title:   input.Title,
		Content: input.Content,
	}

	newNote, err := usecase.NoteFactory.New(ctx, createNoteParams)
	if err != nil {
		return nil, err
	}

	newNote, err = usecase.NoteRepository.Save(ctx, newNote)
	if err != nil {
		return nil, err
	}

	return &CreateNoteOutput{Note: newNote}, nil
}

// Ensure the usecase implements the interface
var _ application.UseCase[CreateNoteInput, CreateNoteOutput] = (*CreateNoteUseCase)(nil)
