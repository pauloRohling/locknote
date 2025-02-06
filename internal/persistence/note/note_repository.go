package note

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/note"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/locknote/internal/persistence/store"
	"github.com/pauloRohling/locknote/pkg/transaction"
	"github.com/pauloRohling/throw"
)

// Repository is the PostgreSQL implementation of [note.Repository]
type Repository struct {
	conn   store.DBTX
	mapper Mapper
}

func NewRepository(conn store.DBTX, mapper Mapper) *Repository {
	return &Repository{
		conn:   conn,
		mapper: mapper,
	}
}

func (repository *Repository) query(ctx context.Context) *store.Queries {
	if tx := transaction.FromContext(ctx); tx != nil {
		return store.New(*tx)
	}
	return store.New(repository.conn)
}

func (repository *Repository) Save(ctx context.Context, note *note.Note) (*note.Note, error) {
	newNote, err := repository.query(ctx).InsertNote(ctx, store.InsertNoteParams{
		ID:        note.ID().UUID(),
		Title:     note.Title().String(),
		Content:   note.Content(),
		CreatedAt: note.Audit().CreatedAt(),
		CreatedBy: note.Audit().CreatedBy().UUID(),
	})

	if err != nil {
		return nil, throw.Internal().Err(err).Msg("could not save note")
	}

	return repository.mapper.Parse(&newNote)
}

func (repository *Repository) FindByID(ctx context.Context, id id.ID) (*note.Note, error) {
	userId, err := audit.GetUserId(ctx)
	if err != nil {
		return nil, throw.Internal().Err(err).Msg("Could not find an UserID in the request")
	}

	params := store.FindNoteByIDParams{
		ID:        id.UUID(),
		CreatedBy: userId.UUID(),
	}

	matchedNote, err := repository.query(ctx).FindNoteByID(ctx, params)
	if err == nil {
		return repository.mapper.Parse(&matchedNote)
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, throw.NotFound().
			Err(err).
			Str("id", id.String()).
			Msg("The provided ID does not match any note")
	}

	return nil, throw.Internal().
		Err(err).
		Str("id", id.String()).
		Msg("A database error occurred while trying to find a note with the provided ID")
}

// Ensure the repository implements the [note.Repository] interface
var _ note.Repository = (*Repository)(nil)
