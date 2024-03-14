package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

type Information struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func editUser(c *gin.Context) {
	var userInfo Information

	err := c.BindJSON(&userInfo)
	if err != nil {
		fmt.Errorf("error binding JSON:%w", err)
		c.Status(400)
		return
	}

	if len(userInfo.Username) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "the username is incorrect",
		})
		return
	}

	userUnderReview, err := db.UsersDB.GetUserByUsername(userInfo.Username)
	if err != nil {
		fmt.Errorf("error getting username from database:%w", err)
		c.Status(400)
		return
	}

	userUnderReview.FirstName=userInfo.Firstname
	userUnderReview.LastName=userInfo.Lastname
	err = db.UsersDB.UpdateUser(userUnderReview)
	if err != nil {
		fmt.Errorf("error updating user:%w", err)
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "firstname and lastname edited successfully",
	})
}
