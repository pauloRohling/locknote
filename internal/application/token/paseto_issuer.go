package token

import (
	"aidanwoods.dev/go-paseto"
	"github.com/pauloRohling/locknote/internal/domain/token"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"time"
)

// PasetoIssuer implements the [token.Issuer] interface using Paseto
type PasetoIssuer struct {
	secretKey  paseto.V4AsymmetricSecretKey
	publicKey  paseto.V4AsymmetricPublicKey
	expiration time.Duration
	issuer     string
}

func NewPasetoIssuer(secretKey, publicKey, issuer string, expiration time.Duration) (*PasetoIssuer, error) {
	parsedSecretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(secretKey)
	if err != nil {
		return nil, err
	}

	parsedPublicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(publicKey)
	if err != nil {
		return nil, err
	}

	return &PasetoIssuer{
		secretKey:  parsedSecretKey,
		publicKey:  parsedPublicKey,
		expiration: expiration,
		issuer:     issuer,
	}, nil
}

func (issuer *PasetoIssuer) Issue(payload token.Payload) (string, id.ID, error) {
	tokenId, err := id.NewID()
	if err != nil {
		return "", tokenId, err
	}

	now := time.Now().UTC()
	pasetoToken := paseto.NewToken()
	pasetoToken.SetIssuedAt(now)
	pasetoToken.SetNotBefore(now)
	pasetoToken.SetExpiration(now.Add(issuer.expiration))
	pasetoToken.SetJti(tokenId.String())
	pasetoToken.SetIssuer(issuer.issuer)
	pasetoToken.SetSubject(payload.UserID.String())

	return pasetoToken.V4Sign(issuer.secretKey, nil), tokenId, nil
}

// Ensure the issuer implements the [token.Issuer] interface
var _ token.Issuer = (*PasetoIssuer)(nil)
