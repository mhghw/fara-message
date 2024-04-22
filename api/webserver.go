package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunWebServer(port int) error {
	addr := fmt.Sprintf(":%d", port)
	router := gin.New()
	router.POST("/user/edit",editUser)
	router.POST("/user/register", Register)
	router.POST("/user/change_password",changePassword)
	// router.POST("/login", login)
	// router.POST("/login", login)
	router.Use(AuthMiddleware)
	router.POST("/new_direct_chat", NewDirectChat)
	router.POST("/new_group_chat", NewGroupChat)
	router.GET("/chat/:id/messages", GetChatMessages)
	router.POST("/send/message", SendMessageHandler)
	router.POST("/delete/message", DeleteMessageHandler)
	err := router.Run(addr)
	return err
}
