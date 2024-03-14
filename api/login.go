package api

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

type loginBody struct {
	Username string `json:username`
	Password string `json:password`
}

type HTTPError struct{
	Message string `json:message`
}

func hash(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	hashedBytes := hasher.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)
	return hashedString
}

func login(c *gin.Context) {
	var loginData loginBody
	err := c.BindJSON(&loginData)
	if err != nil {
		fmt.Errorf("error binding JSON:%w", err)
		c.Status(400)
		return
	}

	errIncorrectUserOrPass := HTTPError{Message: "the username or password is incorrect"}
	errIncorrectUserOrPassJSON, errInMarshalling := json.Marshal(errIncorrectUserOrPass)
	if errInMarshalling!=nil {
		fmt.Errorf("error:%w",errInMarshalling)
		return
	}

	if len(loginData.Username) < 3 || len(loginData.Password) < 8 {
		c.JSON(http.StatusBadRequest,errIncorrectUserOrPassJSON)
		return
	}

	//checking entered data with data that is already stored
	userUnderReveiw, err := db.UsersDB.GetUserByUsername(loginData.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest,errIncorrectUserOrPassJSON)
	}
	if hash(loginData.Password) == userUnderReveiw.Password {
		c.JSON(http.StatusOK, gin.H{
			"message": "username and password are correct",
		})
	}
}
