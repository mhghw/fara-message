package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunWebServer(port *int) error {
	addr := fmt.Sprintf(":%d", *port)
	router := gin.Default()
	router.Use(AuthMiddleware)
	router.POST("/user/register", Register)
	err := router.Run(addr)
	return err
}
