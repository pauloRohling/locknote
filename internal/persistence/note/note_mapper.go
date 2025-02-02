package note

import (
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/note"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/locknote/internal/persistence/store"
)

// Mapper is responsible for mapping [store.Note] objects to the domain model
type Mapper interface {
	Parse(savedNote *store.Note) (*note.Note, error)
}

type DefaultMapper struct {
	factory note.Factory
}

func NewMapper(factory note.Factory) *DefaultMapper {
	return &DefaultMapper{
		factory: factory,
	}
}

func (mapper *DefaultMapper) Parse(savedNote *store.Note) (*note.Note, error) {
	createdBy, err := id.FromUUID(savedNote.CreatedBy)
	if err != nil {
		return nil, err
	}

	return mapper.factory.Parse(note.ParseParams{
		ID:    savedNote.ID,
		Audit: audit.New(savedNote.CreatedAt, createdBy),
		NewParams: note.NewParams{
			Title:   savedNote.Title,
			Content: savedNote.Content,
		},
	})
}

// Ensure the mapper implements the [Mapper] interface
var _ Mapper = (*DefaultMapper)(nil)
