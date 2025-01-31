package text

import (
	"encoding/json"
	"github.com/pauloRohling/throw"
	"strings"
)

// Title defines a resource title with a maximum length of 255 characters
type Title string

// String returns the string representation of the title
func (title Title) String() string {
	return string(title)
}

// MarshalJSON implements the [encoding/json.Marshaler] interface
func (title Title) MarshalJSON() ([]byte, error) {
	return json.Marshal(title.String())
}

// NewTitle creates a new [Title] from a string
func NewTitle(title string) (Title, error) {
	title = strings.TrimSpace(title)
	length := len(title)
	if length == 0 {
		return "", throw.Validation().Msg("title should not be empty")
	}

	if length > 255 {
		return "", throw.Validation().Msg("title length cannot exceed 255 characters")
	}

	return Title(title), nil
}
