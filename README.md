# SUMA Finance App Backend

## Overview
Backend service for the SUMA Finance application, providing APIs for personal finance management, bank integration, and financial analysis.

## Technology Stack
- Go (Golang)
- PostgreSQL
- TrueLayer API Integration
- JWT Authentication
- Gin Web Framework

## Project Structure
```
├── cmd/
│   └── api/            # Application entrypoint
├── internal/
│   ├── models/         # Data models
│   ├── controllers/    # HTTP handlers
│   ├── middleware/     # HTTP middleware
│   ├── services/       # Business logic
│   └── repositories/   # Data access layer
├── scripts/           # Development and deployment scripts
└── migrations/        # Database migrations
```

## Getting Started

### Prerequisites
- Go 1.21 or later
- PostgreSQL 15 or later
- Docker (optional)

### Setup
1. Clone the repository
   ```bash
   git clone git@github.com:rmar-dev/suma-backend.git
   cd suma-backend
   ```

2. Install dependencies
   ```bash
   go mod download
   ```

3. Set up environment variables
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. Run database migrations
   ```bash
   ./scripts/run-migrations.ps1
   ```

5. Run the application
   ```bash
   go run cmd/api/main.go
   ```

### Development

#### Running Tests
```bash
go test ./...
```

#### Code Quality
```bash
# Run linter
golangci-lint run

# Run security checks
gosec ./...
```

## Contributing
1. Create a feature branch
2. Make your changes
3. Run tests and linting
4. Create a pull request

## License
Proprietary - All rights reserved