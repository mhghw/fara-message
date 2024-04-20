package send_message

import "github.com/gin-gonic/gin"

func RunWebServer() {
	router := gin.New()
	router.POST("/send/message", sendMessage)
	router.Run()
}
