package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
	"gorm.io/gorm"
)

type RegisterForm struct {
	Username        string `json:"username"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Gender          string `json:"gender"`
	DateOfBirth     string `json:"date_of_birth"`
}

type tokenJSON struct {
	Token string `json:"token"`
}

func RegisterHandler(c *gin.Context) {
	var requestBody RegisterForm
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Print("failed to bind json", err)
		return
	}
	err = validateUser(requestBody)
	if err != nil {
		log.Print("failed to validate user: ", err)
		return
	}

	user, err := convertRegisterFormToUser(requestBody)
	if err != nil {
		log.Print("failed to convert register form to user")
		return
	}
	repeatedUserName, err := CheckRepeatedUser(user.Username)
	if err != nil {
		log.Printf("failed to check for repeated user: %v", err)
		return
	}
	if repeatedUserName {
		log.Printf("repeated userName")
		c.JSON(400, "repeated username")
		return
	}
	token, err := CreateJWTToken(user.ID.String())
	if err != nil {
		log.Print("failed to create token")
		return
	}
	userToken := tokenJSON{
		Token: token,
	}

	db.Mysql.CreateUser(user)
	c.JSON(http.StatusOK, userToken.Token)
}

func validateUser(form RegisterForm) error {
	if len(form.Username) < 4 || form.Username == "" {
		return errors.New("username length must be more than 5 characters")
	}
	if len(form.FirstName) < 3 || form.FirstName == "" {
		return errors.New("first name length must be more than 3 characters")
	}
	if len(form.LastName) < 3 || form.LastName == "" {
		return errors.New("last name length must be more than 3 characters")
	}
	if form.Gender == "" {
		return errors.New("please fill the gender section")
	}

	if strings.ToLower(form.Gender) != "male" && strings.ToLower(form.Gender) != "female" && strings.ToLower(form.Gender) != "non binary" {
		return errors.New("wrong gender type ")
	}
	if form.DateOfBirth == "" {
		return errors.New("please fill the date of birth section")
	}

	if len(form.Password) < 8 {
		return errors.New("password is too short")
	}
	if form.Password != form.ConfirmPassword {
		return errors.New("password does not match")
	}

	return nil
}
func assignGender(sex string) db.Gender {
	var gender db.Gender
	switch strings.ToLower(sex) {
	case "male":
		gender = db.Male
	case "female":
		gender = db.Female
	case "non binary":
		gender = db.NonBinary
	}
	return gender

}
func convertRegisterFormToUser(form RegisterForm) (db.User, error) {
	layout := "2006-01-02 15:04:05"
	convertTime, err := time.Parse(layout, form.DateOfBirth)
	if err != nil {
		return db.User{}, fmt.Errorf("failed to parse date %w", err)
	}

	gender := assignGender(form.Gender)
	generatedID := generateID()
	password := hash(form.Password)
	user := db.User{
		ID:          generatedID,
		Username:    form.Username,
		FirstName:   form.FirstName,
		LastName:    form.LastName,
		Password:    password,
		Gender:      gender,
		DateOfBirth: convertTime,
		CreatedTime: time.Now(),
		DeletedTime: time.Time{},
	}

	return user, nil
}

func CheckRepeatedUser(username string) (bool, error) {
	result := true
	var err error
	err = nil
	_, err = db.Mysql.ReadUserByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result = false
			err = nil
			return result, err
		}

		log.Printf("failed to get user: %v", err)
		return result, err
	}
	return result, err
}

// func generateID() (string, error) {

// }
