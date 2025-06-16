package auth

import (
	"time"

	"github.com/Lewvy/markable/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("SecretKey")

func GenerateJWT(s *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["authorized"] = true
	claims["user"] = s.UserName
	claims["role"] = s.Role

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
