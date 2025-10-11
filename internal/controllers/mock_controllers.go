package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rmar-dev/suma-backend/internal/mocks"
	"github.com/google/uuid"
	"time"
)

// MockAuthController handles authentication endpoints with mock data
type MockAuthController struct{}

func NewMockAuthController() *MockAuthController {
	return &MockAuthController{}
}

func (c *MockAuthController) Register(ctx *gin.Context) {
	user := mocks.GenerateMockUser()
	tokens := map[string]interface{}{
		"access_token":  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.mock_access_token." + uuid.New().String(),
		"refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.mock_refresh_token." + uuid.New().String(),
		"token_type":    "Bearer",
		"expires_in":    3600,
		"expires_at":    time.Now().Add(time.Hour),
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
		"tokens":  tokens,
	})
}

func (c *MockAuthController) Login(ctx *gin.Context) {
	user := mocks.GenerateMockUser()
	tokens := map[string]interface{}{
		"access_token":  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.mock_access_token." + uuid.New().String(),
		"refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.mock_refresh_token." + uuid.New().String(),
		"token_type":    "Bearer",
		"expires_in":    3600,
		"expires_at":    time.Now().Add(time.Hour),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
		"tokens":  tokens,
	})
}

func (c *MockAuthController) Refresh(ctx *gin.Context) {
	tokens := map[string]interface{}{
		"access_token":  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.mock_access_token." + uuid.New().String(),
		"refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.mock_refresh_token." + uuid.New().String(),
		"token_type":    "Bearer",
		"expires_in":    3600,
		"expires_at":    time.Now().Add(time.Hour),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Token refreshed successfully",
		"tokens":  tokens,
	})
}

func (c *MockAuthController) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logout successful. Please remove tokens from client storage.",
	})
}

// MockUserController handles user profile endpoints with mock data
type MockUserController struct{}

func NewMockUserController() *MockUserController {
	return &MockUserController{}
}

func (c *MockUserController) GetProfile(ctx *gin.Context) {
	user := mocks.GenerateMockUser()
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (c *MockUserController) UpdateProfile(ctx *gin.Context) {
	user := mocks.GenerateMockUser()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"user":    user,
	})
}

func (c *MockUserController) GetSettings(ctx *gin.Context) {
	settings := mocks.GenerateMockSettings()
	ctx.JSON(http.StatusOK, gin.H{
		"settings": settings,
	})
}

func (c *MockUserController) UpdateSettings(ctx *gin.Context) {
	settings := mocks.GenerateMockSettings()
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Settings updated successfully",
		"settings": settings,
	})
}

// MockAccountController handles bank account endpoints with mock data
type MockAccountController struct{}

func NewMockAccountController() *MockAccountController {
	return &MockAccountController{}
}

func (c *MockAccountController) List(ctx *gin.Context) {
	accounts := mocks.GenerateMockAccounts()
	ctx.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
		"total":    len(accounts),
	})
}

func (c *MockAccountController) Create(ctx *gin.Context) {
	account := mocks.GenerateMockAccounts()[0]
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Account added successfully",
		"account": account,
	})
}

func (c *MockAccountController) Get(ctx *gin.Context) {
	account := mocks.GenerateMockAccounts()[0]
	ctx.JSON(http.StatusOK, gin.H{
		"account": account,
	})
}

func (c *MockAccountController) Update(ctx *gin.Context) {
	account := mocks.GenerateMockAccounts()[0]
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Account updated successfully",
		"account": account,
	})
}

func (c *MockAccountController) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Account removed successfully",
	})
}

func (c *MockAccountController) Sync(ctx *gin.Context) {
	transactions := mocks.GenerateMockTransactions()[:10]
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "Sync completed successfully",
		"new_transactions": len(transactions),
		"last_sync":    time.Now(),
	})
}

// MockTransactionController handles transaction endpoints with mock data
type MockTransactionController struct{}

func NewMockTransactionController() *MockTransactionController {
	return &MockTransactionController{}
}

func (c *MockTransactionController) List(ctx *gin.Context) {
	transactions := mocks.GenerateMockTransactions()
	ctx.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
		"total":        len(transactions),
		"page":         1,
		"per_page":     50,
	})
}

func (c *MockTransactionController) Get(ctx *gin.Context) {
	transaction := mocks.GenerateMockTransactions()[0]
	ctx.JSON(http.StatusOK, gin.H{
		"transaction": transaction,
	})
}

func (c *MockTransactionController) Update(ctx *gin.Context) {
	transaction := mocks.GenerateMockTransactions()[0]
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "Transaction updated successfully",
		"transaction": transaction,
	})
}

func (c *MockTransactionController) Categorize(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Transaction categorized successfully",
		"category": "Alimentação",
	})
}

// MockSubscriptionController handles subscription endpoints with mock data
type MockSubscriptionController struct{}

