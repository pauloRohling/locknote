package user

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/types/email"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
)

// Repository is responsible for storing and retrieving users
type Repository interface {
	DeleteById(ctx context.Context, userId id.ID) error
	FindByEmail(ctx context.Context, email email.Email) (*User, error)
	FindByID(ctx context.Context, userId id.ID) (*User, error)
	Save(ctx context.Context, user *User) (*User, error)
	UpdateById(ctx context.Context, user *User) (*User, error)
}
