package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rmar-dev/suma-backend/pkg/auth"
)

// AuthMiddleware creates a JWT authentication middleware
func AuthMiddleware(jwtManager *auth.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Check Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		// Validate token
		tokenString := parts[1]
		claims, err := jwtManager.ValidateToken(tokenString, auth.AccessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("userID", claims.UserID)
		c.Set("userEmail", claims.Email)
		c.Set("claims", claims)

		// Continue to next handler
		c.Next()
	}
}

// OptionalAuthMiddleware creates an optional JWT authentication middleware
// It validates the token if present but doesn't require it
func OptionalAuthMiddleware(jwtManager *auth.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// No token provided, continue without authentication
			c.Next()
			return
		}

		// Check Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			// Invalid format, continue without authentication
			c.Next()
			return
		}

		// Validate token
		tokenString := parts[1]
		claims, err := jwtManager.ValidateToken(tokenString, auth.AccessToken)
		if err != nil {
			// Invalid token, continue without authentication
			c.Next()
			return
		}

		// Set user information in context
		c.Set("userID", claims.UserID)
		c.Set("userEmail", claims.Email)
		c.Set("claims", claims)
		c.Set("authenticated", true)

		// Continue to next handler
		c.Next()
	}
}
