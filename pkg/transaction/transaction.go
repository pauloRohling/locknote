package transaction

import (
	"context"
	"github.com/jackc/pgx/v5"
)

const contextKey = "transaction"

// Inject injects a transaction into the context
func Inject(ctx context.Context, tx *pgx.Tx) context.Context {
	return context.WithValue(ctx, contextKey, tx)
}

// Clean returns a context without the transaction
func Clean(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKey, nil)
}

// FromContext returns the transaction from the context or nil
// if there is no transaction in the context
func FromContext(ctx context.Context) *pgx.Tx {
	if tx, ok := ctx.Value(contextKey).(*pgx.Tx); ok {
		return tx
	}
	return nil
}
