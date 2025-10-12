package repository

import (
	"time"

	"github.com/rmar-dev/suma-backend/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// Create inserts a new transaction
func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

// CreateBatch inserts multiple transactions at once
func (r *TransactionRepository) CreateBatch(transactions []*models.Transaction) error {
	return r.db.CreateInBatches(transactions, 100).Error
}

// GetByID retrieves a transaction by ID
func (r *TransactionRepository) GetByID(id string) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Where("id = ?", id).First(&transaction).Error
	return &transaction, err
}

// GetByAccountID retrieves all transactions for an account
func (r *TransactionRepository) GetByAccountID(accountID string, limit int) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	query := r.db.Where("account_id = ?", accountID).Order("transaction_date DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&transactions).Error
	return transactions, err
}

// GetByDateRange retrieves transactions within a date range
func (r *TransactionRepository) GetByDateRange(accountID string, startDate, endDate time.Time) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.Where("account_id = ? AND transaction_date >= ? AND transaction_date <= ?",
		accountID, startDate, endDate).
		Order("transaction_date DESC").
		Find(&transactions).Error
	return transactions, err
}

// GetByUserID retrieves all transactions for a user across all accounts
func (r *TransactionRepository) GetByUserID(userID string, limit int) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	query := r.db.Where("user_id = ?", userID).Order("transaction_date DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&transactions).Error
	return transactions, err
}

// GetByCategory retrieves transactions by category
func (r *TransactionRepository) GetByCategory(userID, category string) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.Where("user_id = ? AND (category = ? OR user_category = ?)",
		userID, category, category).
		Order("transaction_date DESC").
		Find(&transactions).Error
	return transactions, err
}

// Update saves changes to a transaction
func (r *TransactionRepository) Update(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

// UpdateCategory updates the user-defined category for a transaction
func (r *TransactionRepository) UpdateCategory(id, category string) error {
	return r.db.Model(&models.Transaction{}).
		Where("id = ?", id).
		Update("user_category", category).Error
}

// GetRecurringTransactions finds potential recurring transactions
func (r *TransactionRepository) GetRecurringTransactions(accountID string) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	// Get transactions from last 90 days that are likely recurring
	startDate := time.Now().AddDate(0, 0, -90)

	err := r.db.Where(`account_id = ? AND 
		transaction_date >= ? AND 
		amount < 0 AND
		merchant_name IS NOT NULL AND 
		merchant_name != ''`,
		accountID, startDate).
		Order("merchant_name, transaction_date DESC").
		Find(&transactions).Error

	return transactions, err
}

// GetMonthlySpending calculates total spending for a month
func (r *TransactionRepository) GetMonthlySpending(userID string, year, month int) (float64, error) {
	var total float64

	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, -1)

	err := r.db.Model(&models.Transaction{}).
		Select("SUM(ABS(amount))").
		Where("user_id = ? AND transaction_date >= ? AND transaction_date <= ? AND amount < 0",
			userID, startDate, endDate).
		Scan(&total).Error

	return total, err
}

// GetCategorySpending calculates spending by category
func (r *TransactionRepository) GetCategorySpending(userID string, startDate, endDate time.Time) (map[string]float64, error) {
	type CategorySum struct {
		Category string
		Total    float64
	}

	var results []CategorySum

	err := r.db.Model(&models.Transaction{}).
		Select("COALESCE(user_category, category, 'Uncategorized') as category, SUM(ABS(amount)) as total").
		Where("user_id = ? AND transaction_date >= ? AND transaction_date <= ? AND amount < 0",
			userID, startDate, endDate).
		Group("COALESCE(user_category, category, 'Uncategorized')").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	categoryMap := make(map[string]float64)
	for _, result := range results {
		categoryMap[result.Category] = result.Total
	}

	return categoryMap, nil
}

// CheckDuplicateTransaction checks if a transaction already exists (by external_id)
func (r *TransactionRepository) CheckDuplicateTransaction(externalID string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Transaction{}).
		Where("external_id = ?", externalID).
		Count(&count).Error
	return count > 0, err
}
