package CRUD_User

import(
	"github.com/gin-gonic/gin"
)

func RunWebServer(){
	router:=gin.New()
	router.POST("/user/create",createUser)
	router.POST("/user/read",readUser)
	router.POST("/user/update",updateUser)
	router.POST("/user/delete",deleteUser)
	router.Run()
}

