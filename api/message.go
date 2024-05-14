package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

type Message struct {
	ID       int    `json:"id"`
	SenderID int    `json:"senderID"`
	ChatID   int    `json:"chatID"`
	Content  string `json:"content"`
}

func SendMessageHandler(c *gin.Context) {
	var message Message
	if err := c.BindJSON(&message); err != nil {
		log.Printf("error binding json:%v", err)
		c.Status(400)
		return
	}

	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		log.Print("failed to get authorization token")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unauthorized",
		})
		return
	}
	userID, err := ValidateToken(authorizationHeader)
	if err != nil {
		log.Printf("error get ID:%v", err)
		c.Status(400)
		return
	}

	message.SenderID = userID

	err = db.Mysql.SendMessage(message.SenderID, message.ChatID, message.Content)
	if err != nil {
		log.Printf("error:%v", err)
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "message sent successfully",
	})
}

func DeleteMessageHandler(c *gin.Context) {
	var message Message
	if err := c.BindJSON(&message.ID); err != nil {
		log.Printf("error binding JSON:%v", err)
	}
	err := db.Mysql.DeleteMessage(message.ID)
	if err != nil {
		log.Printf("error:%v", err)
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "message deleted successfully",
	})
}
