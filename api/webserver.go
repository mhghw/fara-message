package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunWebServer(port int) error {
	addr := fmt.Sprintf(":%d", port)
	router := gin.New()
	router.POST("/user/register", Register)
	router.Use(AuthMiddleware)
	router.POST("/user/newchat", NewChatAPI)
	err := router.Run(addr)
	return err
}