func NewMockSubscriptionController() *MockSubscriptionController {
	return &MockSubscriptionController{}
}

func (c *MockSubscriptionController) List(ctx *gin.Context) {
	subscriptions := mocks.GenerateMockSubscriptions()
	ctx.JSON(http.StatusOK, gin.H{
		"subscriptions": subscriptions,
		"total":         len(subscriptions),
		"total_cost":    106.97,
	})
}

func (c *MockSubscriptionController) Create(ctx *gin.Context) {
	subscription := mocks.GenerateMockSubscriptions()[0]
	ctx.JSON(http.StatusCreated, gin.H{
		"message":      "Subscription added successfully",
		"subscription": subscription,
	})
}

func (c *MockSubscriptionController) Get(ctx *gin.Context) {
	subscription := mocks.GenerateMockSubscriptions()[0]
	ctx.JSON(http.StatusOK, gin.H{
		"subscription": subscription,
	})
}

func (c *MockSubscriptionController) Update(ctx *gin.Context) {
	subscription := mocks.GenerateMockSubscriptions()[0]
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "Subscription updated successfully",
		"subscription": subscription,
	})
}

func (c *MockSubscriptionController) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Subscription cancelled successfully",
	})
}

func (c *MockSubscriptionController) Detect(ctx *gin.Context) {
	subscriptions := mocks.GenerateMockSubscriptions()[:2]
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Auto-detection completed",
		"detected":  len(subscriptions),
		"subscriptions": subscriptions,
	})
}

// MockBudgetController handles budget endpoints with mock data
type MockBudgetController struct{}

func NewMockBudgetController() *MockBudgetController {
	return &MockBudgetController{}
}

func (c *MockBudgetController) List(ctx *gin.Context) {
	budgets := mocks.GenerateMockBudgets()
	ctx.JSON(http.StatusOK, gin.H{
		"budgets": budgets,
		"total":   len(budgets),
	})
}

func (c *MockBudgetController) Create(ctx *gin.Context) {
	budget := mocks.GenerateMockBudgets()[0]
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Budget created successfully",
		"budget":  budget,
	})
}

func (c *MockBudgetController) Get(ctx *gin.Context) {
	budget := mocks.GenerateMockBudgets()[0]
	ctx.JSON(http.StatusOK, gin.H{
		"budget": budget,
	})
}

func (c *MockBudgetController) Update(ctx *gin.Context) {
	budget := mocks.GenerateMockBudgets()[0]
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Budget updated successfully",
		"budget":  budget,
	})
}

func (c *MockBudgetController) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Budget deleted successfully",
	})
}

// MockReportController handles report endpoints with mock data
type MockReportController struct{}

func NewMockReportController() *MockReportController {
	return &MockReportController{}
}

func (c *MockReportController) GetSummary(ctx *gin.Context) {
	summary := mocks.GenerateMockReportSummary()
	ctx.JSON(http.StatusOK, gin.H{
		"summary": summary,
	})
}

func (c *MockReportController) GetSpending(ctx *gin.Context) {
	spending := map[string]interface{}{
		"period": map[string]interface{}{
			"start": time.Now().AddDate(0, -1, 0),
			"end":   time.Now(),
		},
		"categories": []map[string]interface{}{
			{"name": "Alimentação", "amount": 543.21, "transactions": 42},
			{"name": "Serviços", "amount": 487.99, "transactions": 8},
			{"name": "Transportes", "amount": 298.50, "transactions": 15},
			{"name": "Entretenimento", "amount": 178.45, "transactions": 12},
			{"name": "Saúde", "amount": 234.67, "transactions": 5},
		},
		"daily_average": 71.59,
		"total":         2147.83,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"spending": spending,
	})
}

func (c *MockReportController) GetSubscriptionCosts(ctx *gin.Context) {
	costs := map[string]interface{}{
		"period": map[string]interface{}{
			"start": time.Now().AddDate(0, -1, 0),
			"end":   time.Now(),
		},
		"monthly_cost": 106.97,
		"yearly_cost":  1283.64,
		"by_category": []map[string]interface{}{
			{"category": "Entretenimento", "amount": 21.98, "count": 2},
			{"category": "Saúde", "amount": 35.00, "count": 1},
			{"category": "Serviços", "amount": 49.99, "count": 1},
		},
		"trend": []map[string]interface{}{
			{"month": "Janeiro", "amount": 106.97},
			{"month": "Fevereiro", "amount": 106.97},
			{"month": "Março", "amount": 106.97},
		},
	}
	ctx.JSON(http.StatusOK, gin.H{
		"subscription_costs": costs,
	})
}

func (c *MockReportController) ExportData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "Export initiated successfully",
		"export_id":    uuid.New().String(),
		"status":       "processing",
		"download_url": "/api/v1/exports/" + uuid.New().String(),
	})
}
