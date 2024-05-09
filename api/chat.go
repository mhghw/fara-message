package api

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/mhghw/fara-message/db"
)

type ChatRequest struct {
	ID   int    `json:"chatId"`
	Name string `json:"chatName"`
}
type GroupChatRequest struct {
	ChatName string         `json:"chatName"`
	Users    []db.UserTable `json:"users"`
}
type DirectChatRequest struct {
	Users []UsernameType `json:"users"`
}

func NewDirectChatHandler(c *gin.Context) {
	var requestBody DirectChatRequest

	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Print("failed to bind json, ", err)
		return
	}
	tokenString := c.GetHeader("Authorization")

	userID, err := ValidateToken(tokenString)
	if err != nil {
		log.Printf("failed to find user by token: %v", err)
	}
	userTable := []db.UserTable{}
	for _, v := range requestBody.Users {
		user, err := db.Mysql.ReadUserByUsername(v.Username)
		if err != nil {
			log.Printf("failed to read user: %v", err)
			return
		}
		userTable = append(userTable, user)
	}

	validUser := false
	for _, user := range userTable {
		if user.ID == userID {
			validUser = true
		}
	}
	if !validUser {
		log.Printf("you're not allowed")
		c.JSON(400, "Invalid token")
		return
	}
	if len(userTable) == 0 {
		log.Print("failed to create chat: no users provided")
		return
	}
	log.Println(userTable)

	if err := db.Mysql.NewChat("", db.Direct, userTable); err != nil {
		log.Print("failed to create chat, ", err)
		return
	}
	log.Print("direct chat created")
}

func NewGroupChatHandler(c *gin.Context) {
	var requestBody GroupChatRequest
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Print("failed to bind json, ", err)
		return
	}

	if err := db.Mysql.NewChat(requestBody.ChatName, db.Group, requestBody.Users); err != nil {
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
	messages, err := db.Mysql.GetChatMessages(int64(chatID))
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
	chatMembers, err := db.Mysql.GetUsersChatMembers(userID)
	if err != nil {
		log.Printf("failed to get users chats: %v", err)
		return
	}

	result, err := db.Mysql.GetUsersChatIDAndChatName(chatMembers)
	if err != nil {
		log.Printf("failed to get users chats: %v", err)
		return
	}
	c.JSON(200, result)
}
