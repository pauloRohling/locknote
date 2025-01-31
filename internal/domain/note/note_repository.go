package note

import "context"

type Repository interface {
	Save(ctx context.Context, note *Note) (*Note, error)
}
