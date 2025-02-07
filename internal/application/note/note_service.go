package note

import "context"

type Service interface {
	Create(ctx context.Context, input *CreateNoteInput) (*CreateNoteOutput, error)
	GetById(ctx context.Context, input *GetNoteInput) (*GetNoteOutput, error)
	List(ctx context.Context, input *ListNotesInput) (*ListNotesOutput, error)
}

type FacadeServiceParams struct {
	CreateNoteUseCase *CreateNoteUseCase
	GetNoteUseCase    *GetNoteUseCase
	ListNotesUseCase  *ListNotesUseCase
}

type FacadeService struct {
	FacadeServiceParams
}

func NewService(params FacadeServiceParams) *FacadeService {
	return &FacadeService{params}
}

func (service *FacadeService) Create(ctx context.Context, input *CreateNoteInput) (*CreateNoteOutput, error) {
	return service.CreateNoteUseCase.Execute(ctx, input)
}

func (service *FacadeService) GetById(ctx context.Context, input *GetNoteInput) (*GetNoteOutput, error) {
	return service.GetNoteUseCase.Execute(ctx, input)
}

func (service *FacadeService) List(ctx context.Context, input *ListNotesInput) (*ListNotesOutput, error) {
	return service.ListNotesUseCase.Execute(ctx, input)
}

// Ensure the service implements the [note.Service] interface
var _ Service = (*FacadeService)(nil)
