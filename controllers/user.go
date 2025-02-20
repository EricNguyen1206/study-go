package controllers

import (
	"net/http"
	"root/config"
	"root/models"

	"github.com/gin-gonic/gin"
)

// UpdateUserProfile updates the user's profile picture
func UpdateUserProfile(c *gin.Context) {
	var user models.User
	var input models.UserUpdateProfileDto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Where("id = ?", input.ID).First(&user)
	if user.ID == 0 {
		c.JSON(401, gin.H{"error": "Invalid user"})
		return
	}
	// Update the object data
	user.Profile = input.Profile

	// Update data in database
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User update successfully"})
}

// GetUserByUsername retrieves a user by username
// func GetUserByUsername(db *sql.DB, username string) (*User, error) {
// 	query := "SELECT id, username, password, profile FROM users WHERE username = ?"
// 	row := db.QueryRow(query, username)

// 	var user User
// 	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Profile); err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		log.Printf("Error fetching user: %v", err)
// 		return nil, err
// 	}
// 	return &user, nil
// }
