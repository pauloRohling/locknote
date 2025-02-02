package user

import (
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/locknote/internal/domain/user"
	"github.com/pauloRohling/locknote/internal/persistence/store"
)

// Mapper is responsible for mapping [store.User] objects to the domain model
type Mapper interface {
	Parse(savedUser *store.User) (*user.User, error)
}

type DefaultMapper struct {
	factory user.Factory
}

func NewMapper(factory user.Factory) *DefaultMapper {
	return &DefaultMapper{
		factory: factory,
	}
}

func (mapper *DefaultMapper) Parse(savedUser *store.User) (*user.User, error) {
	createdBy, err := id.FromUUID(savedUser.CreatedBy)
	if err != nil {
		return nil, err
	}

	return mapper.factory.Parse(user.ParseParams{
		ID:    savedUser.ID,
		Audit: audit.New(savedUser.CreatedAt, createdBy),
		NewParams: user.NewParams{
			Name:     savedUser.Name,
			Email:    savedUser.Email,
			Password: savedUser.Password,
		},
	})
}

// Ensure the mapper implements the [Mapper] interface
var _ Mapper = (*DefaultMapper)(nil)
