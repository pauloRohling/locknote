package password

import (
	"github.com/pauloRohling/throw"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

// Password defines a bcrypt password
type Password string

// String returns the string representation of the password
func (password Password) String() string {
	return string(password)
}

// MarshalJSON implements the [encoding/json.Marshaler] interface
func (password Password) MarshalJSON() ([]byte, error) {
	return []byte{}, nil
}

// Equals compares a bcrypt hashed password with its possible plaintext equivalent
func (password Password) Equals(text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(text))
	return err == nil
}

// New creates a new bcrypt [Password] from a string
func New(password string) (Password, error) {
	length := len(password)
	if length < 8 {
		return "", throw.Validation().Msg("should have at least 8 characters")
	}

	if length > 70 {
		return "", throw.Validation().Msg("should not exceed 70 characters")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", throw.Validation().Err(err).Msg("failed to encrypt password")
	}

	return Password(encryptedPassword), nil
}

// FromEncrypted creates a new [Password] from an encrypted string.
// Should NOT be used to validate a user-inputted password.
func FromEncrypted(password string) (Password, error) {
	pattern, err := regexp.CompilePOSIX("^\\$2[ayb]\\$.{56}$")
	if err != nil {
		return "", throw.Validation().Err(err).Msg("invalid password format")
	}

	if !pattern.MatchString(password) {
		return "", throw.Validation().Msg("invalid password format")
	}

	return Password(password), nil
}
