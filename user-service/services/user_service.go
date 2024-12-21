package services

import (
	"errors"
	"user-service/models"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) RegisterUser(user *models.User) error {
	var existingUser models.User
	if err := s.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return errors.New("user already exists")
	}

	if err := s.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (s *UserService) AuthenticateUser(username, password string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}
	return &user, nil
}
