package delete_message

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "root:Yeganeh_2004@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&Message{})
}

func deleteMessage(c *gin.Context) {
	var messageID int
	if err := c.BindJSON(&messageID); err != nil {
		fmt.Errorf("error binding JSON:%w", err)
	}
	result := db.Delete(&Message{}, messageID)
	if result.Error != nil {
		fmt.Errorf("error:%w", result.Error)
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "message deleted successfully",
	})
}
