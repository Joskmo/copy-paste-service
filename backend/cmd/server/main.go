package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/copy-paste-service/internal/config"
	"github.com/copy-paste-service/internal/database"
	"github.com/copy-paste-service/internal/handler"
	"github.com/copy-paste-service/internal/repository/postgres"
	"github.com/copy-paste-service/internal/service"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Create context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Connect to PostgreSQL
	log.Println("Connecting to PostgreSQL...")
	pool, err := database.NewPostgresPool(ctx, &cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()
	log.Println("Connected to PostgreSQL successfully")

	// Initialize dependencies
	noteRepo := postgres.NewNoteRepository(pool)
	idGenerator := service.NewWordBasedIDGenerator()
	noteService := service.NewNoteService(noteRepo, idGenerator, cfg.Note.TTL)

	// Initialize handlers
	noteHandler := handler.NewNoteHandler(noteService, cfg.Server.BaseURL)
	healthHandler := handler.NewHealthHandler()
	swaggerHandler := handler.NewSwaggerHandler("api/openapi.yaml")

	// Setup router
	router := handler.NewRouter(noteHandler, healthHandler, swaggerHandler)

	// Create server
	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start cleanup goroutine
	noteService.StartCleanup(ctx, cfg.Cleanup.Interval)

	// Start server in goroutine
	go func() {
		log.Printf("Starting server on port %s", cfg.Server.Port)
		log.Printf("Base URL: %s", cfg.Server.BaseURL)
		log.Printf("Swagger UI: http://localhost:%s/swagger/", cfg.Server.Port)
		log.Printf("Note TTL: %v", cfg.Note.TTL)
		log.Printf("Cleanup interval: %v", cfg.Cleanup.Interval)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Cancel context to stop cleanup goroutine
	cancel()

	// Graceful shutdown with timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}
