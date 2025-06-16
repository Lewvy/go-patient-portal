package auth

import (
	"log"
	"time"

	"github.com/Lewvy/markable/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(state *models.State, user *models.User) (string, error) {
	jwtKey := []byte(state.Cfg.JwtSecret)
	token := jwt.New(jwt.SigningMethodHS256)
	log.Println("USER ROLE: ", user.Role)
	log.Println("USER NAME: ", user.UserName)
	log.Println("USER ROLE: ", user.Role)
	log.Println("USER ROLE: ", user.Role)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["authorized"] = true
	claims["user"] = user.UserName
	claims["role"] = user.Role

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
