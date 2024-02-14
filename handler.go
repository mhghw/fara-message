package editUser

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

// suppose users post their usernames to edit information
func postEdit(c *gin.Context) {
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
	for i := 10001; i <= 10004; i++ {
		temp := strconv.Itoa(i)
		currentUser, err := db.UsersDB.GetUser(temp)
		if err != nil {
			fmt.Errorf("error:%w", err)
			return
		}
		if currentUser.Username == username {
			fmt.Println("enter new firstname:")
			fmt.Scan(&currentUser.FirstName)
			fmt.Println("enter new lastname:")
			fmt.Scan(&currentUser.LastName)
			db.UsersDB.UpdateUser(currentUser)
		}

	}
}
