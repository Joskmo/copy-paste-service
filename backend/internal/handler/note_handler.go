package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/copy-paste-service/internal/domain"
	"github.com/copy-paste-service/internal/service"
)

// NoteHandler handles HTTP requests for notes
type NoteHandler struct {
	noteService service.NoteService
	baseURL     string
}

// NewNoteHandler creates a new note handler
func NewNoteHandler(noteService service.NoteService, baseURL string) *NoteHandler {
	return &NoteHandler{
		noteService: noteService,
		baseURL:     baseURL,
	}
}

// CreateNote handles POST /api/notes
func (h *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req CreateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid_request", "Invalid JSON body")
		return
	}

	note, err := h.noteService.CreateNote(r.Context(), req.Content)
	if err != nil {
		if errors.Is(err, domain.ErrEmptyContent) {
			h.respondError(w, http.StatusBadRequest, "empty_content", "Note content cannot be empty")
			return
		}
		h.respondError(w, http.StatusInternalServerError, "internal_error", "Failed to create note")
		return
	}

	response := CreateNoteResponse{
		ID:        note.ID,
		URL:       h.baseURL + "/" + note.ID,
		ExpiresAt: note.ExpiresAt,
	}

	h.respondJSON(w, http.StatusCreated, response)
}

// GetNote handles GET /api/notes/{id}
func (h *NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.respondError(w, http.StatusBadRequest, "missing_id", "Note ID is required")
		return
	}

	note, err := h.noteService.GetNote(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNoteNotFound) || errors.Is(err, domain.ErrNoteExpired) {
			h.respondError(w, http.StatusNotFound, "not_found", "Note not found or has expired")
			return
		}
		h.respondError(w, http.StatusInternalServerError, "internal_error", "Failed to retrieve note")
		return
	}

	response := GetNoteResponse{
		ID:        note.ID,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		ExpiresAt: note.ExpiresAt,
	}

	h.respondJSON(w, http.StatusOK, response)
}

// GetNoteRaw handles GET /api/notes/{id}/raw - returns plain text content
func (h *NoteHandler) GetNoteRaw(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Note ID is required", http.StatusBadRequest)
		return
	}

	note, err := h.noteService.GetNote(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNoteNotFound) || errors.Is(err, domain.ErrNoteExpired) {
			http.Error(w, "Note not found or has expired", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(note.Content))
}

// respondJSON sends a JSON response
func (h *NoteHandler) respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// respondError sends an error response
func (h *NoteHandler) respondError(w http.ResponseWriter, status int, errorCode, message string) {
	h.respondJSON(w, status, ErrorResponse{
		Error:   errorCode,
		Message: message,
	})
}
