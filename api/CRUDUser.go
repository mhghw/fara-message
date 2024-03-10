package CRUD_User

import (
	"fmt"
	"net/http"
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
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader=="" {
		var username string
		err:=c.BindJSON(&username)
		if err!=nil {
			fmt.Errorf("error binding json:%w",err)
			c.Status(400)
			return
		}
		var user User
		result:=db.First(&user,"username=?",username)
		if result.Error != nil {
			fmt.Errorf("error reading user:%w", result.Error)
			c.Status(400)
			return
		} else {
			c.JSON(http.StatusOK,gin.H{
				"username":user.Username,
				"firstname":user.Firstname,
				"lastname":user.Lastname,
			})
		}
	}else{
		parts := strings.Split(authorizationHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
			return
		}
		accessToken := parts[1]
		userID, err := getUserIDFromToken(accessToken)
		if err != nil {
			fmt.Errorf("error:%w", err)
			c.Status(400)
			return
		}
		var user User
		result:=db.First(&user,"ID=?",userID)     //id or ID??????????????????
		if result.Error != nil {
			fmt.Errorf("error reading user:%w", result.Error)
			c.Status(400)
			return
		} else {
			c.JSON(http.StatusOK,gin.H{
				"username":user.Username,
				"firstname":user.Firstname,
				"lastname":user.Lastname,
				"gender":user.Gender,
				"date of birth":user.Date_Of_Birth,
				"created time":user.Created_Time,
			})
		}
	}
}

func updateUser(c *gin.Context) {
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
	userID, err1 := getUserIDFromToken(accessToken) //err1!!!!!!!!!!!!!!!!!
	if err1 != nil {
		fmt.Errorf("error:%w", err1)
		c.Status(400)
		return
	}

	var newInfo newInformation
	err2 := c.BindJSON(&newInfo)
	if err2 != nil {
		fmt.Errorf("error binding JSON:%w", err2)
		c.Status(400)
		return
	}
	result := db.Model(&User{}).Where("ID=?", userID).Updates(User{Firstname: newInfo.Firstname, Lastname: newInfo.Lastname, Gender: newInfo.Gender, Date_Of_Birth: newInfo.Date_Of_Birth})
	if result.Error != nil {
		fmt.Errorf("error updating user:%w", result.Error)
		c.Status(400)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "user updated successfully",
		})
	}
}

func deleteUser(c *gin.Context) {
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
	userID, err := getUserIDFromToken(accessToken) //err1!!!!!!!!!!!!!!!!!
	if err != nil {
		fmt.Errorf("error:%w", err)
		c.Status(400)
		return
	}

	var user User
	result:=db.First(&user,"ID=?",userID)     //id or ID??????????????????
	if result.Error != nil {
		fmt.Errorf("error reading user:%w", result.Error)
		c.Status(400)
		return
	} else {
		result = db.Delete(&user)
		if result.Error != nil {
			fmt.Errorf("error deleting user:%w", result.Error)
			c.Status(400)
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "user deleted successfully",
			})
		}
	}
}
