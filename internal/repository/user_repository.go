package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/rmar-dev/suma-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create inserts a new user into the database
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// GetByID retrieves a user by their ID
func (r *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user by their email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update saves changes to an existing user
func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// Delete permanently removes a user from the database
func (r *UserRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.User{}, id).Error
}

// SoftDelete marks a user as deleted without removing from database
func (r *UserRepository) SoftDelete(id uuid.UUID) error {
	now := time.Now()
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("deleted_at", now).Error
}

// List retrieves all active users with pagination
func (r *UserRepository) List(offset, limit int) ([]*models.User, error) {
	var users []*models.User
	err := r.db.Where("deleted_at IS NULL").Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

// Count returns the total number of active users
func (r *UserRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("deleted_at IS NULL").Count(&count).Error
	return count, err
}

// UpdateLastLogin updates the last login timestamp for a user
func (r *UserRepository) UpdateLastLogin(id uuid.UUID) error {
	now := time.Now()
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("last_login", now).Error
}

// VerifyEmail marks a user's email as verified
func (r *UserRepository) VerifyEmail(id uuid.UUID) error {
	now := time.Now()
	return r.db.Model(&models.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"email_verified":    true,
		"email_verified_at": now,
	}).Error
}

// UpdateGDPRConsent records GDPR consent
func (r *UserRepository) UpdateGDPRConsent(id uuid.UUID) error {
	now := time.Now()
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("gdpr_consent_at", now).Error
}
