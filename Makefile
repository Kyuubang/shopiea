include .env

.PHONY: run build migrate seed docker-up docker-down clean help

# Default target
help:
	@echo "Available commands:"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application"
	@echo "  make migrate      - Run database migrations"
	@echo "  make seed         - Seed the database with default data"
	@echo "  make docker-up    - Start application with Docker Compose"
	@echo "  make docker-down  - Stop Docker Compose containers"
	@echo "  make clean        - Remove build artifacts"

# Build the application
build:
	go build -o bin/shopiea

# Run the application
run:
	go run main.go

# Run database migrations
migrate:
	go run main.go -migrate

# Seed the database
seed:
	go run main.go -seed

# Start Docker containers
docker-up:
	docker-compose up -d

# Stop Docker containers
docker-down:
	docker-compose down

# Clean build artifacts
clean:
	rm -rf bin/