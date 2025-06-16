package api

import (
	"context"
	"time"

	"github.com/Lewvy/markable/internal/auth"
	"github.com/Lewvy/markable/internal/database"
	"github.com/Lewvy/markable/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleRegistration(s *models.State) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		pwHash, err := auth.HashPassword(user.Password)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to hash password"})
			return
		}

		reqUser := &database.CreateUserParams{
			ID:        uuid.New(),
			Name:      user.UserName,
			PwHash:    pwHash,
			Role:      user.Role,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		u, err := s.Db.CreateUser(ctx, *reqUser)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create user"})
			return
		}
		c.JSON(201, u)
	}
}
