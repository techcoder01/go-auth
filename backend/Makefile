# Makefile for Go backend development

# Variables
APP_NAME := backend
BIN_DIR := bin
DB_NAME := goauth
MIGRATIONS_DIR := migrations

# Go variables
GO := go
GO_BUILD := $(GO) build
GO_RUN := $(GO) run
GO_TEST := $(GO) test
GO_MOD := $(GO) mod
GO_LINT := golangci-lint

.PHONY: all build run test clean migrate-up migrate-down lint deps setup

# Build the application
build:
	@echo "Building application..."
	@mkdir -p $(BIN_DIR)
	@$(GO_BUILD) -o $(BIN_DIR)/$(APP_NAME) ./cmd/server
	@echo "Build complete!"

# Run the application
run:
	@echo "Starting application..."
	@$(GO_RUN) cmd/server/main.go

# Run tests
test:
	@echo "Running tests..."
	@$(GO_TEST) -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)
	@$(GO) clean

# Database migrations
migrate-up:
	@echo "Applying database migrations..."
	@migrate -path $(MIGRATIONS_DIR) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migrate-down:
	@echo "Reverting database migrations..."
	@migrate -path $(MIGRATIONS_DIR) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down 1

# Lint the code
lint:
	@echo "Running linter..."
	@$(GO_LINT) run

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@$(GO_MOD) download
	@$(GO) install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Setup development environment
setup: deps migrate-up

# Default target
.DEFAULT_GOAL := run
