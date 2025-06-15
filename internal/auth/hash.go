package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const cost = 12

func HashPassword(password string) (string, error) {
	pwBytes := []byte(password)
	hashedpw, err := bcrypt.GenerateFromPassword(pwBytes, cost)
	if err != nil {
		return "", fmt.Errorf("Error generating password: %q", err)
	}
	return string(hashedpw), nil
}

func MatchPassword(password, hashedPw string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(password)) != nil {
		return false
	}
	return true
}
