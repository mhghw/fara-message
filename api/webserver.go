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
	router.POST("/login", login)
	router.POST("/user/edit",editUser) 
	router.Use(AuthMiddleware)
	err := router.Run(addr)
	return err
}
