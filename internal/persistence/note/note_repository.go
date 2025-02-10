package note

import (
	"context"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/note"
	"github.com/pauloRohling/locknote/internal/domain/pagination"
	"github.com/pauloRohling/locknote/internal/domain/types/id"
	"github.com/pauloRohling/locknote/internal/persistence/postgres"
	"github.com/pauloRohling/locknote/internal/persistence/store"
	"github.com/pauloRohling/locknote/pkg/transaction"
	"time"
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
		UpdatedAt: note.Audit().UpdatedAt(),
		UpdatedBy: note.Audit().UpdatedBy().UUID(),
	})

	if err != nil {
		return nil, postgres.Throw(err)
	}

	return repository.mapper.Parse(newNote)
}

func (repository *Repository) FindByID(ctx context.Context, id id.ID) (*note.Note, error) {
	userId, err := audit.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	params := store.FindNoteByIDParams{
		ID:        id.UUID(),
		CreatedBy: userId.UUID(),
	}

	matchedNote, err := repository.query(ctx).FindNoteByID(ctx, params)
	if err != nil {
		return nil, postgres.ThrowNotFound(err)
	}

	return repository.mapper.Parse(matchedNote)
}

func (repository *Repository) FindAll(ctx context.Context, pagination pagination.Pagination) ([]*note.Note, error) {
	userId, err := audit.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	params := store.FindNotesByUserParams{
		CreatedBy: userId.UUID(),
		Limit:     pagination.Limit(),
		Offset:    pagination.Offset(),
	}

	matchedNotes, err := repository.query(ctx).FindNotesByUser(ctx, params)
	if err != nil {
		return nil, postgres.Throw(err)
	}

	return repository.mapper.ParseMany(matchedNotes)
}

func (repository *Repository) Update(ctx context.Context, note *note.Note) (*note.Note, error) {
	userId, err := audit.GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	params := store.UpdateNoteParams{
		ID:        note.ID().UUID(),
		CreatedBy: userId.UUID(),
		Title:     note.Title().String(),
		Content:   note.Content(),
		UpdatedAt: time.Now().UTC(),
		UpdatedBy: userId.UUID(),
	}

	updatedNote, err := repository.query(ctx).UpdateNote(ctx, params)
	if err != nil {
		return nil, postgres.Throw(err)
	}

	return repository.mapper.Parse(updatedNote)
}

func (repository *Repository) Delete(ctx context.Context, noteId id.ID) error {
	userId, err := audit.GetUserId(ctx)
	if err != nil {
		return err
	}

	err = repository.query(ctx).DeleteNote(ctx, store.DeleteNoteParams{
		ID:        noteId.UUID(),
		CreatedBy: userId.UUID(),
	})

	return postgres.Throw(err)
}

func (repository *Repository) DeleteAllByUser(ctx context.Context, userId id.ID) error {
	err := repository.query(ctx).DeleteNotesByUser(ctx, userId.UUID())
	return postgres.Throw(err)
}

// Ensure the repository implements the [note.Repository] interface
var _ note.Repository = (*Repository)(nil)
