package send_message

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "root:Yeganeh_2004@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&MessageInformation{})
}

func sendMessage(c *gin.Context) {
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

	result := db.Create(&messageInfo)
	if result.Error != nil {
		fmt.Errorf("error:%w",result.Error)
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "message sent successfully",
	})
}
