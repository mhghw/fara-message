package api

import (
	"os"

	"github.com/gin-gonic/gin"
)

func getPort() string {
	if len(os.Args) >= 2 {
		return os.Args[1]
	}
	return "8080"
}

func StartWebServer() {
	router := gin.New()
	router.POST("/login", authenticateUser)
	port := getPort()
	router.Run(port)
}
