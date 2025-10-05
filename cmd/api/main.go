package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/suma/finance-app-api/internal/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default configuration")
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Set Gin mode based on environment
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:5173",
		"http://localhost:3000",
	}
	if frontendURL := os.Getenv("FRONTEND_URL"); frontendURL != "" {
		config.AllowOrigins = append(config.AllowOrigins, frontendURL)
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	
	router.Use(cors.New(config))

	// Setup mock routes
	routes.SetupMockRoutes(router)

	// Start server
	fmt.Printf("🚀 Finance App API (Mock Mode) starting on port %s\n", port)
	fmt.Println("📝 API Documentation: http://localhost:" + port + "/api/v1")
	fmt.Println("🏥 Health Check: http://localhost:" + port + "/health")
	fmt.Println("\n📌 Available endpoints:")
	fmt.Println("  Auth:")
	fmt.Println("    POST   /api/v1/auth/register")
	fmt.Println("    POST   /api/v1/auth/login")
	fmt.Println("    POST   /api/v1/auth/refresh")
	fmt.Println("    POST   /api/v1/auth/logout")
	fmt.Println("  User:")
	fmt.Println("    GET    /api/v1/me")
	fmt.Println("    PUT    /api/v1/me")
	fmt.Println("    GET    /api/v1/settings")
	fmt.Println("    PUT    /api/v1/settings")
	fmt.Println("  Accounts:")
	fmt.Println("    GET    /api/v1/accounts")
	fmt.Println("    POST   /api/v1/accounts")
	fmt.Println("    GET    /api/v1/accounts/:id")
	fmt.Println("    PUT    /api/v1/accounts/:id")
	fmt.Println("    DELETE /api/v1/accounts/:id")
	fmt.Println("    POST   /api/v1/accounts/:id/sync")
	fmt.Println("  Transactions:")
	fmt.Println("    GET    /api/v1/transactions")
	fmt.Println("    GET    /api/v1/transactions/:id")
	fmt.Println("    PUT    /api/v1/transactions/:id")
	fmt.Println("    POST   /api/v1/transactions/:id/categorize")
	fmt.Println("  Subscriptions:")
	fmt.Println("    GET    /api/v1/subscriptions")
	fmt.Println("    POST   /api/v1/subscriptions")
	fmt.Println("    GET    /api/v1/subscriptions/detect")
	fmt.Println("    GET    /api/v1/subscriptions/:id")
	fmt.Println("    PUT    /api/v1/subscriptions/:id")
	fmt.Println("    DELETE /api/v1/subscriptions/:id")
	fmt.Println("  Budgets:")
	fmt.Println("    GET    /api/v1/budgets")
	fmt.Println("    POST   /api/v1/budgets")
	fmt.Println("    GET    /api/v1/budgets/:id")
	fmt.Println("    PUT    /api/v1/budgets/:id")
	fmt.Println("    DELETE /api/v1/budgets/:id")
	fmt.Println("  Reports:")
	fmt.Println("    GET    /api/v1/reports/summary")
	fmt.Println("    GET    /api/v1/reports/spending")
	fmt.Println("    GET    /api/v1/reports/subscriptions")
	fmt.Println("    POST   /api/v1/reports/export")
	fmt.Println("  Configuration:")
	fmt.Println("    GET    /api/v1/config/app")
	fmt.Println("    GET    /api/v1/config/features")
	fmt.Println("    GET    /api/v1/config/languages")
	fmt.Println("    GET    /api/v1/config/themes")
	fmt.Println("    GET    /api/v1/config/currencies")
	fmt.Println("    GET    /api/v1/config/categories")
	fmt.Println("    GET    /api/v1/config/user [Auth Required]")
	fmt.Println("    GET    /api/v1/config/menus [Auth Required]")
	fmt.Println("    GET    /api/v1/config/settings [Auth Required]")
	fmt.Println("  Navigation:")
	fmt.Println("    GET    /api/v1/navigation/menu")
	fmt.Println("\n✨ Mock mode enabled - returning sample data")
	
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
