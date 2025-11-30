package config

import (
	"os"
	"strconv"
	"time"

	"github.com/copy-paste-service/internal/database"
)

// Config holds all application configuration
type Config struct {
	Server   ServerConfig
	Database database.PostgresConfig
	Note     NoteConfig
	Cleanup  CleanupConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port    string
	BaseURL string
}

// NoteConfig holds note-related configuration
type NoteConfig struct {
	TTL time.Duration
}

// CleanupConfig holds cleanup-related configuration
type CleanupConfig struct {
	Interval time.Duration
}

// Load loads configuration from environment variables with defaults
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:    getEnv("SERVER_PORT", "8080"),
			BaseURL: getEnv("BASE_URL", "http://localhost:8080/api/notes"),
		},
		Database: database.PostgresConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Database: getEnv("DB_NAME", "copypaste"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Note: NoteConfig{
			TTL: getDurationEnv("NOTE_TTL_HOURS", 3) * time.Hour,
		},
		Cleanup: CleanupConfig{
			Interval: getDurationEnv("CLEANUP_INTERVAL_MINUTES", 5) * time.Minute,
		},
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getDurationEnv gets a duration from environment or returns default
func getDurationEnv(key string, defaultValue int) time.Duration {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return time.Duration(intVal)
		}
	}
	return time.Duration(defaultValue)
}
