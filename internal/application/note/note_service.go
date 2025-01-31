package note

import "context"

type Service interface {
	Create(ctx context.Context, input *CreateNoteInput) (*CreateNoteOutput, error)
}

type FacadeServiceParams struct {
	CreateUsecase *CreateNoteUseCase
}

type FacadeService struct {
	FacadeServiceParams
}

func NewService(params FacadeServiceParams) *FacadeService {
	return &FacadeService{params}
}

func (service *FacadeService) Create(ctx context.Context, input *CreateNoteInput) (*CreateNoteOutput, error) {
	return service.CreateUsecase.Execute(ctx, input)
}

// Ensure the service implements the [note.Service] interface
var _ Service = (*FacadeService)(nil)
