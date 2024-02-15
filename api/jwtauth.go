package api

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("farawin")

const (
	TokenExpireTime = "exp"
	TokenUserID     = "id"
)

func CreateJWTToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		TokenUserID:     userID,
		TokenExpireTime: time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}
	return tokenString, nil

}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims format")
	}
	userID := claims[TokenUserID].(string)
	expirationTime := claims[TokenExpireTime].(time.Time)
	if expirationTime.Before(time.Now()) {
		return "", fmt.Errorf("token has expired")
	}

	return userID, nil
}
