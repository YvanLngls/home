package utils

import (
	"back/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("your-secret-key")

func GenerateJWT(username string) (string, time.Time, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	return tokenString, expirationTime, err
}
