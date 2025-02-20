package middleware

import (
	"root/config"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID := c.Request.Header.Get("Session-ID")
		if sessionID == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		username, err := config.Redis.Get(config.Ctx, "session:"+sessionID).Result()
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid session"})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Next()
	}
}
