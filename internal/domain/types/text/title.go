package text

import (
	"encoding/json"
	"github.com/pauloRohling/throw"
	"strings"
)

// Title defines a resource title with a maximum length of 255 characters
type Title string

// String returns the string representation of the title
func (name Title) String() string {
	return string(name)
}

// MarshalJSON implements the [encoding/json.Marshaler] interface
func (name Title) MarshalJSON() ([]byte, error) {
	return json.Marshal(name.String())
}

// NewTitle creates a new [Title] from a string
func NewTitle(name string) (Title, error) {
	name = strings.TrimSpace(name)
	length := len(name)
	if length == 0 || length > 255 {
		return "", throw.Validation().Msg("invalid naming length")
	}

	return Title(name), nil
}
