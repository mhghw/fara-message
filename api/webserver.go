package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunWebServer(port int) error {
	addr := fmt.Sprintf(":%d", port)
	router := gin.Default()
	router.POST("/user/register", Register)
	router.Use(AuthMiddleware)
	err := router.Run(addr)
	return err
}
