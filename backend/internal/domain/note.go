package domain

import (
	"time"
)

// Note represents a text note with expiration
type Note struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// IsExpired checks if the note has expired
func (n *Note) IsExpired() bool {
	return time.Now().After(n.ExpiresAt)
}

// NewNote creates a new note with the given content and TTL
func NewNote(id, content string, ttl time.Duration) *Note {
	now := time.Now()
	return &Note{
		ID:        id,
		Content:   content,
		CreatedAt: now,
		ExpiresAt: now.Add(ttl),
	}
}
