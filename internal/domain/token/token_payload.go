package token

import (
	"github.com/pauloRohling/locknote/internal/domain/types/id"
)

// Payload is the payload of the token
type Payload struct {
	UserID id.ID
}

func NewPayload(userId id.ID) Payload {
	return Payload{UserID: userId}
}
