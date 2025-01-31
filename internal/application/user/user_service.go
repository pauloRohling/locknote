package user

import "context"

type Service interface {
	Create(ctx context.Context, input *CreateUserInput) (*CreateUserOutput, error)
	Login(ctx context.Context, input *LoginInput) (*LoginOutput, error)
}

type FacadeServiceParams struct {
	CreateUsecase *CreateUserUseCase
	LoginUsecase  *LoginUseCase
}

type FacadeService struct {
	FacadeServiceParams
}

func NewService(params FacadeServiceParams) *FacadeService {
	return &FacadeService{params}
}

func (service *FacadeService) Create(ctx context.Context, input *CreateUserInput) (*CreateUserOutput, error) {
	return service.CreateUsecase.Execute(ctx, input)
}

func (service *FacadeService) Login(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	return service.LoginUsecase.Execute(ctx, input)
}

// Ensure the service implements the [user.Service] interface
var _ Service = (*FacadeService)(nil)
