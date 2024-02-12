package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegUser struct {
	ID              string `json:"id"`
	Username        string `json:"username"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Gender          string `json:"gender"`
	DateOfBirth     string `json:"date_of_birth"`
}

// handle func for register,receives a json string according to the RegUser struct and generate a JWT with ID parameter
func Register(c *gin.Context) {
	var requestBody RegUser
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Print("failed to bind json", err)
		return
	}
	if requestBody.Password != requestBody.ConfirmPassword {
		c.String(http.StatusBadRequest, "password does not match")
		return
	}
	token, err := CreateJWTToken(requestBody.ID)
	if err != nil {
		log.Print("failed to create token")
		return
	}
	c.JSON(http.StatusOK, token)
	fmt.Println(requestBody)
}
