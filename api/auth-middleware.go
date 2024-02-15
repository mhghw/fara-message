package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{

			"error": "Authorization is required",
		})

		c.Abort()
		return
	}
	userID, err := ValidateToken(tokenString)
	if err != nil {
		log.Print("failed to validate token")
		c.Abort()
		return
	}
	_, err = db.UsersDB.GetUser(userID)
	if err != nil {
		log.Print("user ID is not in the DataBase: %w", err)
		c.Abort()
		return
	}

	c.Next()
}
