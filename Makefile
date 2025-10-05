# Variables
APP_NAME=finance-app-api
MAIN_PATH=cmd/api/main.go
BUILD_PATH=build/
BINARY_NAME=$(APP_NAME)

# Colors for output
GREEN=\033[0;32m
RED=\033[0;31m
NC=\033[0m # No Color

.PHONY: help
help: ## Display this help screen
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'

.PHONY: install
install: ## Install dependencies
	@echo "$(GREEN)Installing dependencies...$(NC)"
	go mod download
	go mod tidy

.PHONY: build
build: ## Build the application
	@echo "$(GREEN)Building application...$(NC)"
	go build -o $(BUILD_PATH)$(BINARY_NAME) $(MAIN_PATH)

.PHONY: run
run: ## Run the application
	@echo "$(GREEN)Running application...$(NC)"
	go run $(MAIN_PATH)

.PHONY: dev
dev: ## Run the application with hot reload (requires air)
	@echo "$(GREEN)Running in development mode...$(NC)"
	air

.PHONY: test
test: ## Run tests
	@echo "$(GREEN)Running tests...$(NC)"
	go test -v ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage
	@echo "$(GREEN)Running tests with coverage...$(NC)"
	go test -v -cover -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: lint
lint: ## Run linter (requires golangci-lint)
	@echo "$(GREEN)Running linter...$(NC)"
	golangci-lint run

.PHONY: fmt
fmt: ## Format code
	@echo "$(GREEN)Formatting code...$(NC)"
	go fmt ./...
	gofmt -s -w .

.PHONY: clean
clean: ## Clean build files
	@echo "$(GREEN)Cleaning build files...$(NC)"
	rm -rf $(BUILD_PATH)
	rm -f coverage.out coverage.html

.PHONY: migrate-up
migrate-up: ## Run database migrations up
	@echo "$(GREEN)Running migrations up...$(NC)"
	migrate -path migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

.PHONY: migrate-down
migrate-down: ## Run database migrations down
	@echo "$(GREEN)Running migrations down...$(NC)"
	migrate -path migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

.PHONY: migrate-create
migrate-create: ## Create a new migration file (usage: make migrate-create name=migration_name)
	@echo "$(GREEN)Creating migration: $(name)$(NC)"
	migrate create -ext sql -dir migrations -seq $(name)

.PHONY: docker-build
docker-build: ## Build Docker image
	@echo "$(GREEN)Building Docker image...$(NC)"
	docker build -t $(APP_NAME):latest .

.PHONY: docker-run
docker-run: ## Run Docker container
	@echo "$(GREEN)Running Docker container...$(NC)"
	docker run -p 8080:8080 --env-file .env $(APP_NAME):latest

.PHONY: install-tools
install-tools: ## Install development tools
	@echo "$(GREEN)Installing development tools...$(NC)"
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: setup
setup: install install-tools ## Setup the project (install dependencies and tools)
	@echo "$(GREEN)Project setup complete!$(NC)"
	@echo "Create a .env file based on .env.example and configure your environment variables."
	@echo "Run 'make run' to start the application."