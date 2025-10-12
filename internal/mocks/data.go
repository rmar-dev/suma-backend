package mocks

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// User mock data
func GenerateMockUser() map[string]interface{} {
	return map[string]interface{}{
		"id":                 uuid.New().String(),
		"email":              "john.doe@example.com",
		"first_name":         "John",
		"last_name":          "Doe",
		"phone_number":       "+351912345678",
		"email_verified":     true,
		"two_factor_enabled": false,
		"language":           "pt",
		"currency":           "EUR",
		"timezone":           "Europe/Lisbon",
		"created_at":         time.Now().Add(-30 * 24 * time.Hour),
		"last_login":         time.Now().Add(-2 * time.Hour),
		"is_active":          true,
	}
}

// Account mock data
func GenerateMockAccounts() []map[string]interface{} {
	accounts := []map[string]interface{}{
		{
			"id":             uuid.New().String(),
			"user_id":        uuid.New().String(),
			"bank_name":      "Millennium BCP",
			"account_name":   "Conta à Ordem",
			"account_type":   "checking",
			"account_number": "**** **** **** 1234",
			"iban":           "PT50 **** **** **** **** **** 1",
			"balance":        2543.67,
			"currency":       "EUR",
			"is_primary":     true,
			"last_sync":      time.Now().Add(-1 * time.Hour),
			"created_at":     time.Now().Add(-180 * 24 * time.Hour),
			"status":         "active",
		},
		{
			"id":             uuid.New().String(),
			"user_id":        uuid.New().String(),
			"bank_name":      "Santander Totta",
			"account_name":   "Conta Poupança",
			"account_type":   "savings",
			"account_number": "**** **** **** 5678",
			"iban":           "PT50 **** **** **** **** **** 2",
			"balance":        15234.50,
			"currency":       "EUR",
			"is_primary":     false,
			"last_sync":      time.Now().Add(-2 * time.Hour),
			"created_at":     time.Now().Add(-90 * 24 * time.Hour),
			"status":         "active",
		},
	}
	return accounts
}

// Transaction mock data
func GenerateMockTransactions() []map[string]interface{} {
	categories := []string{"Alimentação", "Transportes", "Entretenimento", "Serviços", "Saúde", "Educação", "Compras"}
	merchants := []string{"Continente", "Pingo Doce", "Galp", "Netflix", "Spotify", "Farmácia Santos", "Fnac", "Worten"}

	var transactions []map[string]interface{}
	now := time.Now()

	for i := 0; i < 50; i++ {
		isExpense := rand.Float32() > 0.2 // 80% expenses, 20% income
		amount := rand.Float64() * 200
		if !isExpense {
			amount = rand.Float64()*3000 + 500 // Income between 500-3500
		}

		transaction := map[string]interface{}{
			"id":           uuid.New().String(),
			"account_id":   uuid.New().String(),
			"amount":       fmt.Sprintf("%.2f", amount),
			"type":         map[bool]string{true: "expense", false: "income"}[isExpense],
			"category":     categories[rand.Intn(len(categories))],
			"merchant":     merchants[rand.Intn(len(merchants))],
			"description":  fmt.Sprintf("Transaction %d", i+1),
			"date":         now.AddDate(0, 0, -i),
			"created_at":   now.AddDate(0, 0, -i),
			"is_recurring": rand.Float32() > 0.8,
			"notes":        "",
			"tags":         []string{},
		}
		transactions = append(transactions, transaction)
	}

	return transactions
}

// Subscription mock data
func GenerateMockSubscriptions() []map[string]interface{} {
	subscriptions := []map[string]interface{}{
		{
			"id":            uuid.New().String(),
			"user_id":       uuid.New().String(),
			"name":          "Netflix",
			"description":   "Streaming de vídeo",
			"amount":        11.99,
			"currency":      "EUR",
			"billing_cycle": "monthly",
			"next_billing":  time.Now().AddDate(0, 0, 15),
			"category":      "Entretenimento",
			"is_active":     true,
			"auto_detected": true,
			"logo_url":      "/images/subscriptions/netflix.png",
			"created_at":    time.Now().AddDate(0, -6, 0),
			"total_spent":   71.94,
		},
		{
			"id":            uuid.New().String(),
			"user_id":       uuid.New().String(),
			"name":          "Spotify",
			"description":   "Streaming de música",
			"amount":        9.99,
			"currency":      "EUR",
			"billing_cycle": "monthly",
			"next_billing":  time.Now().AddDate(0, 0, 10),
			"category":      "Entretenimento",
			"is_active":     true,
			"auto_detected": true,
			"logo_url":      "/images/subscriptions/spotify.png",
			"created_at":    time.Now().AddDate(-1, 0, 0),
			"total_spent":   119.88,
		},
		{
			"id":            uuid.New().String(),
			"user_id":       uuid.New().String(),
			"name":          "Ginásio Fitness Hut",
			"description":   "Mensalidade ginásio",
			"amount":        35.00,
			"currency":      "EUR",
			"billing_cycle": "monthly",
			"next_billing":  time.Now().AddDate(0, 0, 5),
			"category":      "Saúde",
			"is_active":     true,
			"auto_detected": false,
			"logo_url":      "/images/subscriptions/gym.png",
			"created_at":    time.Now().AddDate(0, -3, 0),
			"total_spent":   105.00,
		},
		{
			"id":            uuid.New().String(),
			"user_id":       uuid.New().String(),
			"name":          "MEO",
			"description":   "Internet + TV + Telefone",
			"amount":        49.99,
			"currency":      "EUR",
			"billing_cycle": "monthly",
			"next_billing":  time.Now().AddDate(0, 0, 20),
			"category":      "Serviços",
			"is_active":     true,
			"auto_detected": true,
			"logo_url":      "/images/subscriptions/meo.png",
			"created_at":    time.Now().AddDate(-2, 0, 0),
			"total_spent":   1199.76,
		},
	}
	return subscriptions
}

