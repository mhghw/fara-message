package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{

			"error": "Authorization is required",
		})
		c.Abort()
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return GetSecretKey(), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		c.Abort()
	}
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token is not valid",
		})
		c.Abort()
		return
	}
	c.Next()
}
