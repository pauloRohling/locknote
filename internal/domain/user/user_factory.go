package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/types/email"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/locknote/internal/domain/types/password"
	"github.com/pauloRohling/locknote/internal/domain/types/text"
)

// NewParams represents the necessary parameters for creating a new [user.User]
type NewParams struct {
	Name     string
	Email    string
	Password string
}

// ParseParams represents the necessary parameters for parsing a previously saved [user.User]
type ParseParams struct {
	ID    uuid.UUID
	Audit audit.Audit
	NewParams
}

// Factory is used for creating new [user.User] objects
type Factory interface {
	New(ctx context.Context, params NewParams) (*User, error)
	Parse(params ParseParams) (*User, error)
}

type DefaultFactory struct {
}

func NewFactory() *DefaultFactory {
	return &DefaultFactory{}
}

func (factory *DefaultFactory) New(ctx context.Context, params NewParams) (*User, error) {
	userId, err := id.NewID()
	if err != nil {
		return nil, err
	}

	return factory.Parse(ParseParams{
		ID:        userId.UUID(),
		Audit:     audit.NewDefault(userId),
		NewParams: params,
	})
}

func (factory *DefaultFactory) Parse(params ParseParams) (*User, error) {
	userId, err := id.FromUUID(params.ID)
	if err != nil {
		return nil, err
	}

	userName, err := text.NewPersonName(params.Name)
	if err != nil {
		return nil, err
	}

	userEmail, err := email.NewEmail(params.Email)
	if err != nil {
		return nil, err
	}

	userPassword, err := password.New(params.Password)
	if err != nil {
		return nil, err
	}

	return &User{
		id:       userId,
		name:     userName,
		email:    userEmail,
		password: userPassword,
		audit:    params.Audit,
	}, nil
}

// Ensure the factory implements the [Factory] interface
var _ Factory = (*DefaultFactory)(nil)
