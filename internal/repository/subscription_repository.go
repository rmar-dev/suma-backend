package repository

import (
	"time"

	"github.com/suma/finance-app-api/internal/models"
	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

// Create inserts a new subscription
func (r *SubscriptionRepository) Create(subscription *models.Subscription) error {
	return r.db.Create(subscription).Error
}

// GetByID retrieves a subscription by ID
func (r *SubscriptionRepository) GetByID(id string) (*models.Subscription, error) {
	var subscription models.Subscription
	err := r.db.Where("id = ?", id).First(&subscription).Error
	return &subscription, err
}

// GetByUserID retrieves all subscriptions for a user
func (r *SubscriptionRepository) GetByUserID(userID string) ([]*models.Subscription, error) {
	var subscriptions []*models.Subscription
	err := r.db.Where("user_id = ?", userID).Find(&subscriptions).Error
	return subscriptions, err
}

// GetActiveByUserID retrieves only active subscriptions for a user
func (r *SubscriptionRepository) GetActiveByUserID(userID string) ([]*models.Subscription, error) {
	var subscriptions []*models.Subscription
	err := r.db.Where("user_id = ? AND status = ?", userID, "active").Find(&subscriptions).Error
	return subscriptions, err
}

// Update saves changes to a subscription
func (r *SubscriptionRepository) Update(subscription *models.Subscription) error {
	return r.db.Save(subscription).Error
}

// Delete removes a subscription
func (r *SubscriptionRepository) Delete(id string) error {
	return r.db.Delete(&models.Subscription{}, "id = ?", id).Error
}

// GetUpcomingPayments retrieves subscriptions with upcoming billing dates
func (r *SubscriptionRepository) GetUpcomingPayments(userID string, endDate time.Time) ([]*models.Subscription, error) {
	var subscriptions []*models.Subscription
	err := r.db.Where("user_id = ? AND status = ? AND next_billing_date <= ?", 
		userID, "active", endDate).
		Order("next_billing_date ASC").
		Find(&subscriptions).Error
	return subscriptions, err
}

// GetByAccountID retrieves all subscriptions for a specific account
func (r *SubscriptionRepository) GetByAccountID(accountID string) ([]*models.Subscription, error) {
	var subscriptions []*models.Subscription
	err := r.db.Where("account_id = ?", accountID).Find(&subscriptions).Error
	return subscriptions, err
}

// CheckDuplicateSubscription checks if a similar subscription already exists
func (r *SubscriptionRepository) CheckDuplicateSubscription(userID, merchantName string, amount float64) (bool, error) {
	var count int64
	err := r.db.Model(&models.Subscription{}).
		Where("user_id = ? AND merchant_name = ? AND amount = ? AND status = ?", 
			userID, merchantName, amount, "active").
		Count(&count).Error
	return count > 0, err
}