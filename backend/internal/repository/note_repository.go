package repository

import (
	"context"

	"github.com/copy-paste-service/internal/domain"
)

// NoteRepository defines the interface for note storage operations
type NoteRepository interface {
	// Save stores a note
	Save(ctx context.Context, note *domain.Note) error

	// FindByID retrieves a note by its ID
	FindByID(ctx context.Context, id string) (*domain.Note, error)

	// Delete removes a note by its ID
	Delete(ctx context.Context, id string) error

	// DeleteExpired removes all expired notes
	DeleteExpired(ctx context.Context) (int, error)
}
