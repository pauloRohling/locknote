package email

import (
	"encoding/json"
	"github.com/pauloRohling/throw"
	"net/mail"
	"strings"
)

// Email defines a email address
type Email string

// String returns the string representation of the email
func (email Email) String() string {
	return string(email)
}

// MarshalJSON implements the [encoding/json.Marshaler] interface
func (email Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(email.String())
}

// NewEmail creates a new [Email] from a string
func NewEmail(email string) (Email, error) {
	email = strings.TrimSpace(email)
	length := len(email)
	if length == 0 {
		return "", throw.Validation().Msg("email should	not be empty")
	}

	if length > 255 {
		return "", throw.Validation().Msg("email should not be longer than 255 characters")
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return "", throw.Validation().Msg("invalid email format")
	}

	return Email(email), nil
}
