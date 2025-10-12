# SUMA Finance App Backend

## Overview
SUMA Finance is a comprehensive personal finance management platform that helps users track their spending, manage subscriptions, and gain insights into their financial health. The backend service provides secure APIs for bank integration, transaction analysis, and financial management.

### Key Features
- 🏦 Bank Integration via TrueLayer
- 💰 Transaction Tracking & Categorization
- 📊 Spending Analytics & Insights
- 📅 Subscription Management & Detection
- 📱 Multi-device Support
- 🔒 Secure Authentication & Authorization

## System Architecture

### Technology Stack
- **Language:** Go 1.21+
- **Framework:** Gin Web Framework
- **Database:** PostgreSQL 15
- **Authentication:** JWT with refresh tokens
- **Bank Integration:** TrueLayer API
- **Documentation:** OpenAPI/Swagger
- **Testing:** Go testing framework with testify
- **CI/CD:** GitHub Actions

### Core Components
- **Authentication Service:** User management and security
- **Transaction Service:** Transaction processing and categorization
- **Account Service:** Bank account management and syncing
- **Analytics Service:** Financial analysis and insights
- **Subscription Service:** Recurring payment detection and management

## Current Status
The project is under active development. Current implementation includes:

✅ Basic project structure
✅ CI/CD setup
✅ Development guidelines
🚧 Database migrations (in progress)
📋 Feature planning complete

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