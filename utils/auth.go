package utils

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

// jwtKey is the secret key used for signing JWT tokens.
var jwtKey = []byte("secret-key")

// GenerateJWT creates a new JWT token for a given username with a 72-hour expiration.
func GenerateJWT(username string) (string, error) {
	// Define token claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}