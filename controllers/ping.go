package controllers

import (
	"net/http"
	"root/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Ping(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	user := username.(string)

	// Rate Limit (Chỉ cho phép 2 request / 60s)
	keyRateLimit := "ratelimit:" + user
	count, err := config.Redis.Get(config.Ctx, keyRateLimit).Int()
	if err != nil && err.Error() != "redis: nil" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
		return
	}
	if count >= 2 {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded (max 2 requests per 60s)"})
		return
	}

	// Tăng số lần gọi API & đặt timeout 60s
	config.Redis.Incr(config.Ctx, keyRateLimit)
	config.Redis.Expire(config.Ctx, keyRateLimit, 60*time.Second)

	// Singleton Lock (Chỉ cho phép 1 user gọi API cùng lúc trong 5s)
	lockKey := "ping-lock"
	lockValue := uuid.NewString() // Dùng UUID để đảm bảo khóa không bị xóa bởi thread khác

	locked := config.Redis.SetNX(config.Ctx, lockKey, lockValue, 5*time.Second).Val()
	if !locked {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Only one user can call /ping at a time"})
		return
	}

	// Đảm bảo khóa chỉ bị xóa nếu đúng user đó đặt
	defer func() {
		val, _ := config.Redis.Get(config.Ctx, lockKey).Result()
		if val == lockValue {
			config.Redis.Del(config.Ctx, lockKey)
		}
	}()

	// Tăng số lần user đã gọi API
	keyUserCount := "ping-count:" + user
	config.Redis.Incr(config.Ctx, keyUserCount)

	// Lưu user vào HyperLogLog (lưu số lượng user gọi API)
	config.Redis.PFAdd(config.Ctx, "ping-users", user)

	// Delay 5s để giả lập xử lý lâu
	time.Sleep(5 * time.Second)

	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
