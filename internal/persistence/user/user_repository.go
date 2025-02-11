package user

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/types/email"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/locknote/internal/domain/user"
	"github.com/pauloRohling/locknote/internal/persistence/postgres"
	"github.com/pauloRohling/locknote/internal/persistence/store"
	"github.com/pauloRohling/locknote/pkg/transaction"
	"time"
)

// Repository is the PostgreSQL implementation of [user.Repository]
type Repository struct {
	conn   store.DBTX
	mapper Mapper
}

func NewRepository(conn store.DBTX, mapper Mapper) *Repository {
	return &Repository{
		conn:   conn,
		mapper: mapper,
	}
}

func (repository *Repository) query(ctx context.Context) *store.Queries {
	if tx := transaction.FromContext(ctx); tx != nil {
		return store.New(*tx)
	}
	return store.New(repository.conn)
}

func (repository *Repository) DeleteById(ctx context.Context, userId id.ID) error {
	err := repository.query(ctx).DeleteUserById(ctx, userId.UUID())
	return postgres.Throw(err)
}

func (repository *Repository) FindByEmail(ctx context.Context, email email.Email) (*user.User, error) {
	matchedUser, err := repository.query(ctx).FindUserByEmail(ctx, email.String())
	if err != nil {
		return nil, postgres.ThrowNotFound(err)
	}
	return repository.mapper.Parse(matchedUser)
}

func (repository *Repository) FindByID(ctx context.Context, userId id.ID) (*user.User, error) {
	matchedUser, err := repository.query(ctx).FindUserByID(ctx, userId.UUID())
	if err != nil {
		return nil, postgres.ThrowNotFound(err)
	}
	return repository.mapper.Parse(matchedUser)
}

func (repository *Repository) UpdateById(ctx context.Context, user *user.User) (*user.User, error) {
	userId, err := audit.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	updatedUser, err := repository.query(ctx).UpdateUserById(ctx, store.UpdateUserByIdParams{
		ID:        user.ID().UUID(),
		Name:      user.Name().String(),
		UpdatedAt: time.Now().UTC(),
		UpdatedBy: userId.UUID(),
	})

	if err != nil {
		return nil, postgres.Throw(err)
	}

	return repository.mapper.Parse(updatedUser)
}

func (repository *Repository) Save(ctx context.Context, user *user.User) (*user.User, error) {
	newUser, err := repository.query(ctx).InsertUser(ctx, store.InsertUserParams{
		ID:        user.ID().UUID(),
		Name:      user.Name().String(),
		Email:     user.Email().String(),
		Password:  user.Password().String(),
		CreatedAt: user.Audit().CreatedAt(),
		CreatedBy: user.Audit().CreatedBy().UUID(),
		UpdatedAt: user.Audit().UpdatedAt(),
		UpdatedBy: user.Audit().UpdatedBy().UUID(),
	})

	if err != nil {
		return nil, postgres.Throw(err)
	}

	return repository.mapper.Parse(newUser)
}

// Ensure the repository implements the [user.Repository] interface
var _ user.Repository = (*Repository)(nil)
