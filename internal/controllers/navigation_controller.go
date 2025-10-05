package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NavigationController struct{}

// NewNavigationController creates a new navigation controller
func NewNavigationController() *NavigationController {
	return &NavigationController{}
}

// MenuItem represents a navigation menu item
type MenuItem struct {
	ID          string     `json:"id"`
	Label       string     `json:"label"`
	Path        string     `json:"path"`
	Icon        string     `json:"icon"`
	Badge       *string    `json:"badge,omitempty"`
	Children    []MenuItem `json:"children,omitempty"`
	Permissions []string   `json:"permissions,omitempty"`
	IsActive    bool       `json:"is_active"`
	Order       int        `json:"order"`
}

// GetMenuOptions returns the available navigation menu options
func (nc *NavigationController) GetMenuOptions(c *gin.Context) {
	// Get user from context (set by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// For now, return static menu structure
	// In a real application, this would be dynamic based on user permissions
	menuItems := []MenuItem{
		{
			ID:       "dashboard",
			Label:    "Dashboard",
			Path:     "/dashboard",
			Icon:     "dashboard",
			IsActive: true,
			Order:    1,
		},
		{
			ID:       "accounts",
			Label:    "Accounts",
			Path:     "/accounts",
			Icon:     "account_balance",
			IsActive: true,
			Order:    2,
		},
		{
			ID:       "transactions",
			Label:    "Transactions",
			Path:     "/transactions",
			Icon:     "receipt_long",
			IsActive: true,
			Order:    3,
		},
		{
			ID:       "subscriptions",
			Label:    "Subscriptions",
			Path:     "/subscriptions",
			Icon:     "subscriptions",
			Badge:    stringPtr("12"),
			IsActive: true,
			Order:    4,
		},
		{
			ID:       "budgets",
			Label:    "Budgets",
			Path:     "/budgets",
			Icon:     "account_balance_wallet",
			IsActive: true,
			Order:    5,
		},
		{
			ID:       "reports",
			Label:    "Reports",
			Path:     "/reports",
			Icon:     "analytics",
			IsActive: true,
			Order:    6,
		},
		{
			ID:       "settings",
			Label:    "Settings",
			Path:     "/settings",
			Icon:     "settings",
			IsActive: true,
			Order:    7,
		},
	}

	// Add user-specific menu items
	userMenuItems := []MenuItem{
		{
			ID:       "profile",
			Label:    "Profile",
			Path:     "/profile",
			Icon:     "person",
			IsActive: true,
			Order:    8,
		},
		{
			ID:       "logout",
			Label:    "Logout",
			Path:     "/logout",
			Icon:     "logout",
			IsActive: true,
			Order:    9,
		},
	}

	// Combine main menu and user menu
	allMenuItems := append(menuItems, userMenuItems...)

	c.JSON(http.StatusOK, gin.H{
		"menu_items": allMenuItems,
		"user_id":    userID,
	})
}

// Helper function to create string pointer
func stringPtr(s string) *string {
	return &s
}
