package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

type UserData struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// suppose users post their usernames to edit information
func changePassword(c *gin.Context) {
	var user UserData

	err := c.BindJSON(&user)
	if err != nil {
		log.Printf("error binding JSON:%v", err)
		c.Status(400)
		return
	}

	if len(user.Username) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "the username is incorrect",
		})
		return
	}
	if user.Password != user.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password does not match",
		})
		return
	}

	userUnderReviewTable, err := db.Mysql.ReadUserByUsername(user.Username)
	if err != nil {
		log.Printf("error getting user from database:%v", err)
		c.Status(400)
		return
	}
	userUnderReviewTable.Password = user.Password

	err = db.Mysql.UpdateUser(userUnderReviewTable.ID, userUnderReviewTable)
	if err != nil {
		log.Printf("error updating user:%v", err)
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "password change successfully",
	})
}
