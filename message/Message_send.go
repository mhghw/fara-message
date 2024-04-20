package send_message

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

func SendMessage(c *gin.Context) {
	var messageInfo MessageInformation
	if err := c.BindJSON(&messageInfo); err != nil {
		fmt.Errorf("error binding json:", err)
		c.Status(400)
		return
	}

	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unauthorized",
		})
		return
	}
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unauthorized",
		})
		return
	}
	accessToken := parts[1]
	userID, err := getUserIDFromToken(accessToken)
	if err != nil {
		fmt.Errorf("error get ID:%w", err)
		c.Status(400)
		return
	}
	ID, _ := strconv.Atoi(userID)
	messageInfo.SenderID = ID

	result := db.DB.Create(&messageInfo)
	if result.Error != nil {
		fmt.Errorf("error:%w", result.Error)
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "message sent successfully",
	})
}
