package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("JWT_SECRET_KEY")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func init() {
	if len(jwtKey) == 0 {
		panic("JWT_SECRET_KEY environment variable is not set")
	}
}

func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	if CheckIfTokenRevoked(tokenString) {
		return nil, fmt.Errorf("token has been revoked")
	}

	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func RefreshToken(oldTokenString string) (string, error) {
	claims, err := ValidateToken(oldTokenString)
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	newClaims := &Claims{
		Email: claims.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	return newToken.SignedString(jwtKey)
}
