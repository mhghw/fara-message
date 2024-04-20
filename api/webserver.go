package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	send_message "github.com/mhghw/fara-message/message"
)

func RunWebServer(port int) error {
	addr := fmt.Sprintf(":%d", port)
	router := gin.New()
	router.POST("/user/register", Register)
	router.Use(AuthMiddleware)
	router.POST("/new_direct_chat", NewDirectChat)
	router.POST("/new_group_chat", NewGroupChat)
	router.GET("/chat/:id/messages", GetChatMessages)
	router.POST("/send/message", send_message.SendMessage)
	err := router.Run(addr)
	return err
}
