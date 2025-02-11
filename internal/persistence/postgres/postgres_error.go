package postgres

import (
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pauloRohling/throw"
)

const (
	DefaultErrorMsg  = "An unexpected database error occurred"
	NotFoundErrorMsg = "No records found matching the given criteria"
)

func Throw(err error) error {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return throw.Internal().Err(err).Msg(DefaultErrorMsg)
	}

	switch pgErr.Code {
	case pgerrcode.UniqueViolation:
		if msg, exists := UniqueViolationErrors[pgErr.ConstraintName]; exists {
			return throw.Conflict().Err(err).Msg(msg)
		}
		return throw.Conflict().
			Err(err).
			Str("code", pgErr.Code).
			Str("detail", pgErr.Detail).
			Str("constraintName", pgErr.ConstraintName).
			Msg("A unique constraint violation occurred")
	default:
		return throw.Internal().
			Err(err).
			Str("code", pgErr.Code).
			Str("detail", pgErr.Detail).
			Msg(DefaultErrorMsg)
	}
}

func ThrowNotFound(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return throw.NotFound().Err(err).Msg(NotFoundErrorMsg)
	}

	return throw.Internal().Err(err).Msg(DefaultErrorMsg)
}
