package user

import (
	"context"
	"github.com/pauloRohling/locknote/internal/application"
	"github.com/pauloRohling/locknote/internal/domain/user"
)

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

type CreateUserOutput struct {
	User *user.User
}

type CreateUserParams struct {
	UserFactory    user.Factory
	UserRepository user.Repository
}

type CreateUserUseCase struct {
	CreateUserParams
}

func NewCreateUserUseCase(params CreateUserParams) *CreateUserUseCase {
	return &CreateUserUseCase{CreateUserParams: params}
}

func (usecase *CreateUserUseCase) Execute(ctx context.Context, input *CreateUserInput) (*CreateUserOutput, error) {
	createUserParams := user.NewParams{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	newUser, err := usecase.UserFactory.New(ctx, createUserParams)
	if err != nil {
		return nil, err
	}

	newUser, err = usecase.UserRepository.Save(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &CreateUserOutput{User: newUser}, nil
}

// Ensure the usecase implements the interface
var _ application.UseCase[CreateUserInput, CreateUserOutput] = (*CreateUserUseCase)(nil)
