package api

import (
	"github.com/gin-gonic/gin"
)

func RunLogin() {
	router := gin.Default()
	router.POST("/login", postChecking)
}
