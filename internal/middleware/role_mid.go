package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"slices"
)

func RequiredRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("user")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
			c.Abort()
			return
		}
		role, ok := claims.(jwt.MapClaims)["role"].(string)
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
			c.Abort()
			return
		}
		if slices.Contains(roles, role) {
			c.Next()
			return
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		c.Abort()
	}
}
