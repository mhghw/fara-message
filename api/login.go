package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(c *gin.Context) {
	var loginBody loginBody
	err := c.BindJSON(&loginBody)
	if err != nil {
		log.Printf("error binding JSON:%v", err)
		c.Status(400)
		return
	}

	errIncorrectUserOrPass := HTTPError{Message: "the username or password is incorrect"}
	errIncorrectUserOrPassJSON, errInMarshalling := json.Marshal(errIncorrectUserOrPass)
	if errInMarshalling != nil {
		log.Printf("error:%v", errInMarshalling)
		return
	}

	if len(loginBody.Username) < 3 || len(loginBody.Password) < 8 {
		c.JSON(http.StatusBadRequest, string(errIncorrectUserOrPassJSON))
		return
	}

	//checking entered data with data that is already stored
	userUnderReview, err := db.Mysql.ReadUserByUsername(loginBody.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, string(errIncorrectUserOrPassJSON))
		return
	}
	if hash(loginBody.Password) != userUnderReview.Password {
		c.JSON(http.StatusOK, gin.H{
			"message": "username or password are incorrect",
		})
		return
	}

	token, err := CreateJWTToken(userUnderReview.ID)
	if err != nil {
		log.Printf("failed to create token: %v", err)
		return
	}

	c.JSON(http.StatusOK, token)
}
