package text

import (
	"encoding/json"
	"github.com/pauloRohling/throw"
	"strings"
)

// PersonName defines a person naming with a maximum length of 50 characters
type PersonName string

// String returns the string representation of the naming
func (name PersonName) String() string {
	return string(name)
}

// MarshalJSON implements the [encoding/json.Marshaler] interface
func (name PersonName) MarshalJSON() ([]byte, error) {
	return json.Marshal(name.String())
}

// NewPersonName creates a new [PersonName] from a string
func NewPersonName(name string) (PersonName, error) {
	name = strings.TrimSpace(name)
	length := len(name)
	if length == 0 {
		return "", throw.Validation().Msg("should not be empty")
	}

	if length > 50 {
		return "", throw.Validation().Msg("should not exceed 50 characters")
	}

	return PersonName(name), nil
}
