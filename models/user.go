package models

import (
	"database/sql"
	"encoding/base64"
	"log"
	"os"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Profile  string `json:"profile"` // Base64-encoded image
}

// CreateUser creates a new user in the database
func CreateUser(db *sql.DB, user User) error {
	query := "INSERT INTO users (username, password, profile) VALUES (?, ?, ?)"
	_, err := db.Exec(query, user.Username, user.Password, user.Profile)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}

// GetUserByUsername retrieves a user by username
func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	query := "SELECT id, username, password, profile FROM users WHERE username = ?"
	row := db.QueryRow(query, username)

	var user User
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Profile); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error fetching user: %v", err)
		return nil, err
	}
	return &user, nil
}

// UpdateUserProfile updates the user's profile picture
func UpdateUserProfile(db *sql.DB, username, profileBase64 string) error {
	query := "UPDATE users SET profile = ? WHERE username = ?"
	_, err := db.Exec(query, profileBase64, username)
	if err != nil {
		log.Printf("Error updating user profile: %v", err)
		return err
	}
	return nil
}

// ConvertImageToBase64 converts an image file to a Base64 string
func ConvertImageToBase64(filePath string) (string, error) {
	imageData, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading image file: %v", err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(imageData), nil
}
