package services

import (
	"encoding/base64"
	"errors"
	"os"

	"root/models"
	"root/repositories"
)

// UserService handles the business logic for user operations
type UserService struct {
	UserRepo *repositories.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

// RegisterUser registers a new user
func (service *UserService) RegisterUser(username, password string) error {
	// Check if user already exists
	existingUser, err := service.UserRepo.FindByUsername(username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// Create a new user
	user := models.User{
		Username: username,
		Password: password,
		Profile:  "", // Default to empty profile
	}
	return service.UserRepo.Create(user)
}

// LoginUser validates user credentials
func (service *UserService) LoginUser(username, password string) (bool, error) {
	user, err := service.UserRepo.FindByUsername(username)
	if err != nil {
		return false, err
	}
	if user == nil || user.Password != password {
		return false, nil // Invalid credentials
	}
	return true, nil
}

// UpdateUserProfile updates a user's profile image
func (service *UserService) UpdateUserProfile(username, filePath string) error {
	// Convert image to Base64
	imageData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	profileBase64 := base64.StdEncoding.EncodeToString(imageData)

	// Update the profile in the database
	return service.UserRepo.UpdateProfile(username, profileBase64)
}
