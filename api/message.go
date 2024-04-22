package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mhghw/fara-message/db"
)

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

type Message struct {
	ID       int    `json:"id`
	SenderID int    `json:"senderID"`
	ChatID   int    `json:"chatID"`
	Content  string `json:"content"`
}

func DeleteMessageHandler(c *gin.Context) {
	var message Message ///////////////////////////////struct beshe!!!!!!!!!
	if err := c.BindJSON(&message.ID); err != nil {
		fmt.Errorf("error binding JSON:%w", err)
	}
	err := db.DeleteMessage(message.ID)
	if err != nil {
		fmt.Errorf("error:%w", err)
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "message deleted successfully",
	})
}

func SendMessageHandler(c *gin.Context) {
	var message Message
	if err := c.BindJSON(&message); err != nil {
		fmt.Errorf("error binding json:", err)
		c.Status(400)
		return
	}

	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unauthorized",
		})
		return
	}
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unauthorized",
		})
		return
	}
	accessToken := parts[1]
	userID, err := getUserIDFromToken(accessToken)
	if err != nil {
		fmt.Errorf("error get ID:%w", err)
		c.Status(400)
		return
	}
	ID, _ := strconv.Atoi(userID)
	message.SenderID = ID

	err = db.SendMessage(message.ID, message.SenderID, message.ChatID, message.Content)
	if err != nil {
		fmt.Errorf("error:%w", err)
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "message sent successfully",
	})
}
