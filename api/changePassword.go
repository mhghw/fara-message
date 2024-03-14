package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// suppose users post their usernames to edit information
func changePassword(c *gin.Context) {
	var userData UserData

	err := c.BindJSON(&userData.Username)
	if err != nil {
		log.Printf("error binding JSON:%v", err)
		c.Status(400)
		return
	}

	if len(userData.Username) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "the username is incorrect",
		})
		return
	}

	userUnderReview, err := db.UsersDB.GetUserByUsername(userData.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "the username is incorrect",
		})
		return
	}
	userUnderReview.Password = userData.Password
	err = db.UsersDB.UpdateUser(userUnderReview)
	if err != nil {
		log.Printf("error updating user:%v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "password change successfully",
	})
}
