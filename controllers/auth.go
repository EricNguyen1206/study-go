package controllers

import (
	"net/http"
	"root/config"
	"root/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var user models.User
	var input models.UserLoginDto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Where("username = ?", input.Username).First(&user)
	if user.ID == 0 {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}

	sessionID := uuid.New().String()
	config.Redis.Set(config.Ctx, "session:"+sessionID, user.Username, time.Hour)

	c.JSON(200, gin.H{"session_id": sessionID})
}

func SignUp(c *gin.Context) {
	var input models.UserSignUpDto

	// Validate the input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check User exist or not
	var existingUser models.User
	if err := config.DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Hash the password before save to DB
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create new User
	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
		Profile:  input.Profile,
	}

	// Save to DB
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
