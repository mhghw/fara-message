package api

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/mhghw/fara-message/db"
)

func NewDirectChat(c *gin.Context) {
	var requestBody NewDirectChatRequest
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

func NewGroupChat(c *gin.Context) {
	var requestBody NewGroupChatRequest
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

func GetChatMessages(c *gin.Context) {
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

func GetUsersChats(c *gin.Context) {
	userIDString := c.Param("id")
	userID, _ := strconv.Atoi(userIDString)
	chatMembers, err := db.GetUsersChatMembers(userID)
	if err != nil {
		log.Print("failed to get users chat members")
		return
	}
	c.JSON(200, chatMembers)
}
