package note

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/pagination"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
)

type Repository interface {
	Save(ctx context.Context, note *Note) (*Note, error)
	FindByID(ctx context.Context, id id.ID) (*Note, error)
	FindAll(ctx context.Context, pagination pagination.Pagination) ([]*Note, error)
	Update(ctx context.Context, note *Note) (*Note, error)
	Delete(ctx context.Context, noteId id.ID) error
	DeleteAllByUser(ctx context.Context, userId id.ID) error
}
