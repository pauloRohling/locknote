package user

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/types/email"
	"github.com/pauloRohling/locknote/internal/domain/user"
	"github.com/pauloRohling/locknote/internal/persistence/store"
	"github.com/pauloRohling/locknote/pkg/transaction"
	"github.com/pauloRohling/throw"
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

func (repository *Repository) Save(ctx context.Context, user *user.User) (*user.User, error) {
	newUser, err := repository.query(ctx).InsertUser(ctx, store.InsertUserParams{
		ID:        user.ID().UUID(),
		Name:      user.Name().String(),
		Email:     user.Email().String(),
		Password:  user.Password().String(),
		CreatedAt: user.Audit().CreatedAt(),
		CreatedBy: user.Audit().CreatedBy().UUID(),
	})

	if err != nil {
		return nil, throw.Internal().Err(err).Msg("could not save user")
	}

	return repository.mapper.Parse(&newUser)
}

func (repository *Repository) FindByEmail(ctx context.Context, email email.Email) (*user.User, error) {
	matchedUser, err := repository.query(ctx).FindUserByEmail(ctx, email.String())
	if err != nil {
		return nil, throw.Internal().Err(err).Msg("could not find user")
	}

	return repository.mapper.Parse(&matchedUser)
}

// Ensure the repository implements the [user.Repository] interface
var _ user.Repository = (*Repository)(nil)
