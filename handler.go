package editUser

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

func editUser(c *gin.Context) {
	var username string

	err := c.BindJSON(&username)
	if err != nil {
		fmt.Errorf("error binding JSON:%w", err)
		c.Status(400)
		return
	}

	if len(username) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "the username or password is incorrect",
		})
		return
	}

	userUnderReview, err1 := db.UsersDB.GetUserByUsername(username)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "the username is incorrect",
		})
	}
	fmt.Println("enter new firstname:")
	fmt.Scan(&userUnderReview.FirstName)
	fmt.Println("enter new lastname:")
	fmt.Scan(&userUnderReview.LastName)
	err2 := db.UsersDB.UpdateUser(userUnderReview)
	if err2 != nil {
		fmt.Errorf("error:%w", err2)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "firstname and lastname edited successfully",
	})
}
