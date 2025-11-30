package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter creates and configures the HTTP router
func NewRouter(noteHandler *NoteHandler, healthHandler *HealthHandler, swaggerHandler *SwaggerHandler) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(corsMiddleware)

	// Health check
	r.Get("/health", healthHandler.Health)

	// Swagger UI
	r.Get("/swagger", swaggerHandler.ServeUI)
	r.Get("/swagger/", swaggerHandler.ServeUI)
	r.Get("/swagger/openapi.yaml", swaggerHandler.ServeSpec)

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Route("/notes", func(r chi.Router) {
			r.Post("/", noteHandler.CreateNote)
			r.Get("/{id}", noteHandler.GetNote)
			r.Get("/{id}/raw", noteHandler.GetNoteRaw)
		})
	})

	return r
}

// corsMiddleware handles CORS headers
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
