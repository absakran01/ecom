package auth

import (
	"fmt"
	"time"

	"github.com/absakran01/ecom/config"
	"github.com/dgrijalva/jwt-go"
)

func GenJWT(userId int) (string, error) {
	// debug 
	fmt.Println("Generating JWT for user ID:", userId)


	// Validate inputs
	if len(config.Envs.JWTSecret) == 0 {
		return "", fmt.Errorf("secret key is empty")
	}
	if userId <= 0 {
		return "", fmt.Errorf("invalid user ID: %d", userId)
	}

	// Create a new JWT token with claims
	expireTime := time.Duration(config.Envs.JWTExpiration) * time.Second
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userId,
		"exp":  time.Now().Add(expireTime).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString(config.Envs.JWTSecret)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return tokenString, nil
}
