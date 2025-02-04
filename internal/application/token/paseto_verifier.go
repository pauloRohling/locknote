package token

import (
	"aidanwoods.dev/go-paseto"
	"github.com/pauloRohling/locknote/internal/domain/token"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/throw"
)

// PasetoVerifier implements the [token.Verifier] interface using Paseto
type PasetoVerifier struct {
	publicKey paseto.V4AsymmetricPublicKey
	parser    *paseto.Parser
}

func NewPasetoVerifier(publicKey, issuer string) (*PasetoVerifier, error) {
	parsedPublicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(publicKey)
	if err != nil {
		return nil, err
	}

	parser := paseto.NewParser()
	parser.AddRule(paseto.IssuedBy(issuer))
	parser.AddRule(paseto.NotExpired())
	parser.AddRule(paseto.NotBeforeNbf())

	return &PasetoVerifier{
		publicKey: parsedPublicKey,
		parser:    &parser,
	}, nil
}

func (verifier *PasetoVerifier) Verify(pasetoToken string) (*token.Payload, error) {
	parsedToken, err := verifier.parser.ParseV4Public(verifier.publicKey, pasetoToken, nil)
	if err != nil {
		return nil, err
	}

	subjectClaim, err := parsedToken.GetSubject()
	if err != nil {
		return nil, throw.Unauthorized().Err(err).Msg("subject claim not found")
	}

	subject, err := id.FromString(subjectClaim)
	if err != nil {
		return nil, throw.Unauthorized().Err(err).Msg("subject should be a valid UUID V7")
	}

	return &token.Payload{UserID: subject}, nil
}

// Ensure the verifier implements the [token.Verifier] interface
var _ token.Verifier = (*PasetoVerifier)(nil)
