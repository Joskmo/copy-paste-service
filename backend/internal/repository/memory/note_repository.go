package memory

import (
	"context"
	"sync"
	"time"

	"github.com/copy-paste-service/internal/domain"
)

// NoteRepository implements in-memory storage for notes
type NoteRepository struct {
	mu    sync.RWMutex
	notes map[string]*domain.Note
}

// NewNoteRepository creates a new in-memory note repository
func NewNoteRepository() *NoteRepository {
	return &NoteRepository{
		notes: make(map[string]*domain.Note),
	}
}

// Save stores a note in memory
func (r *NoteRepository) Save(ctx context.Context, note *domain.Note) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.notes[note.ID] = note
	return nil
}

// FindByID retrieves a note by its ID
func (r *NoteRepository) FindByID(ctx context.Context, id string) (*domain.Note, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	note, exists := r.notes[id]
	if !exists {
		return nil, domain.ErrNoteNotFound
	}

	if note.IsExpired() {
		return nil, domain.ErrNoteExpired
	}

	return note, nil
}

// Delete removes a note by its ID
func (r *NoteRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.notes, id)
	return nil
}

// DeleteExpired removes all expired notes and returns the count of deleted notes
func (r *NoteRepository) DeleteExpired(ctx context.Context) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	deleted := 0

	for id, note := range r.notes {
		if now.After(note.ExpiresAt) {
			delete(r.notes, id)
			deleted++
		}
	}

	return deleted, nil
}

