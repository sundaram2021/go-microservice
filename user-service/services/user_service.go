// services/user_service.go
package services

import (
	"errors"
	"user-service/models"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

// NewUserService creates a new UserService instance
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

// RegisterUser registers a new user in the database
func (s *UserService) RegisterUser(user *models.User) error {
	// Check if the user already exists
	var existingUser models.User
	if err := s.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return errors.New("user already exists")
	}

	// Create a new user
	if err := s.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

// AuthenticateUser verifies the user's credentials and returns a boolean status
func (s *UserService) AuthenticateUser(username, password string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}
	return &user, nil
}
