package api

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/mhghw/fara-message/db"
)

type ChatResponse struct {
	ID       int       `json:"chatId"`
	Name     string    `json:"chatName"`
	Messages []Message `json:"messages"`
	Users    []User    `json:"users"`
}

type ChatRequest struct {
	ID   int    `json:"chatId"`
	Name string `json:"chatName"`
}
type NewGroupChatRequest struct {
	ChatName string    `json:"chatName"`
	Users    []db.User `json:"users"`
}
type NewDirectChatRequest struct {
	Users []db.User `json:"users"`
}

func NewDirectChatHandler(c *gin.Context) {
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

func NewGroupChatHandler(c *gin.Context) {
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

func GetChatMessagesHandler(c *gin.Context) {
	var requestBody ChatRequest
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

func GetUsersChatsHandler(c *gin.Context) {
	userIDString := c.Param("id")
	userID, _ := strconv.Atoi(userIDString)
	chatMembers, err := db.GetUsersChatMembers(userID)
	if err != nil {
		log.Print("failed to get users chat members")
		return
	}
	c.JSON(200, chatMembers)
}
