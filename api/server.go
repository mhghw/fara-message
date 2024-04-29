package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunWebServer(port int) error {
	addr := fmt.Sprintf(":%d", port)
	router := gin.New()
	router.POST("/user/register", RegisterHandler)
	router.POST("/change_password", changePassword)
	router.POST("/login", authenticateUser)
	router.Use(AuthMiddlewareHandler)
	router.POST("/user/info", ReadUserHandler)
	router.POST("/user/update", UpdateUserHandler)
	router.POST("/user/delete", DeleteUserHandler)
	router.POST("/user/edit", editUser)
	router.POST("/send/message", SendMessageHandler)
	router.DELETE("/delete/message", DeleteMessageHandler)
	router.POST("/chat/direct", NewDirectChatHandler)
	router.POST("/chat/group", NewGroupChatHandler)
	router.GET("/chat/:id/messages", GetChatMessagesHandler)
	err := router.Run(addr)
	return err
}
