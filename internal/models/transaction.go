package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionType string
type TransactionStatus string

const (
	TransactionTypeDebit  TransactionType = "debit"
	TransactionTypeCredit TransactionType = "credit"

	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusCompleted TransactionStatus = "completed"
	TransactionStatusFailed    TransactionStatus = "failed"
)

type Transaction struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	AccountID  uuid.UUID `gorm:"type:uuid;not null;index" json:"account_id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	ExternalID string    `gorm:"uniqueIndex" json:"-"` // ID from bank API

	// Transaction details
	Amount   float64           `gorm:"not null" json:"amount"`
	Currency string            `gorm:"default:'EUR'" json:"currency"`
	Type     TransactionType   `gorm:"type:varchar(20);not null" json:"type"`
	Status   TransactionStatus `gorm:"type:varchar(20);default:'completed'" json:"status"`

	// Description and categorization
	Description  string `gorm:"not null" json:"description"`
	MerchantName string `json:"merchant_name,omitempty"`
	Category     string `json:"category,omitempty"`
	SubCategory  string `json:"sub_category,omitempty"`
	UserCategory string `json:"user_category,omitempty"` // User-defined category
	Notes        string `json:"notes,omitempty"`

	// Dates
	TransactionDate time.Time `gorm:"not null;index" json:"transaction_date"`
	PostedDate      time.Time `json:"posted_date"`

	// Subscription detection
	IsRecurring    bool       `gorm:"default:false" json:"is_recurring"`
	SubscriptionID *uuid.UUID `gorm:"type:uuid" json:"subscription_id,omitempty"`

	// Location data (if available)
	Location string `json:"location,omitempty"`
	Country  string `json:"country,omitempty"`

	// Additional metadata
	Metadata map[string]interface{} `gorm:"type:jsonb" json:"metadata,omitempty"`

	// Timestamps
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Account      Account       `gorm:"foreignKey:AccountID" json:"-"`
	User         User          `gorm:"foreignKey:UserID" json:"-"`
	Subscription *Subscription `gorm:"foreignKey:SubscriptionID" json:"subscription,omitempty"`
}

// BeforeCreate hook to set UUID
func (t *Transaction) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}

// TableName specifies the table name for the Transaction model
func (Transaction) TableName() string {
	return "transactions"
}

// IsDebit returns true if the transaction is a debit
func (t *Transaction) IsDebit() bool {
	return t.Type == TransactionTypeDebit
}

// IsCredit returns true if the transaction is a credit
func (t *Transaction) IsCredit() bool {
	return t.Type == TransactionTypeCredit
}
