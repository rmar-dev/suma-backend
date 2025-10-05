package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suma/finance-app-api/internal/services"
)

type ConfigController struct {
	configService *services.ConfigService
}

func NewConfigController(configService *services.ConfigService) *ConfigController {
	return &ConfigController{
		configService: configService,
	}
}

// RegisterRoutes registers all configuration routes
func (c *ConfigController) RegisterRoutes(router *gin.RouterGroup) {
	config := router.Group("/config")
	{
		// Public endpoints (no auth required)
		config.GET("/app", c.GetAppConfig)
		config.GET("/features", c.GetFeatures)
		config.GET("/languages", c.GetLanguages)
		config.GET("/themes", c.GetThemes)
		config.GET("/currencies", c.GetCurrencies)
		config.GET("/categories", c.GetCategories)
		
		// Authenticated endpoints (in mock mode, we'll skip auth for now)
		// In production, these would use: middleware.AuthMiddleware(jwtManager)
		config.GET("/user", c.GetUserConfig)
		config.GET("/menus", c.GetMenus)
		config.GET("/settings", c.GetSettings)
	}
}

// GetAppConfig returns the complete app configuration
// @Summary Get app configuration
// @Description Get complete app configuration including features, menus, settings
// @Tags Config
// @Accept json
// @Produce json
// @Param lang query string false "Language code (pt, en, es)"
// @Success 200 {object} services.AppConfig
// @Router /api/v1/config/app [get]
func (c *ConfigController) GetAppConfig(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", "pt")
	
	// For public access, assume free user
	config := c.configService.GetAppConfig("public", false, lang)
	
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    config,
	})
}

// GetFeatures returns available feature flags
// @Summary Get feature flags
// @Description Get list of available features and their status
// @Tags Config
// @Accept json
// @Produce json
// @Success 200 {object} services.FeatureFlags
// @Router /api/v1/config/features [get]
func (c *ConfigController) GetFeatures(ctx *gin.Context) {
	// Public features (free tier)
	features := c.configService.GetFeatureFlags("public", false)
	
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    features,
	})
}

// GetUserConfig returns user-specific configuration
// @Summary Get user configuration
// @Description Get configuration specific to authenticated user
// @Tags Config
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/config/user [get]
func (c *ConfigController) GetUserConfig(ctx *gin.Context) {
	// In mock mode, use a default user ID
	userID := ctx.GetString("user_id")
	if userID == "" {
		// Mock user ID for testing
		userID = "123e4567-e89b-12d3-a456-426614174000"
	}
	
	// TODO: Check if user is premium from database
	isPremium := false // This should come from user service
	
	config, err := c.configService.GetUserConfig(userID, isPremium)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to get user configuration",
		})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    config,
	})
}

// GetMenus returns navigation menus for authenticated user
// @Summary Get navigation menus
// @Description Get all navigation menus based on user permissions
// @Tags Config
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} services.MenuConfig
// @Router /api/v1/config/menus [get]
func (c *ConfigController) GetMenus(ctx *gin.Context) {
	// In mock mode, use a default user ID
	userID := ctx.GetString("user_id")
	if userID == "" {
		// Mock user ID for testing
		userID = "123e4567-e89b-12d3-a456-426614174000"
	}
	
	// TODO: Get user premium status from database
	isPremium := false
	
	features := c.configService.GetFeatureFlags("user", isPremium)
	
	menus := services.MenuConfig{
		MainMenu:     c.configService.GetMainMenu(features, isPremium),
		UserMenu:     c.configService.GetUserMenu(isPremium),
		SettingsMenu: c.configService.GetSettingsMenu(features),
		QuickActions: c.configService.GetQuickActions(features),
	}
	
	// Add notification badges if needed
	// TODO: Get notification counts from services
	for i, item := range menus.MainMenu {
		if item.ID == "subscriptions" {
			// Example: Add upcoming payment count
			menus.MainMenu[i].Badge = &services.Badge{
				Value: "3",
				Type:  "warning",
			}
		}
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    menus,
	})
}

// GetSettings returns app settings for authenticated user
// @Summary Get app settings
// @Description Get application settings
// @Tags Config
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} services.AppSettings
// @Router /api/v1/config/settings [get]
func (c *ConfigController) GetSettings(ctx *gin.Context) {
	settings := c.configService.GetAppSettings()
	
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    settings,
	})
}

// GetLanguages returns supported languages
// @Summary Get supported languages
// @Description Get list of supported languages
// @Tags Config
// @Accept json
// @Produce json
// @Success 200 {array} services.Language
// @Router /api/v1/config/languages [get]
func (c *ConfigController) GetLanguages(ctx *gin.Context) {
	languages := c.configService.GetLanguages()
	
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    languages,
	})
}

// GetThemes returns available themes
// @Summary Get available themes
// @Description Get list of available UI themes
// @Tags Config
// @Accept json
// @Produce json
// @Success 200 {array} services.Theme
// @Router /api/v1/config/themes [get]
func (c *ConfigController) GetThemes(ctx *gin.Context) {
	themes := c.configService.GetThemes()
	
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    themes,
	})
}

// GetCurrencies returns supported currencies
// @Summary Get supported currencies
// @Description Get list of supported currencies with exchange rates
// @Tags Config
// @Accept json
// @Produce json
// @Success 200 {array} services.Currency
// @Router /api/v1/config/currencies [get]
func (c *ConfigController) GetCurrencies(ctx *gin.Context) {
	currencies := c.configService.GetCurrencies()
	
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    currencies,
	})
}

// GetCategories returns transaction categories
// @Summary Get transaction categories
// @Description Get list of transaction categories
// @Tags Config
// @Accept json
// @Produce json
// @Param lang query string false "Language code (pt, en, es)"
// @Success 200 {array} services.Category
// @Router /api/v1/config/categories [get]
func (c *ConfigController) GetCategories(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", "pt")
	categories := c.configService.GetCategories(lang)
	
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    categories,
	})
}

// GetSubscriptionTypes returns billing cycle types
// @Summary Get subscription billing cycles
// @Description Get list of subscription billing cycle types
// @Tags Config
// @Accept json
// @Produce json
// @Param lang query string false "Language code (pt, en, es)"
// @Success 200 {array} services.SubscriptionType
// @Router /api/v1/config/subscription-types [get]
func (c *ConfigController) GetSubscriptionTypes(ctx *gin.Context) {
	lang := ctx.DefaultQuery("lang", "pt")
	types := c.configService.GetSubscriptionTypes(lang)
	
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    types,
	})
}