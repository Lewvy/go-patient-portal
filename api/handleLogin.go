package api

import (
	"context"
	"time"

	"github.com/Lewvy/markable/internal/auth"
	"github.com/Lewvy/markable/internal/models"
	"github.com/gin-gonic/gin"
)

func HandleLogin(s *models.State) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		staffMember, err := s.Db.GetStaffMember(ctx, user.UserName)
		if err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		if !auth.MatchPassword(user.Password, staffMember.PwHash) {
			c.JSON(401, gin.H{"message": "Incorrect Password"})
			return
		}
		user.Role = staffMember.Role
		token, err := auth.GenerateJWT(s, &user)
		if err != nil {
			c.JSON(401, gin.H{"Error generating jwt token": err})
		}
		c.JSON(200, gin.H{
			"message": "Logged in successfully",
			"token":   token,
		})
	}
}
