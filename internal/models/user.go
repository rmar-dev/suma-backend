package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID                uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	Email             string     `gorm:"uniqueIndex;not null" json:"email"`
	Password          string     `gorm:"not null" json:"-"` // Never send password in JSON
	FirstName         string     `gorm:"not null" json:"first_name"`
	LastName          string     `gorm:"not null" json:"last_name"`
	PhoneNumber       string     `json:"phone_number,omitempty"`
	EmailVerified     bool       `gorm:"default:false" json:"email_verified"`
	EmailVerifiedAt   *time.Time `json:"email_verified_at,omitempty"`
	TwoFactorEnabled  bool       `gorm:"default:false" json:"two_factor_enabled"`
	ProfilePictureURL string     `json:"profile_picture_url,omitempty"`
	Language          string     `gorm:"default:'pt'" json:"language"`
	Currency          string     `gorm:"default:'EUR'" json:"currency"`
	Timezone          string     `gorm:"default:'Europe/Lisbon'" json:"timezone"`

	// Privacy & Compliance
	GDPRConsentAt     *time.Time `json:"gdpr_consent_at,omitempty"`
	TermsAcceptedAt   *time.Time `json:"terms_accepted_at,omitempty"`
	PrivacyAcceptedAt *time.Time `json:"privacy_accepted_at,omitempty"`
	MarketingConsent  bool       `gorm:"default:false" json:"marketing_consent"`

	// Account Status
	Active    bool       `gorm:"default:true" json:"active"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	LastLogin *time.Time `json:"last_login,omitempty"`

	// Timestamps
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Accounts      []Account      `gorm:"foreignKey:UserID" json:"accounts,omitempty"`
	Subscriptions []Subscription `gorm:"foreignKey:UserID" json:"subscriptions,omitempty"`
}

// BeforeCreate hook to set UUID and hash password
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	if u.Password != "" {
		hashedPassword, err := u.HashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hashedPassword
	}

	return nil
}

// HashPassword hashes the user password
func (u *User) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword checks if the provided password is correct
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}

// UserResponse is the response structure for user data (without sensitive info)
type UserResponse struct {
	ID                uuid.UUID  `json:"id"`
	Email             string     `json:"email"`
	FirstName         string     `json:"first_name"`
	LastName          string     `json:"last_name"`
	PhoneNumber       string     `json:"phone_number,omitempty"`
	EmailVerified     bool       `json:"email_verified"`
	TwoFactorEnabled  bool       `json:"two_factor_enabled"`
	ProfilePictureURL string     `json:"profile_picture_url,omitempty"`
	Language          string     `json:"language"`
	Currency          string     `json:"currency"`
	Timezone          string     `json:"timezone"`
	CreatedAt         time.Time  `json:"created_at"`
	LastLogin         *time.Time `json:"last_login,omitempty"`
}

// ToResponse converts User to UserResponse
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:                u.ID,
		Email:             u.Email,
		FirstName:         u.FirstName,
		LastName:          u.LastName,
		PhoneNumber:       u.PhoneNumber,
		EmailVerified:     u.EmailVerified,
		TwoFactorEnabled:  u.TwoFactorEnabled,
		ProfilePictureURL: u.ProfilePictureURL,
		Language:          u.Language,
		Currency:          u.Currency,
		Timezone:          u.Timezone,
		CreatedAt:         u.CreatedAt,
		LastLogin:         u.LastLogin,
	}
}
