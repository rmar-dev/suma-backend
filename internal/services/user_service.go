package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rmar-dev/suma-backend/internal/models"
	"github.com/rmar-dev/suma-backend/internal/repository"
	"github.com/rmar-dev/suma-backend/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo       *repository.UserRepository
	jwtManager *auth.JWTManager
}

func NewUserService(repo *repository.UserRepository, jwtManager *auth.JWTManager) *UserService {
	return &UserService{
		repo:       repo,
		jwtManager: jwtManager,
	}
}

// Register creates a new user account
func (s *UserService) Register(email, password, firstName, lastName string) (*models.User, error) {
	// Check if user already exists
	existingUser, _ := s.repo.GetByEmail(email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		ID:        uuid.New(),
		Email:     email,
		Password:  string(hashedPassword),
		FirstName: firstName,
		LastName:  lastName,
		Language:  "pt",
		Currency:  "EUR",
		Timezone:  "Europe/Lisbon",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save to database
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	// Clear password before returning
	user.Password = ""
	return user, nil
}

// Login authenticates a user and returns JWT tokens
func (s *UserService) Login(email, password string) (*models.User, string, string, error) {
	// Get user by email
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, "", "", errors.New("invalid credentials")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", "", errors.New("invalid credentials")
	}

	// Generate tokens
	accessToken, refreshToken, err := s.jwtManager.GenerateTokenPair(user.ID, user.Email)
	if err != nil {
		return nil, "", "", err
	}

	// Update last login
	user.LastLogin = &time.Time{}
	*user.LastLogin = time.Now()
if err := s.repo.Update(user); err != nil {
		return nil, "", "", err
	}

	// Clear password before returning
	user.Password = ""
	return user, accessToken, refreshToken, nil
}

// GetProfile returns user profile information
func (s *UserService) GetProfile(userID string) (*models.User, error) {
	id, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Clear sensitive information
	user.Password = ""
	return user, nil
}

// UpdateProfile updates user profile information
func (s *UserService) UpdateProfile(userID string, updates map[string]interface{}) (*models.User, error) {
	id, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Apply updates (add validation as needed)
	if firstName, ok := updates["first_name"].(string); ok {
		user.FirstName = firstName
	}
	if lastName, ok := updates["last_name"].(string); ok {
		user.LastName = lastName
	}
	if language, ok := updates["language"].(string); ok {
		user.Language = language
	}
	if currency, ok := updates["currency"].(string); ok {
		user.Currency = currency
	}
	if timezone, ok := updates["timezone"].(string); ok {
		user.Timezone = timezone
	}

	user.UpdatedAt = time.Now()

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

// DeleteAccount soft deletes a user account
func (s *UserService) DeleteAccount(userID string) error {
	id, err := uuid.Parse(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	return s.repo.SoftDelete(id)
}
