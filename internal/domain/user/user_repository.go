package user

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/types/email"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
)

// Repository is responsible for storing and retrieving users
type Repository interface {
	Save(ctx context.Context, user *User) (*User, error)
	FindByEmail(ctx context.Context, email email.Email) (*User, error)
	FindByID(ctx context.Context, userId id.ID) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, userId id.ID) error
}
