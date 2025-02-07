package pagination

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/pagination"
	"github.com/pauloRohling/throw"
)

const (
	// PaginationContextKey is the key used to store the pagination params in the context
	PaginationContextKey = "pagination-id"
)

// GetPagination returns the pagination params from the context
func GetPagination(ctx context.Context) (pagination.Pagination, error) {
	params, ok := ctx.Value(PaginationContextKey).(pagination.Pagination)
	if !ok {
		return pagination.Pagination{}, throw.Validation().Msg("Pagination is not available in the context")
	}
	return params, nil
}

// SetPagination sets the pagination params in the context
func SetPagination(ctx context.Context, params pagination.Pagination) context.Context {
	return context.WithValue(ctx, PaginationContextKey, params)
}
