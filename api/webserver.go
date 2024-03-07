package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunWebServer(port int) error {
	addr := fmt.Sprintf(":%d", port)
	router := gin.New()
	router.POST("/user/register", Register)
	router.POST("/user/change_password",changePassword)
	router.POST("/login", login)
	router.Use(AuthMiddleware)
	err := router.Run(addr)
	return err
}
