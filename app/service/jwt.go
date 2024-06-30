// Package: service
package service

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = os.Getenv("JWT_SECRET")

// EncrytData encrypts data and returns a token
func EncrytData(data map[string]interface{}) (string, error) {
	jwtExp, _ := strconv.Atoi(os.Getenv("JWT_EXP"))
	service := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp":  time.Now().Add(time.Hour * time.Duration(jwtExp)).Unix(),
			"iat":  time.Now().Unix(),
			"data": data,
		},
	)
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
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(jwtSecret), nil
		},
	)

	if err != nil {
		return nil, err
	}
	claims := service.Claims.(jwt.MapClaims)["data"].(map[string]interface{})
	return claims, nil
}
