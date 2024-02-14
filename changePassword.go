package changePassword

import(
	"github.com/gin-gonic/gin"
)

func StartWebServer() {
	router := gin.New()
	router.POST("/user/change_password",changePassword)
	router.Run()
}
