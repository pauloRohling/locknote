package token

// Verifier is responsible for verifying a token
type Verifier interface {
	Verify(token string) (*Payload, error)
}
