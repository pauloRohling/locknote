package user

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/types/email"
)

// Repository is responsible for storing and retrieving users
type Repository interface {
	Save(ctx context.Context, user *User) (*User, error)
	FindByEmail(ctx context.Context, email email.Email) (*User, error)
}
