package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	validatorlib "github.com/go-playground/validator/v10"
	"github.com/rmar-dev/suma-backend/internal/database"
	"github.com/rmar-dev/suma-backend/internal/models"
	"github.com/rmar-dev/suma-backend/pkg/auth"
	"gorm.io/gorm"
)

type AuthController struct {
	db         *gorm.DB
	jwtManager *auth.JWTManager
	validator  *validatorlib.Validate
}

// NewAuthController creates a new authentication controller
func NewAuthController(jwtManager *auth.JWTManager) *AuthController {
	return &AuthController{
		db:         database.GetDB(),
		jwtManager: jwtManager,
		validator:  validatorlib.New(),
	}
}

// RegisterRequest represents the registration request body
type RegisterRequest struct {
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required,min=8"`
	PasswordConfirm  string `json:"password_confirm" validate:"required,eqfield=Password"`
	FirstName        string `json:"first_name" validate:"required,min=2,max=50"`
	LastName         string `json:"last_name" validate:"required,min=2,max=50"`
	PhoneNumber      string `json:"phone_number,omitempty"`
	MarketingConsent bool   `json:"marketing_consent"`
	TermsAccepted    bool   `json:"terms_accepted" validate:"required,eq=true"`
	PrivacyAccepted  bool   `json:"privacy_accepted" validate:"required,eq=true"`
	GDPRConsent      bool   `json:"gdpr_consent" validate:"required,eq=true"`
}

// LoginRequest represents the login request body
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// RefreshRequest represents the token refresh request body
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// Register handles user registration
func (ac *AuthController) Register(c *gin.Context) {
	var req RegisterRequest
	
	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate request
	if err := ac.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user exists
	var existingUser models.User
	if err := ac.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Create new user
	now := time.Now()
	user := models.User{
		Email:             req.Email,
		Password:          req.Password, // Will be hashed in BeforeCreate hook
		FirstName:         req.FirstName,
		LastName:          req.LastName,
		PhoneNumber:       req.PhoneNumber,
		MarketingConsent:  req.MarketingConsent,
		TermsAcceptedAt:   &now,
		PrivacyAcceptedAt: &now,
		GDPRConsentAt:     &now,
	}

	// Save user to database
	if err := ac.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate tokens
	accessToken, refreshToken, err := ac.jwtManager.GenerateTokenPair(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user.ToResponse(),
		"tokens":  auth.NewTokenResponse(accessToken, refreshToken, time.Hour),
	})
}

// Login handles user login
func (ac *AuthController) Login(c *gin.Context) {
	var req LoginRequest

	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate request
	if err := ac.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user
	var user models.User
	if err := ac.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	if !user.CheckPassword(req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check if user is active
	if !user.Active {
		c.JSON(http.StatusForbidden, gin.H{"error": "Account is deactivated"})
		return
	}

	// Update last login
	now := time.Now()
	user.LastLogin = &now
	ac.db.Save(&user)

	// Generate tokens
	accessToken, refreshToken, err := ac.jwtManager.GenerateTokenPair(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user.ToResponse(),
		"tokens":  auth.NewTokenResponse(accessToken, refreshToken, time.Hour),
	})
}

// RefreshToken handles token refresh
func (ac *AuthController) RefreshToken(c *gin.Context) {
	var req RefreshRequest

	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate request
	if err := ac.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate refresh token and get claims
	claims, err := ac.jwtManager.ValidateToken(req.RefreshToken, auth.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Check if user still exists and is active
	var user models.User
	if err := ac.db.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if !user.Active {
		c.JSON(http.StatusForbidden, gin.H{"error": "Account is deactivated"})
		return
	}

	// Generate new tokens
	accessToken, newRefreshToken, err := ac.jwtManager.GenerateTokenPair(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	// Return new tokens
	c.JSON(http.StatusOK, gin.H{
		"message": "Token refreshed successfully",
		"tokens":  auth.NewTokenResponse(accessToken, newRefreshToken, time.Hour),
	})
}

// Logout handles user logout (client-side token removal)
func (ac *AuthController) Logout(c *gin.Context) {
	// In a JWT-based system, logout is typically handled client-side
	// by removing the tokens. Optionally, you can implement a token
	// blacklist here if needed.
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful. Please remove tokens from client storage.",
	})
}

// GetMe returns the current authenticated user
func (ac *AuthController) GetMe(c *gin.Context) {
	// Get user from context (set by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var user models.User
	if err := ac.db.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user.ToResponse(),
	})
}
