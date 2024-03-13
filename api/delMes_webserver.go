package delete_message

import "github.com/gin-gonic/gin"

func RunWebServer() {
	router := gin.New()
	router.POST("/delete/message", deleteMessage)
	router.Run()
}
