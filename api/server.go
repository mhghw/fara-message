package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	send_message "github.com/mhghw/fara-message/message"
)

func RunWebServer(port int) error {
	addr := fmt.Sprintf(":%d", port)
	router := gin.New()
	router.POST("/user/register", RegisterHandler)
	router.POST("/login", login)
	router.POST("/login", login)
	router.Use(AuthMiddlewareHandler)
	router.POST("/user/change_password", changePassword)
	router.POST("/user/edit", editUser)
	router.POST("/new_direct_chat", NewDirectChatHandler)
	router.POST("/new_group_chat", NewGroupChatHandler)
	router.GET("/chat/:id/messages", GetChatMessagesHandler)
	router.POST("/send/message", send_message.SendMessage)
	err := router.Run(addr)
	return err
}
