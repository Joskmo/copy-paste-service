package service

import (
	"context"
	"strings"
	"time"

	"github.com/copy-paste-service/internal/domain"
	"github.com/copy-paste-service/internal/repository"
)

// NoteService defines the interface for note operations
type NoteService interface {
	// CreateNote creates a new note and returns its ID
	CreateNote(ctx context.Context, content string) (*domain.Note, error)

	// GetNote retrieves a note by its ID
	GetNote(ctx context.Context, id string) (*domain.Note, error)

	// StartCleanup starts the background cleanup process
	StartCleanup(ctx context.Context, interval time.Duration)
}

// noteService implements NoteService
type noteService struct {
	repo        repository.NoteRepository
	idGenerator IDGenerator
	noteTTL     time.Duration
}

// NewNoteService creates a new note service
func NewNoteService(repo repository.NoteRepository, idGenerator IDGenerator, ttl time.Duration) NoteService {
	return &noteService{
		repo:        repo,
		idGenerator: idGenerator,
		noteTTL:     ttl,
	}
}

// CreateNote creates a new note with the given content
func (s *noteService) CreateNote(ctx context.Context, content string) (*domain.Note, error) {
	content = strings.TrimSpace(content)
	if content == "" {
		return nil, domain.ErrEmptyContent
	}

	id := s.idGenerator.Generate()
	note := domain.NewNote(id, content, s.noteTTL)

	if err := s.repo.Save(ctx, note); err != nil {
		return nil, err
	}

	return note, nil
}

// GetNote retrieves a note by its ID
func (s *noteService) GetNote(ctx context.Context, id string) (*domain.Note, error) {
	return s.repo.FindByID(ctx, id)
}

// StartCleanup starts a background goroutine that periodically removes expired notes
func (s *noteService) StartCleanup(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				_, _ = s.repo.DeleteExpired(ctx)
			}
		}
	}()
}
