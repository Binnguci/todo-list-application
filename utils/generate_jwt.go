package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
	"todo-app/config"
)

func GenerateJWT(username string, userID int) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"userID":   userID,
		"exp":      time.Now().Add(config.JWT_EXPIRY).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT_SECRET))
}
