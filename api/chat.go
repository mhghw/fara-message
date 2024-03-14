package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/mhghw/fara-message/db"
)

func NewDirectChatRequest(c *gin.Context) {
	var requestBody NewDirectChat
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Print("failed to bind json, ", err)
		return
	}

	if err := db.NewChat("", db.Direct, requestBody.Users); err != nil {
		log.Print("failed to create chat, ", err)
		return
	}
}

func NewGroupChatRequest(c *gin.Context) {
	var requestBody NewGroupChat
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Print("failed to bind json, ", err)
		return
	}

	if err := db.NewChat(requestBody.ChatName, db.Group, requestBody.Users); err != nil {
		log.Printf("failed to create chat: %v", err)
		return
	}
}

func GetChatMessagesAPI(c *gin.Context) {
	var requestBody Chat
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Printf("failed to bind json: %v", err)
		return
	}
	chatID := requestBody.ID
	messages, err := db.GetChatMessages(int64(chatID))
	if err != nil {
		log.Print(err)
		return
	}
	chatMessages, err := json.Marshal(messages)
	if err != nil {
		log.Print("failed to marshal json", err)
		return
	}
	c.JSON(200, chatMessages)

}
