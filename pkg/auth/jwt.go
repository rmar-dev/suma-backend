package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

// Claims represents the JWT claims
type Claims struct {
	UserID    uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	TokenType TokenType `json:"token_type"`
	jwt.RegisteredClaims
}

// JWTManager handles JWT operations
type JWTManager struct {
	accessSecret  string
	refreshSecret string
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

// NewJWTManager creates a new JWT manager
func NewJWTManager(accessSecret, refreshSecret string, accessExpiry, refreshExpiry time.Duration) *JWTManager {
	return &JWTManager{
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
	}
}

// GenerateTokenPair generates both access and refresh tokens
func (j *JWTManager) GenerateTokenPair(userID uuid.UUID, email string) (accessToken, refreshToken string, err error) {
	// Generate access token
	accessToken, err = j.GenerateToken(userID, email, AccessToken)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}

	// Generate refresh token
	refreshToken, err = j.GenerateToken(userID, email, RefreshToken)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}

// GenerateToken generates a JWT token
func (j *JWTManager) GenerateToken(userID uuid.UUID, email string, tokenType TokenType) (string, error) {
	var secret string
	var expiry time.Duration

	switch tokenType {
	case AccessToken:
		secret = j.accessSecret
		expiry = j.accessExpiry
	case RefreshToken:
		secret = j.refreshSecret
		expiry = j.refreshExpiry
	default:
		return "", errors.New("invalid token type")
	}

	claims := &Claims{
		UserID:    userID,
		Email:     email,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "finance-app-api",
			Subject:   userID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateToken validates and parses a JWT token
func (j *JWTManager) ValidateToken(tokenString string, tokenType TokenType) (*Claims, error) {
	var secret string

	switch tokenType {
	case AccessToken:
		secret = j.accessSecret
	case RefreshToken:
		secret = j.refreshSecret
	default:
		return nil, errors.New("invalid token type")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Verify token type
	if claims.TokenType != tokenType {
		return nil, errors.New("token type mismatch")
	}

	return claims, nil
}

// RefreshAccessToken generates a new access token using a refresh token
func (j *JWTManager) RefreshAccessToken(refreshToken string) (string, error) {
	// Validate the refresh token
	claims, err := j.ValidateToken(refreshToken, RefreshToken)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token: %w", err)
	}

	// Generate a new access token
	accessToken, err := j.GenerateToken(claims.UserID, claims.Email, AccessToken)
	if err != nil {
		return "", fmt.Errorf("failed to generate new access token: %w", err)
	}

	return accessToken, nil
}

// GetUserIDFromToken extracts the user ID from a token
func (j *JWTManager) GetUserIDFromToken(tokenString string, tokenType TokenType) (uuid.UUID, error) {
	claims, err := j.ValidateToken(tokenString, tokenType)
	if err != nil {
		return uuid.Nil, err
	}
	return claims.UserID, nil
}

// TokenResponse represents the token response structure
type TokenResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int       `json:"expires_in"`
	ExpiresAt    time.Time `json:"expires_at"`
}

// NewTokenResponse creates a new token response
func NewTokenResponse(accessToken, refreshToken string, expiresIn time.Duration) TokenResponse {
	return TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int(expiresIn.Seconds()),
		ExpiresAt:    time.Now().Add(expiresIn),
	}
}