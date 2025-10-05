package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionStatus string
type BillingCycle string

const (
	SubscriptionStatusActive   SubscriptionStatus = "active"
	SubscriptionStatusPaused   SubscriptionStatus = "paused"
	SubscriptionStatusCanceled SubscriptionStatus = "canceled"
	SubscriptionStatusExpired  SubscriptionStatus = "expired"
	
	BillingCycleMonthly   BillingCycle = "monthly"
	BillingCycleQuarterly BillingCycle = "quarterly"
	BillingCycleYearly    BillingCycle = "yearly"
	BillingCycleWeekly    BillingCycle = "weekly"
	BillingCycleCustom    BillingCycle = "custom"
)

type Subscription struct {
	ID              uuid.UUID          `gorm:"type:uuid;primary_key" json:"id"`
	UserID          uuid.UUID          `gorm:"type:uuid;not null;index" json:"user_id"`
	AccountID       uuid.UUID          `gorm:"type:uuid;not null;index" json:"account_id"`
	
	// Subscription details
	Name            string             `gorm:"not null" json:"name"`
	Description     string             `json:"description,omitempty"`
	MerchantName    string             `json:"merchant_name"`
	Category        string             `json:"category"`
	
	// Billing information
	Amount          float64            `gorm:"not null" json:"amount"`
	Currency        string             `gorm:"default:'EUR'" json:"currency"`
	BillingCycle    BillingCycle       `gorm:"type:varchar(20);not null" json:"billing_cycle"`
	CustomDays      int                `json:"custom_days,omitempty"` // For custom billing cycles
	
	// Dates
	StartDate       time.Time          `gorm:"not null" json:"start_date"`
	NextBillingDate time.Time          `gorm:"not null;index" json:"next_billing_date"`
	LastBillingDate *time.Time         `json:"last_billing_date,omitempty"`
	EndDate         *time.Time         `json:"end_date,omitempty"`
	
	// Status and detection
	Status          SubscriptionStatus `gorm:"type:varchar(20);default:'active'" json:"status"`
	AutoDetected    bool               `gorm:"default:false" json:"auto_detected"`
	Confidence      float32            `json:"confidence,omitempty"` // Detection confidence 0-1
	
	// Notification settings
	NotifyDaysBefore int               `gorm:"default:3" json:"notify_days_before"`
	NotifyEnabled    bool              `gorm:"default:true" json:"notify_enabled"`
	
	// Cost tracking
	TotalPaid       float64            `json:"total_paid"`
	PaymentCount    int                `json:"payment_count"`
	
	// Additional data
	Website         string             `json:"website,omitempty"`
	LogoURL         string             `json:"logo_url,omitempty"`
	Color           string             `json:"color,omitempty"` // Brand color for UI
	Notes           string             `json:"notes,omitempty"`
	
	// Cancellation
	CancelURL       string             `json:"cancel_url,omitempty"`
	CancelledAt     *time.Time         `json:"cancelled_at,omitempty"`
	CancellationReason string          `json:"cancellation_reason,omitempty"`
	
	// Timestamps
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
	
	// Relations
	User            User               `gorm:"foreignKey:UserID" json:"-"`
	Account         Account            `gorm:"foreignKey:AccountID" json:"-"`
	Transactions    []Transaction      `gorm:"foreignKey:SubscriptionID" json:"transactions,omitempty"`
}

// BeforeCreate hook to set UUID
func (s *Subscription) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}

// TableName specifies the table name for the Subscription model
func (Subscription) TableName() string {
	return "subscriptions"
}

// IsActive returns true if the subscription is active
func (s *Subscription) IsActive() bool {
	return s.Status == SubscriptionStatusActive
}

// GetMonthlyAmount calculates the monthly cost of the subscription
func (s *Subscription) GetMonthlyAmount() float64 {
	switch s.BillingCycle {
	case BillingCycleMonthly:
		return s.Amount
	case BillingCycleQuarterly:
		return s.Amount / 3
	case BillingCycleYearly:
		return s.Amount / 12
	case BillingCycleWeekly:
		return s.Amount * 4.33 // Average weeks per month
	case BillingCycleCustom:
		if s.CustomDays > 0 {
			return s.Amount * (30.0 / float64(s.CustomDays))
		}
		return s.Amount
	default:
		return s.Amount
	}
}

// GetYearlyAmount calculates the yearly cost of the subscription
func (s *Subscription) GetYearlyAmount() float64 {
	return s.GetMonthlyAmount() * 12
}