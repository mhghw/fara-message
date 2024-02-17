package api

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

func hash(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	hashedBytes := hasher.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)
	return hashedString
}

func authenticateUser(c *gin.Context) {
	var loginData loginBody
	err := c.BindJSON(&loginData)

	if err != nil {
		fmt.Errorf("error binding JSON:%w", err)
		c.Status(400)
		return
	}

	if len(loginData.Username) < 3 || len(loginData.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "the username or password is incorrect",
		})
		return
	}

	//checking entered data with data that is already stored
	userUnderReveiw, err1 := db.UsersDB.GetUserByUsername(loginData.Username)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "the username or password is incorrect",
		})
	}
	if hash(loginData.Password) == userUnderReveiw.Password {
		c.JSON(http.StatusOK, gin.H{
			"message": "username and password are correct",
		})
	}
}
