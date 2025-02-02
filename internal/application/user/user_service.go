package user

import "context"

type Service interface {
	Create(ctx context.Context, input *CreateUserInput) (*CreateUserOutput, error)
	Login(ctx context.Context, input *LoginInput) (*LoginOutput, error)
}

type FacadeServiceParams struct {
	CreateUseCase *CreateUserUseCase
	LoginUseCase  *LoginUseCase
}

type FacadeService struct {
	FacadeServiceParams
}

func NewService(params FacadeServiceParams) *FacadeService {
	return &FacadeService{params}
}

func (service *FacadeService) Create(ctx context.Context, input *CreateUserInput) (*CreateUserOutput, error) {
	return service.CreateUseCase.Execute(ctx, input)
}

func (service *FacadeService) Login(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	return service.LoginUseCase.Execute(ctx, input)
}

// Ensure the service implements the [user.Service] interface
var _ Service = (*FacadeService)(nil)
