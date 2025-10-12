package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DocsController handles API documentation
type DocsController struct{}

// NewDocsController creates a new docs controller
func NewDocsController() *DocsController {
	return &DocsController{}
}

// GetAPIInfo returns API information
func (d *DocsController) GetAPIInfo(c *gin.Context) {
	info := gin.H{
		"name":        "SUMA Finance API",
		"version":     "1.0.0",
		"description": "Personal Finance Management Platform API",
		"mode":        "mock",
		"endpoints": gin.H{
			"auth": []gin.H{
				{"method": "POST", "path": "/api/v1/auth/register", "description": "Register new user"},
				{"method": "POST", "path": "/api/v1/auth/login", "description": "Login user"},
				{"method": "POST", "path": "/api/v1/auth/refresh", "description": "Refresh access token"},
				{"method": "POST", "path": "/api/v1/auth/logout", "description": "Logout user"},
			},
			"config": []gin.H{
				{"method": "GET", "path": "/api/v1/config/app", "description": "Get app configuration"},
				{"method": "GET", "path": "/api/v1/config/features", "description": "Get feature flags"},
				{"method": "GET", "path": "/api/v1/config/menus", "description": "Get navigation menus"},
				{"method": "GET", "path": "/api/v1/config/languages", "description": "Get supported languages"},
				{"method": "GET", "path": "/api/v1/config/themes", "description": "Get available themes"},
				{"method": "GET", "path": "/api/v1/config/currencies", "description": "Get supported currencies"},
				{"method": "GET", "path": "/api/v1/config/categories", "description": "Get transaction categories"},
			},
			"subscriptions": []gin.H{
				{"method": "GET", "path": "/api/v1/subscriptions", "description": "List all subscriptions"},
				{"method": "POST", "path": "/api/v1/subscriptions", "description": "Create subscription"},
				{"method": "GET", "path": "/api/v1/subscriptions/detect", "description": "Detect recurring payments"},
				{"method": "GET", "path": "/api/v1/subscriptions/:id", "description": "Get subscription details"},
				{"method": "PUT", "path": "/api/v1/subscriptions/:id", "description": "Update subscription"},
				{"method": "DELETE", "path": "/api/v1/subscriptions/:id", "description": "Delete subscription"},
			},
			"transactions": []gin.H{
				{"method": "GET", "path": "/api/v1/transactions", "description": "List transactions"},
				{"method": "GET", "path": "/api/v1/transactions/:id", "description": "Get transaction details"},
				{"method": "PUT", "path": "/api/v1/transactions/:id", "description": "Update transaction"},
				{"method": "POST", "path": "/api/v1/transactions/:id/categorize", "description": "Categorize transaction"},
			},
			"accounts": []gin.H{
				{"method": "GET", "path": "/api/v1/accounts", "description": "List bank accounts"},
				{"method": "POST", "path": "/api/v1/accounts", "description": "Add bank account"},
				{"method": "GET", "path": "/api/v1/accounts/:id", "description": "Get account details"},
				{"method": "PUT", "path": "/api/v1/accounts/:id", "description": "Update account"},
				{"method": "DELETE", "path": "/api/v1/accounts/:id", "description": "Delete account"},
				{"method": "POST", "path": "/api/v1/accounts/:id/sync", "description": "Sync account transactions"},
			},
		},
		"features": []string{
			"Subscription Detection Algorithm",
			"Multi-language Support (PT, EN, ES)",
			"Dynamic Menu System",
			"Feature Flags",
			"Dark/Light Theme",
			"GDPR Compliant",
			"PSD2 Ready",
		},
	}

	c.JSON(http.StatusOK, info)
}
