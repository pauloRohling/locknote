package audit

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/throw"
)

const (
	// UserIdContextKey is the key used to store in the context the user id of the current user
	UserIdContextKey = "user-id"
)

// GetUserId returns the user id of the current user from the context
func GetUserId(ctx context.Context) (id.ID, error) {
	userId, ok := ctx.Value(UserIdContextKey).(id.ID)
	if !ok {
		return id.Nil, throw.Validation().Msg("user id is not available in the context")
	}
	return userId, nil
}

// SetUserId sets the user id of the current user in the context
func SetUserId(ctx context.Context, userId id.ID) context.Context {
	return context.WithValue(ctx, UserIdContextKey, userId)
}
