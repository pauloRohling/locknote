package token

import "github.com/pauloRohling/locknote/internal/domain/types/id"

// Issuer is responsible for generating a token
type Issuer interface {
	Issue(payload Payload) (string, id.ID, error)
}
