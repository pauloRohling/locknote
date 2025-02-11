package note

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/pagination"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
)

type Repository interface {
	DeleteAllByUser(ctx context.Context, userId id.ID) error
	DeleteById(ctx context.Context, noteId id.ID) error
	FindAll(ctx context.Context, pagination pagination.Pagination) ([]*Note, error)
	FindByID(ctx context.Context, id id.ID) (*Note, error)
	Save(ctx context.Context, note *Note) (*Note, error)
	UpdateById(ctx context.Context, note *Note) (*Note, error)
}
