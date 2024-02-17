package api

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

func hashString(input string) string {
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "the username or password is incorrect",
		})
		return
	}
	//checking entered data with data that is already stored
	for i := 10001; i <= 10004; i++ {
		temp := strconv.Itoa(i)
		currentUser, err := db.UsersDB.GetUser(temp)
		if err != nil {
			fmt.Errorf("error:%w", err)
			return
		}
		loginData.Password = hashString(loginData.Password)
		if loginData.Username == currentUser.Username && loginData.Password == currentUser.Password {
			c.JSON(http.StatusOK, gin.H{
				"message": "username and password are correct",
			})
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "the username or password is incorrect",
	})
}
