package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func WebServer(port *int) {

	addr := fmt.Sprintf(":%d", *port)
	router := gin.Default()
	router.Use(AuthMiddleware)
	router.POST("/user/register", Register)
	router.Run(addr)
}
