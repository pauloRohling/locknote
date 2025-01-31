package note

import (
	"context"
	"github.com/google/uuid"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/locknote/internal/domain/types/text"
)

// NewParams represents the necessary parameters for creating a new [note.Note]
type NewParams struct {
	Title   string
	Content string
}

// ParseParams represents the necessary parameters for parsing a previously saved [note.Note]
type ParseParams struct {
	ID    uuid.UUID
	Audit audit.Audit
	NewParams
}

// Factory is used for creating new [note.Note] objects
type Factory interface {
	New(ctx context.Context, params NewParams) (*Note, error)
	Parse(params ParseParams) (*Note, error)
}

type DefaultFactory struct {
}

func NewFactory() *DefaultFactory {
	return &DefaultFactory{}
}

func (factory *DefaultFactory) New(ctx context.Context, params NewParams) (*Note, error) {
	noteId, err := id.NewID()
	if err != nil {
		return nil, err
	}

	userId, err := audit.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	return factory.Parse(ParseParams{
		ID:        noteId.UUID(),
		Audit:     audit.NewDefault(userId),
		NewParams: params,
	})
}

func (factory *DefaultFactory) Parse(params ParseParams) (*Note, error) {
	noteId, err := id.FromUUID(params.ID)
	if err != nil {
		return nil, err
	}

	title, err := text.NewTitle(params.Title)
	if err != nil {
		return nil, err
	}

	return &Note{
		id:      noteId,
		title:   title,
		content: params.Content,
		audit:   params.Audit,
	}, nil
}

// Ensure the factory implements the [Factory] interface
var _ Factory = (*DefaultFactory)(nil)
