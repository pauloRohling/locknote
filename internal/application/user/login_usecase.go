package user

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/token"
	"github.com/pauloRohling/locknote/internal/domain/types/email"
	"github.com/pauloRohling/locknote/internal/domain/user"
	"github.com/pauloRohling/throw"
)

const (
	defaultLoginErrMessage = "Invalid credentials. Please check your email and password"
)

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	AccessToken string
}

type LoginUsecaseParams struct {
	TokenIssuer    token.Issuer
	UserRepository user.Repository
}

type LoginUseCase struct {
	LoginUsecaseParams
}

func NewLoginUseCase(params LoginUsecaseParams) *LoginUseCase {
	return &LoginUseCase{params}
}

func (usecase *LoginUseCase) Execute(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	userEmail, err := email.NewEmail(input.Email)
	if err != nil {
		return nil, throw.Unauthorized().Err(err).Msg(defaultLoginErrMessage)
	}

	matchedUser, err := usecase.UserRepository.FindByEmail(ctx, userEmail)
	if err != nil {
		return nil, throw.Unauthorized().Err(err).Msg(defaultLoginErrMessage)
	}

	if !matchedUser.Password().Equals(input.Password) {
		return nil, throw.Unauthorized().Err(err).Msg(defaultLoginErrMessage)
	}

	tokenPayload := token.NewPayload(matchedUser.ID())

	issuedToken, _, err := usecase.TokenIssuer.Issue(tokenPayload)
	if err != nil {
		return nil, throw.Unauthorized().Err(err).Msg(defaultLoginErrMessage)
	}

	return &LoginOutput{AccessToken: issuedToken}, nil
}
