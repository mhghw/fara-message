package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

type UsernameType struct {
	Username string `json:"username"`
}

func ReadUserHandler(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	userID, err := ValidateToken(authorizationHeader)
	if err != nil {
		log.Printf("failed to validate token: %v", err)
		return
	}
	userTable, err := db.Mysql.ReadUser(userID)
	if err != nil {
		log.Printf("failed to read user: %v", err)
		return
	}
	RegisterForm := convertUserTableToRegisterForm(userTable)
	RegisterForm.Password = ""
	RegisterForm.ConfirmPassword = ""
	c.JSON(http.StatusOK, RegisterForm)

}

func UpdateUserHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	userID, err := ValidateToken(tokenString)
	if err != nil {
		log.Printf("error validating user: %v", err)
		c.JSON(400, "error validating user")

		return
	}
	oldUser, err := db.Mysql.ReadUser(userID)
	if err != nil {
		log.Printf("error reading user:%v", err)
		c.JSON(400, "error reading user")
		return
	}

	fmt.Println(userID)
	fmt.Println(oldUser)

	oldUserRegisterForm := convertUserTableToRegisterForm(oldUser)
	var newInfoRequest RegisterForm
	err = c.BindJSON(&newInfoRequest)
	if err != nil {
		log.Printf("error binding JSON:%v", err)
		c.Status(400)
		return
	}
	newInfoRequest.DateOfBirth = oldUserRegisterForm.DateOfBirth

	user, err := convertRegisterFormToUser(newInfoRequest)
	if err != nil {
		log.Printf("error converting registerForm to user:%v", err)
		c.Status(400)
		return
	}
	newUserTable := db.ConvertUserToUserTable(user)
	fmt.Println(newUserTable)
	err = db.Mysql.UpdateUser(userID, newUserTable)
	if err != nil {
		log.Printf("error updating user:%v", err)
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
	})

}

func DeleteUserHandler(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	userID, err := ValidateToken(authorizationHeader)
	if err != nil {
		log.Printf("error validating token: %v", err)
		c.Status(400)
		return
	}

	err = db.Mysql.DeleteUser(userID)
	if err != nil {
		log.Printf("failed to delete user:%v", err)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("failed to delete user:%v", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})

}
