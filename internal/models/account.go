package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountType string
type AccountProvider string

const (
	AccountTypeChecking   AccountType = "checking"
	AccountTypeSavings    AccountType = "savings"
	AccountTypeCredit     AccountType = "credit"
	AccountTypeInvestment AccountType = "investment"

	// Portuguese banks
	ProviderCGD        AccountProvider = "cgd"
	ProviderMillennium AccountProvider = "millennium"
	ProviderNovoBanco  AccountProvider = "novo_banco"
	ProviderSantander  AccountProvider = "santander"
	ProviderBPI        AccountProvider = "bpi"
	ProviderTink       AccountProvider = "tink"
	ProviderManual     AccountProvider = "manual"
)

type Account struct {
	ID            uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	UserID        uuid.UUID       `gorm:"type:uuid;not null;index" json:"user_id"`
	Name          string          `gorm:"not null" json:"name"`
	AccountType   AccountType     `gorm:"type:varchar(20);not null" json:"account_type"`
	Provider      AccountProvider `gorm:"type:varchar(50);not null" json:"provider"`
	AccountNumber string          `gorm:"encrypted" json:"-"` // Encrypted in database
	MaskedNumber  string          `json:"masked_number"`      // Last 4 digits only
	Balance       float64         `json:"balance"`
	Currency      string          `gorm:"default:'EUR'" json:"currency"`

	// Bank connection details
	ExternalID      string     `json:"-"`                  // ID from bank/Tink API
	ConnectionToken string     `gorm:"encrypted" json:"-"` // Encrypted token
	LastSyncAt      *time.Time `json:"last_sync_at,omitempty"`
	SyncError       string     `json:"sync_error,omitempty"`

	// Status
	Active    bool       `gorm:"default:true" json:"active"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`

	// Timestamps
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	User         User          `gorm:"foreignKey:UserID" json:"-"`
	Transactions []Transaction `gorm:"foreignKey:AccountID" json:"transactions,omitempty"`
}

// BeforeCreate hook to set UUID
func (a *Account) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}

// TableName specifies the table name for the Account model
func (Account) TableName() string {
	return "accounts"
}
