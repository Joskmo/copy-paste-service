package handler

import "time"

// CreateNoteRequest represents the request body for creating a note
type CreateNoteRequest struct {
	Content string `json:"content"`
}

// CreateNoteResponse represents the response after creating a note
type CreateNoteResponse struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	ExpiresAt time.Time `json:"expires_at"`
}

// GetNoteResponse represents the response when retrieving a note
type GetNoteResponse struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
