package send_message

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("farawin")

const (
	TokenExpireTime = "expiration_time"
	TokenUserID     = "user_id"
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

func getUserIDFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims format")
	}
	userID := claims[TokenUserID].(string)
	expirationTime := claims[TokenExpireTime].(time.Time)
	if expirationTime.Before(time.Now()) {
		return "", errors.New("token has expired")
	}

	return userID, nil
}
