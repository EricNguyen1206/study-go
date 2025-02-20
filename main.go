package main

import (
	"root/config"
	"root/controllers"
	"root/middleware"
	"root/models"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	config.InitRedis()

	// auto migrate models (User, Session)
	err := config.DB.AutoMigrate(&models.User{}, &models.Session{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Auto migration for all tables completed successfully.")

	// Initialize router
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
			auth.POST("/signup", controllers.SignUp)
		}

		protect := v1.Group("/")
		protect.Use(middleware.AuthMiddleware())
		{
			protect.GET("/ping", controllers.Ping)

			user := protect.Group("/user")
			{
				user.PATCH("/:id", controllers.UpdateUserProfile)
			}
		}
	}

	r.Run(":8080")
}
