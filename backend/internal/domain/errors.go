package domain

import "errors"

var (
	// ErrNoteNotFound is returned when a note is not found
	ErrNoteNotFound = errors.New("note not found")

	// ErrNoteExpired is returned when a note has expired
	ErrNoteExpired = errors.New("note has expired")

	// ErrEmptyContent is returned when note content is empty
	ErrEmptyContent = errors.New("note content cannot be empty")
)
