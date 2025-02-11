package user

import (
	"context"
	"github.com/pauloRohling/locknote/internal/application"
	"github.com/pauloRohling/locknote/internal/domain/user"
)

type GetUserOutput struct {
	User *user.User
}

type GetUserParams struct {
	UserRepository user.Repository
}

type GetUserUseCase struct {
	GetUserParams
}

func NewGetUserUseCase(params GetUserParams) *GetUserUseCase {
	return &GetUserUseCase{GetUserParams: params}
}

func (usecase *GetUserUseCase) Execute(ctx context.Context) (*GetUserOutput, error) {
	matchedUser, err := usecase.UserRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	return &GetUserOutput{User: matchedUser}, nil
}

// Ensure the usecase implements the interface
var _ application.InnerUseCase[GetUserOutput] = (*GetUserUseCase)(nil)
