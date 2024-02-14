package api

import (
	"github.com/gin-gonic/gin"
)

func StartWebServer() {
	router := gin.New()
	router.POST("/login", authenticateUser)
	router.Run()
}
