package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create test router
	router := gin.New()
	if err := SetupMockRoutes(router); err != nil {
		t.Fatal("Failed to setup routes:", err)
	}

	// Create test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)

	// Perform request
	router.ServeHTTP(w, req)

	// Check response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "healthy")
}

func TestMockEndpoints(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create test router
	router := gin.New()
	if err := SetupMockRoutes(router); err != nil {
		t.Fatal("Failed to setup routes:", err)
	}

	// Test cases
	testCases := []struct {
		name   string
		method string
		path   string
	}{
		{"Auth Register", "POST", "/api/v1/auth/register"},
		{"Auth Login", "POST", "/api/v1/auth/login"},
		{"Get User", "GET", "/api/v1/me"},
		{"Get Accounts", "GET", "/api/v1/accounts"},
		{"Get Transactions", "GET", "/api/v1/transactions"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tc.method, tc.path, nil)
			router.ServeHTTP(w, req)

			if tc.name == "Auth Register" {
				assert.Equal(t, http.StatusCreated, w.Code, "Register should return Created status")
				assert.Contains(t, w.Body.String(), "User registered successfully")
			} else {
				assert.Equal(t, http.StatusOK, w.Code, "Response code should be OK")
			}

			// Verify response contains required fields
			switch tc.name {
			case "Auth Login":
				assert.Contains(t, w.Body.String(), "Login successful")
				assert.Contains(t, w.Body.String(), "tokens")
				assert.Contains(t, w.Body.String(), "user")
			case "Get User":
				assert.Contains(t, w.Body.String(), "user")
			case "Get Accounts":
				assert.Contains(t, w.Body.String(), "accounts")
			case "Get Transactions":
				assert.Contains(t, w.Body.String(), "transactions")
			}
		})
	}
}
