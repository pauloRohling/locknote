package user

import "context"

type Service interface {
	Create(ctx context.Context, input *CreateUserInput) (*CreateUserOutput, error)
	Login(ctx context.Context, input *LoginInput) (*LoginOutput, error)
	Get(ctx context.Context) (*GetUserOutput, error)
}

type FacadeServiceParams struct {
	CreateUseCase *CreateUserUseCase
	LoginUseCase  *LoginUseCase
	GetUseCase    *GetUserUseCase
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

func (service *FacadeService) Get(ctx context.Context) (*GetUserOutput, error) {
	return service.GetUseCase.Execute(ctx)
}

// Ensure the service implements the [user.Service] interface
var _ Service = (*FacadeService)(nil)
