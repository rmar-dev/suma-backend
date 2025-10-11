package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rmar-dev/suma-backend/internal/controllers"
	"github.com/rmar-dev/suma-backend/internal/services"
)

// SetupMockRoutes configures all mock API routes
func SetupMockRoutes(router *gin.Engine) {
	// Set trusted proxies (for production, specify actual proxy IPs)
	router.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	
	// Serve static files
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	
	// Root route - API Documentation page
	router.GET("/", func(c *gin.Context) {
		c.File("./static/api-docs.html")
	})
	
	// Alternative documentation routes
	router.GET("/docs", func(c *gin.Context) {
		c.File("./static/api-docs.html")
	})
	router.GET("/swagger", func(c *gin.Context) {
		c.File("./static/api-docs.html")
	})
	
	// Setup CORS for development
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Initialize mock controllers
	authController := controllers.NewMockAuthController()
	userController := controllers.NewMockUserController()
	accountController := controllers.NewMockAccountController()
	transactionController := controllers.NewMockTransactionController()
	subscriptionController := controllers.NewMockSubscriptionController()
	budgetController := controllers.NewMockBudgetController()
	reportController := controllers.NewMockReportController()
	navigationController := controllers.NewNavigationController()
	
	// Initialize configuration service and controller
	configService := services.NewConfigService()
	configController := controllers.NewConfigController(configService)
	docsController := controllers.NewDocsController()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "finance-app-api",
			"mode":    "mock",
		})
	})

	// API v1 group
	v1 := router.Group("/api/v1")
	{
		// API Documentation
		v1.GET("/", docsController.GetAPIInfo)
		v1.GET("", docsController.GetAPIInfo)
		// Authentication routes (public)
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
			auth.POST("/refresh", authController.Refresh)
			auth.POST("/logout", authController.Logout)
		}

		// Protected routes (would normally require auth middleware)
		// For mock purposes, all routes are accessible

		// User profile routes
		v1.GET("/me", userController.GetProfile)
		v1.PUT("/me", userController.UpdateProfile)
		v1.GET("/settings", userController.GetSettings)
		v1.PUT("/settings", userController.UpdateSettings)

		// Navigation routes
		v1.GET("/navigation/menu", navigationController.GetMenuOptions)
		
		// Configuration routes
		configController.RegisterRoutes(v1)

		// Account routes
		accounts := v1.Group("/accounts")
		{
			accounts.GET("", accountController.List)
			accounts.POST("", accountController.Create)
			accounts.GET("/:id", accountController.Get)
			accounts.PUT("/:id", accountController.Update)
			accounts.DELETE("/:id", accountController.Delete)
			accounts.POST("/:id/sync", accountController.Sync)
		}

		// Transaction routes
		transactions := v1.Group("/transactions")
		{
			transactions.GET("", transactionController.List)
			transactions.GET("/:id", transactionController.Get)
			transactions.PUT("/:id", transactionController.Update)
			transactions.POST("/:id/categorize", transactionController.Categorize)
		}

		// Subscription routes
		subscriptions := v1.Group("/subscriptions")
		{
			subscriptions.GET("", subscriptionController.List)
			subscriptions.POST("", subscriptionController.Create)
			subscriptions.GET("/detect", subscriptionController.Detect)
			subscriptions.GET("/:id", subscriptionController.Get)
			subscriptions.PUT("/:id", subscriptionController.Update)
			subscriptions.DELETE("/:id", subscriptionController.Delete)
		}

		// Budget routes
		budgets := v1.Group("/budgets")
		{
			budgets.GET("", budgetController.List)
			budgets.POST("", budgetController.Create)
			budgets.GET("/:id", budgetController.Get)
			budgets.PUT("/:id", budgetController.Update)
			budgets.DELETE("/:id", budgetController.Delete)
		}

		// Report routes
		reports := v1.Group("/reports")
		{
			reports.GET("/summary", reportController.GetSummary)
			reports.GET("/spending", reportController.GetSpending)
			reports.GET("/subscriptions", reportController.GetSubscriptionCosts)
			reports.POST("/export", reportController.ExportData)
		}
	}
}

