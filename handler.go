package changePassword

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

func hashString(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	hashedBytes := hasher.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)
	return hashedString
}

// suppose users post their usernames to edit information
func changePassword(c *gin.Context) {
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
			var newPassword string
			fmt.Println("enter new password:")
			fmt.Scan(&newPassword)
			newPassword = hashString(newPassword)
			currentUser.Password = newPassword
			err := db.UsersDB.UpdateUser(currentUser)
			if err != nil {
				fmt.Errorf("error:%w", err)
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "password change successfully",
			})
		}
	}
}
