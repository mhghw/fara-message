package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/mhghw/fara-message/db"
)

func NewChatAPI(c *gin.Context) {
	var requestBody NewChat
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Print("failed to bind json", err)
		return
	}
	if err := db.NewChat(requestBody.ChatName, requestBody.ChatType, requestBody.Users); err != nil {
		log.Print("failed to create chat", err)
		return
	}
}

func GetChatMessagesAPI(c *gin.Context) {
	var requestBody Chat
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Print("failed to bind json", err)
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
