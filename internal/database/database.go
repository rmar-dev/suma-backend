package database

import (
	"fmt"
	"log"
	"time"

	"github.com/rmar-dev/suma-backend/internal/config"
	"github.com/rmar-dev/suma-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Initialize sets up the database connection
func Initialize(cfg *config.DatabaseConfig) error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	var gormConfig *gorm.Config

	// Configure logging based on environment
	if cfg.SSLMode == "disable" { // Development mode
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	} else {
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		}
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Get the underlying SQL database to configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established successfully")
	return nil
}

// AutoMigrate runs the auto migration for all models
func AutoMigrate() error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	// Enable UUID extension in PostgreSQL
	if err := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		log.Printf("Warning: Could not create UUID extension: %v", err)
	}

	// Run auto migration for all models
	err := DB.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Transaction{},
		&models.Subscription{},
	)

	if err != nil {
		return fmt.Errorf("failed to run auto migration: %w", err)
	}

	log.Println("Database migration completed successfully")
	return nil
}

// Close closes the database connection
func Close() error {
	if DB == nil {
		return nil
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}

// Health checks the database health
func Health() error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}
