.PHONY: build run test clean docker-build docker-run docker-stop sqlc migrate

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
BINARY_NAME=server

# Database
DB_DSN=postgres://postgres:postgres@localhost:5432/copypaste?sslmode=disable

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/server

# Run the application
run:
	$(GOCMD) run ./cmd/server

# Run tests
test:
	$(GOTEST) -v ./...

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Download dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Generate sqlc code
sqlc:
	sqlc generate

# Run migrations up (requires goose CLI)
migrate-up:
	goose -dir migrations postgres "$(DB_DSN)" up

# Run migrations down
migrate-down:
	goose -dir migrations postgres "$(DB_DSN)" down

# Show migration status
migrate-status:
	goose -dir migrations postgres "$(DB_DSN)" status

# Create new migration
migrate-create:
	@read -p "Migration name: " name; \
	goose -dir migrations create $$name sql

# Build Docker images
docker-build:
	docker compose build

# Run with Docker Compose (full stack)
docker-run:
	docker compose up -d

# Stop Docker containers
docker-stop:
	docker compose down

# View logs
docker-logs:
	docker compose logs -f

# View app logs only
docker-logs-app:
	docker compose logs -f app

# Rebuild and restart
docker-restart: docker-stop docker-build docker-run

# Start only PostgreSQL (for local development)
postgres-start:
	docker compose up -d postgres

# Run migrations only (requires postgres running)
docker-migrate:
	docker compose up migrations

# Stop PostgreSQL
postgres-stop:
	docker compose stop postgres

# Full reset: stop, remove volumes, rebuild
docker-reset:
	docker compose down -v
	docker compose up --build -d
