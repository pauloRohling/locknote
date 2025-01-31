package id

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pauloRohling/throw"
)

// Nil is a constant for a nil UUID
var Nil = ID(uuid.Nil)

// ID is a type alias for a UUID V7
type ID uuid.UUID

// String returns the string representation of the [ID]
func (id ID) String() string {
	return uuid.UUID(id).String()
}

// UUID returns the UUID representation of the [ID]
func (id ID) UUID() uuid.UUID {
	return uuid.UUID(id)
}

// MarshalJSON implements the [encoding/json.Marshaler] interface
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(uuid.UUID(id))
}

// NewID creates a new [ID] using the uuid.NewV7() function
func NewID() (ID, error) {
	uuidV7, err := uuid.NewV7()
	if err != nil {
		return ID(uuid.Nil), throw.Internal().Err(err).Msg("failed to create new id")
	}
	return ID(uuidV7), nil
}

// FromString creates a new [ID] from a string
func FromString(id string) (ID, error) {
	uuidV7, err := uuid.Parse(id)
	if err != nil || uuidV7 == uuid.Nil {
		return ID(uuid.Nil), throw.Validation().Err(err).Msg("invalid id format")
	}

	if uuidV7.Version() != 7 {
		return ID(uuid.Nil), throw.Validation().Msg("invalid id version")
	}

	return ID(uuidV7), nil
}

// FromUUID creates a new [ID] from a UUID
func FromUUID(uuid uuid.UUID) (ID, error) {
	return FromString(uuid.String())
}

// FromOptionalUUID creates a new [ID] from a pointer to a UUID
func FromOptionalUUID(uuid *uuid.UUID) (*ID, error) {
	if uuid == nil {
		return nil, nil
	}

	newId, err := FromUUID(*uuid)
	if err != nil {
		return nil, err
	}

	return &newId, nil
}
