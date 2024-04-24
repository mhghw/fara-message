package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

type Username struct {
	username string
}

func CreateUserHandler(c *gin.Context) {
	var newUser db.User
	err := c.BindJSON(&newUser)
	if err != nil {
		log.Printf("error binding JSON:%v", err)
		c.Status(400)
		return
	}
	err = db.CreateUser(newUser)
	if err != nil {
		log.Printf("error inserting user:%v", err)
		c.Status(400)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "user inserted successfully",
		})
	}

}

func ReadUserHandler(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		var username Username
		err := c.BindJSON(&username.username)
		if err != nil {
			log.Printf("error binding json:%v", err)
			c.Status(400)
			return
		}
		user, err := db.ReadAnotherUser(username.username)
		if err != nil {
			log.Printf("error reading user:%v", err)
			c.Status(400)
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"username":  user.Username,
				"firstname": user.FirstName,
				"lastname":  user.LastName,
			})
		}
	} else {
		parts := strings.Split(authorizationHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
			return
		}
		accessToken := parts[1]
		userID, err := getUserIDFromToken(accessToken)
		if err != nil {
			log.Printf("error:%v", err)
			c.Status(400)
			return
		}
		user, err := db.ReadUser(userID)
		if err != nil {
			log.Printf("error reading user:%v", err)
			c.Status(400)
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"username":      user.Username,
				"firstname":     user.FirstName,
				"lastname":      user.LastName,
				"gender":        user.Gender,
				"date of birth": user.DateOfBirth,
				"created time":  user.CreatedTime,
			})
		}
	}
}

func UpdateUserHandler(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	accessToken := parts[1]
	userID, err := getUserIDFromToken(accessToken)
	if err != nil {
		log.Printf("error:%v", err)
		c.Status(400)
		return
	}
	var newInfo db.User
	err = c.BindJSON(&newInfo)
	if err != nil {
		log.Printf("error binding JSON:%v", err)
		c.Status(400)
		return
	}
	err = db.UpdateUser(userID, newInfo)
	if err != nil {
		log.Printf("error updating user:%v", err)
		c.Status(400)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "user updated successfully",
		})
	}
}

func DeleteUserHandler(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	accessToken := parts[1]
	userID, err := getUserIDFromToken(accessToken)
	if err != nil {
		log.Printf("error:%v", err)
		c.Status(400)
		return
	}
	err = db.DeleteUser(userID)
	if err != nil {
		log.Printf("error:%v", err)
		c.Status(400)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "user deleted successfully",
		})
	}
}
