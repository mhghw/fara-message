package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
)

type RegisterForm struct {
	ID              string `json:"id"`
	Username        string `json:"username"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Gender          string `json:"gender"`
	DateOfBirth     string `json:"date_of_birth"`
}

// type Gender int8

// const(
// 	Male Gender = iota
// 	Female
// )

func convertRegisterFormToUser(form RegisterForm) (db.User, error) {
	convertTime, err := time.Parse("2006-01-02", form.DateOfBirth)
	if err != nil {
		return db.User{}, fmt.Errorf("failed to parse date %w", err)
	}
	var gender db.Gender
	switch strings.ToLower(form.Gender) {
	case "male":
		gender = db.Male
	case "female":
		gender = db.Female
	}
	user := db.User{
		ID:          form.ID,
		Username:    form.Username,
		FirstName:   form.FirstName,
		LastName:    form.LastName,
		Password:    form.Password,
		Gender:      gender,
		DateOfBirth: convertTime,
	}
	return user, nil
}

// handle func for register,receives a json string according to the RegisterForm struct and generate a JWT with ID parameter
func Register(c *gin.Context) {
	var requestBody RegisterForm
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Print("failed to bind json", err)
		return
	}
	if requestBody.Password != requestBody.ConfirmPassword {
		c.String(http.StatusBadRequest, "password does not match")
		return
	}
	token, err := CreateJWTToken(requestBody.ID)
	if err != nil {
		log.Print("failed to create token")
		return
	}
	user, err := convertRegisterFormToUser(requestBody)
	if err != nil {
		log.Print("failed to convert register form to user")
		return
	}
	db.UsersDB.CreateUser(user)
	// check, _ := db.UsersDB.GetUser(user.ID)

	c.JSON(http.StatusOK, check)

	c.JSON(http.StatusOK, token)
}
