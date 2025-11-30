package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/copy-paste-service/internal/domain"
	"github.com/copy-paste-service/internal/repository/postgres/sqlcgen"
)

// NoteRepository implements repository.NoteRepository using PostgreSQL
type NoteRepository struct {
	pool    *pgxpool.Pool
	queries *sqlcgen.Queries
}

// NewNoteRepository creates a new PostgreSQL note repository
func NewNoteRepository(pool *pgxpool.Pool) *NoteRepository {
	return &NoteRepository{
		pool:    pool,
		queries: sqlcgen.New(pool),
	}
}

// Save stores a note in PostgreSQL
func (r *NoteRepository) Save(ctx context.Context, note *domain.Note) error {
	params := sqlcgen.CreateNoteParams{
		ID:      note.ID,
		Content: note.Content,
		CreatedAt: pgtype.Timestamptz{
			Time:  note.CreatedAt,
			Valid: true,
		},
		ExpiresAt: pgtype.Timestamptz{
			Time:  note.ExpiresAt,
			Valid: true,
		},
	}

	return r.queries.CreateNote(ctx, params)
}

// FindByID retrieves a note by its ID
func (r *NoteRepository) FindByID(ctx context.Context, id string) (*domain.Note, error) {
	note, err := r.queries.GetNoteByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNoteNotFound
		}
		return nil, err
	}

	return &domain.Note{
		ID:        note.ID,
		Content:   note.Content,
		CreatedAt: note.CreatedAt.Time,
		ExpiresAt: note.ExpiresAt.Time,
	}, nil
}

// Delete removes a note by its ID
func (r *NoteRepository) Delete(ctx context.Context, id string) error {
	return r.queries.DeleteNote(ctx, id)
}

// DeleteExpired removes all expired notes and returns the count of deleted notes
func (r *NoteRepository) DeleteExpired(ctx context.Context) (int, error) {
	count, err := r.queries.DeleteExpiredNotes(ctx)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
