// Package: service
package service

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = os.Getenv("JWT_SECRET")

// EncrytData encrypts data and returns a token
func EncrytData(data map[string]interface{}) (string, error) {
	service := jwt.New(jwt.SigningMethodHS256)
	claims := service.Claims.(jwt.MapClaims)
	for key, value := range data {
		claims[key] = value
	}
	token, err := service.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// DecryptData decrypts token and returns the data
func DecryptData(token string) (map[string]interface{}, error) {
	service, err := jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims := service.Claims.(jwt.MapClaims)
	return claims, nil
}
