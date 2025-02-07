package note

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/pagination"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
)

type Repository interface {
	Save(ctx context.Context, note *Note) (*Note, error)
	FindByID(ctx context.Context, id id.ID) (*Note, error)
	FindAllNotes(ctx context.Context, pagination pagination.Pagination) ([]*Note, error)
}
