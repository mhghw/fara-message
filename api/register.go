package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/jwt"
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

		log.Print(http.StatusBadRequest, err)
	}
	if requestBody.Password == requestBody.ConfirmPassword {
		token, err := jwt.CreateToken(requestBody.ID)
		if err != nil {
			panic("failed to create token")
		}
		c.JSON(http.StatusOK, token)
		fmt.Println(requestBody)
	} else {
		c.JSON(http.StatusNotAcceptable, "password does not match")
	}

}
