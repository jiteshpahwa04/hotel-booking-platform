package utils

import (
	config "AuthInGo/config/env"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error encrypting password:", err)
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(plainPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func CreateJWTToken(userID int64, payload *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(config.GetString("JWT_SECRET", "default-secret")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}