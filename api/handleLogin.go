package api

import (
	"context"
	"fmt"
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
		hashedPw, err := s.Db.GetUserPasswdHash(ctx, user.UserName)
		if err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		fmt.Println(user.UserName, user.Password, hashedPw)
		if !auth.MatchPassword(user.Password, hashedPw) {
			c.JSON(401, gin.H{"message": "Incorrect Password"})
			return
		}
		token, err := auth.GenerateJWT(&user)
		if err != nil {
			c.JSON(401, gin.H{"Error generating jwt token": err})
		}
		c.JSON(200, gin.H{
			"message": "Logged in successfully",
			"token":   token,
		})
	}
}
