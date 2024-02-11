package api

import (
	"net/http"
	"github.com/mhghw/fara-message/db"
	"github.com/gin-gonic/gin"
)

func postChecking(c *gin.Context) {

	var loginData loginBody
	err := c.BindJSON(&loginData)
	if err != nil {
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
	for _, value := range db.sampleUsers() {
		if value.Username == loginData.Username {
			if value.Password == loginData.Password {
				c.IndentedJSON(http.StatusOK, gin.H{
					"message": "The username and password are correct. You are now logged in.",
				})
				return
			}
		}
	}

	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"error": "the username or password is incorrect",
	})
}
