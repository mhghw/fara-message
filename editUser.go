package editUser

import(
	"github.com/gin-gonic/gin"
)

func StartWebServer(){
	router:=gin.New()
	router.POST("/user/edit",postEdit)  
	router.Run() 
}