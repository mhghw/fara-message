package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

type ChatAPI struct {
	ChatName string    `json:"chatName"`
	ChatType string    `json:"chatType"`
	Users    []db.User `json:"users"`
}

func NewChatAPI(c *gin.Context) {
	var requestBody ChatAPI
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
