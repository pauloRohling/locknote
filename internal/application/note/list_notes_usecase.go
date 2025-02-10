package note

import (
	"context"
	"github.com/pauloRohling/locknote/internal/application"
	"github.com/pauloRohling/locknote/internal/domain/note"
	"github.com/pauloRohling/locknote/internal/domain/pagination"
)

type ListNotesInput struct {
	Pagination pagination.Pagination
}

type ListNotesOutput struct {
	Notes []*note.Note
}

type ListNotesParams struct {
	NoteRepository note.Repository
}

type ListNotesUseCase struct {
	ListNotesParams
}

func NewListNotesUseCase(params ListNotesParams) *ListNotesUseCase {
	return &ListNotesUseCase{ListNotesParams: params}
}

func (usecase *ListNotesUseCase) Execute(ctx context.Context, input *ListNotesInput) (*ListNotesOutput, error) {
	notes, err := usecase.NoteRepository.FindAll(ctx, input.Pagination)
	if err != nil {
		return nil, err
	}

	return &ListNotesOutput{Notes: notes}, nil
}

// Ensure the usecase implements the interface
var _ application.UseCase[ListNotesInput, ListNotesOutput] = (*ListNotesUseCase)(nil)
