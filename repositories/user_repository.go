package repositories

import (
	"database/sql"
	"log"

	"root/models"
)

// UserRepository handles database operations for users
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// Create inserts a new user into the database
func (repo *UserRepository) Create(user models.User) error {
	query := "INSERT INTO users (username, password, profile) VALUES (?, ?, ?)"
	_, err := repo.DB.Exec(query, user.Username, user.Password, user.Profile)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return err
	}
	return nil
}

// FindByUsername retrieves a user by their username
func (repo *UserRepository) FindByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, password, profile FROM users WHERE username = ?"
	row := repo.DB.QueryRow(query, username)

	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Profile); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error finding user by username: %v", err)
		return nil, err
	}
	return &user, nil
}

// UpdateProfile updates the profile image for a user
func (repo *UserRepository) UpdateProfile(username, profileBase64 string) error {
	query := "UPDATE users SET profile = ? WHERE username = ?"
	_, err := repo.DB.Exec(query, profileBase64, username)
	if err != nil {
		log.Printf("Error updating user profile: %v", err)
		return err
	}
	return nil
}
