package CRUD_User

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

	db.AutoMigrate(&User{})
}

func createUser(c *gin.Context) {
	var newUser User
	err := c.BindJSON(&newUser)
	if err != nil {
		fmt.Errorf("error binding JSON:%w", err)
		c.Status(400)
		return
	}
	result := db.Create(&newUser)
	if result.Error != nil {
		fmt.Errorf("error inserting user:%w", result.Error)
		c.Status(400)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "user inserted successfully",
		})
	}

}

func readUser(c *gin.Context) {
	//read user with username
	var user User
	var username string
	err := c.BindJSON(&username)
	if err != nil {
		fmt.Errorf("error binding JSON:%w", err)
		c.Status(400)
		return
	}
	result := db.First(&user, "username=?", username)
	if result.Error != nil {
		fmt.Errorf("error reading user:%w", result.Error)
		c.Status(400)
		return
	} else {
		convertUserToJSON, err := json.Marshal(struct {
			User_ID       int
			Username      string    `json:"username"`
			Firstname     string    `json:"firstname"`
			Lastname      string    `json:"lastname"`
			Gender        string    `json:"gender"`
			Date_Of_Birth time.Time `json:"date_of_birthday"`
			Created_Time  time.Time `json:"created_time"`
		}{user.User_ID, user.Username, user.Firstname, user.Lastname, user.Gender, user.Date_Of_Birth, user.Created_Time})
		if err != nil {
			fmt.Errorf("error marshalling:%w", err)
			c.Status(400)
			return
		}
		c.JSON(http.StatusOK, convertUserToJSON)
	}
}

func updateUser(c *gin.Context){
	//read user with username
	var user User
	var username string
	err := c.BindJSON(&username)
	if err != nil {
		fmt.Errorf("error binding JSON:%w", err)
		c.Status(400)
		return
	}
	result := db.First(&user, "username=?", username)
	if result.Error != nil {
		fmt.Errorf("error reading user:%w", result.Error)
		c.Status(400)
		return
	} else {
		
	}
}

func deleteUser(c *gin.Context){
	//read user with username
	var user User
	var username string
	err := c.BindJSON(&username)
	if err != nil {
		fmt.Errorf("error binding JSON:%w", err)
		c.Status(400)
		return
	}
	result := db.First(&user, "username=?", username)
	if result.Error != nil {
		fmt.Errorf("error reading user:%w", result.Error)
		c.Status(400)
		return
	} else {
		result=db.Delete(&user)
		if result.Error != nil {
			fmt.Errorf("error deleting user:%w", result.Error)
			c.Status(400)
			return
		}else{
			c.JSON(http.StatusOK,gin.H{
				"message":"user deleted successfully",
			})
		}
	}
}