// Budget mock data
func GenerateMockBudgets() []map[string]interface{} {
	budgets := []map[string]interface{}{
		{
			"id":         uuid.New().String(),
			"user_id":    uuid.New().String(),
			"name":       "Orçamento Mensal",
			"period":     "monthly",
			"start_date": time.Now().AddDate(0, 0, -time.Now().Day()+1),
			"end_date":   time.Now().AddDate(0, 1, -time.Now().Day()),
			"categories": []map[string]interface{}{
				{
					"name":      "Alimentação",
					"allocated": 400.00,
					"spent":     287.43,
					"color":     "#10B981",
				},
				{
					"name":      "Transportes",
					"allocated": 150.00,
					"spent":     98.50,
					"color":     "#3B82F6",
				},
				{
					"name":      "Entretenimento",
					"allocated": 100.00,
					"spent":     67.98,
					"color":     "#8B5CF6",
				},
				{
					"name":      "Serviços",
					"allocated": 200.00,
					"spent":     185.99,
					"color":     "#F59E0B",
				},
			},
			"total_allocated": 850.00,
			"total_spent":     639.90,
			"created_at":      time.Now().AddDate(0, -3, 0),
			"is_active":       true,
		},
		{
			"id":         uuid.New().String(),
			"user_id":    uuid.New().String(),
			"name":       "Poupança para Férias",
			"period":     "custom",
			"start_date": time.Now().AddDate(0, -6, 0),
			"end_date":   time.Now().AddDate(0, 6, 0),
			"categories": []map[string]interface{}{
				{
					"name":      "Poupança",
					"allocated": 3000.00,
					"spent":     1850.00,
					"color":     "#059669",
				},
			},
			"total_allocated": 3000.00,
			"total_spent":     1850.00,
			"created_at":      time.Now().AddDate(0, -6, 0),
			"is_active":       true,
		},
	}
	return budgets
}

// Report summary mock data
func GenerateMockReportSummary() map[string]interface{} {
	return map[string]interface{}{
		"period": map[string]interface{}{
			"start": time.Now().AddDate(0, -1, 0),
			"end":   time.Now(),
		},
		"total_income":        3250.00,
		"total_expenses":      2147.83,
		"net_savings":         1102.17,
		"savings_rate":        33.91,
		"average_daily_spend": 71.59,
		"top_categories": []map[string]interface{}{
			{"name": "Alimentação", "amount": 543.21, "percentage": 25.3},
			{"name": "Serviços", "amount": 487.99, "percentage": 22.7},
			{"name": "Transportes", "amount": 298.50, "percentage": 13.9},
		},
		"monthly_trend": []map[string]interface{}{
			{"month": "Janeiro", "income": 3250.00, "expenses": 2234.56},
			{"month": "Fevereiro", "income": 3250.00, "expenses": 2147.83},
			{"month": "Março", "income": 3250.00, "expenses": 1987.43},
		},
		"accounts_summary": []map[string]interface{}{
			{"name": "Conta à Ordem", "balance": 2543.67, "change": 234.56},
			{"name": "Conta Poupança", "balance": 15234.50, "change": 1102.17},
		},
		"active_subscriptions": 4,
		"subscription_cost":    106.97,
		"upcoming_bills": []map[string]interface{}{
			{"name": "Ginásio Fitness Hut", "amount": 35.00, "date": time.Now().AddDate(0, 0, 5)},
			{"name": "Spotify", "amount": 9.99, "date": time.Now().AddDate(0, 0, 10)},
		},
	}
}

// Settings preferences
func GenerateMockSettings() map[string]interface{} {
	return map[string]interface{}{
		"notifications": map[string]interface{}{
			"email_notifications":    true,
			"push_notifications":     false,
			"transaction_alerts":     true,
			"budget_alerts":          true,
			"subscription_reminders": true,
			"weekly_summary":         true,
			"monthly_report":         true,
		},
		"privacy": map[string]interface{}{
			"show_balance_dashboard": true,
			"require_auth_view":      false,
			"two_factor_enabled":     false,
		},
		"preferences": map[string]interface{}{
			"language":          "pt",
			"currency":          "EUR",
			"timezone":          "Europe/Lisbon",
			"date_format":       "DD/MM/YYYY",
			"number_format":     "1.234,56",
			"start_of_week":     "monday",
			"fiscal_year_start": "january",
		},
		"data_management": map[string]interface{}{
			"auto_categorize":       true,
			"auto_detect_recurring": true,
			"data_retention_months": 24,
			"last_export":           time.Now().AddDate(0, -1, 0),
		},
	}
}
