package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mhghw/fara-message/db"
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

func generateID() string {
	const charset = "0123456789"
	rand.NewSource(10)
	id := make([]byte, 5)
	for idx := range id {
		id[idx] = charset[rand.Intn(len(charset))]
	}

	return string(id)
}

// other validation fields will be added...
func validateUser(form RegisterForm) error {
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
	}
	return gender

}
func convertRegisterFormToUser(form RegisterForm) (db.User, error) {
	convertTime, err := time.Parse(time.DateOnly, form.DateOfBirth)

	if err != nil {
		return db.User{}, fmt.Errorf("failed to parse date %w", err)
	}
	gender := assignGender(form.Gender)
	generatedID := generateID()
	user := db.User{
		ID:          generatedID,
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
	err = validateUser(requestBody)
	if err != nil {
		log.Print("failed to validate user", err)
		return
	}

	user, err := convertRegisterFormToUser(requestBody)
	if err != nil {
		log.Print("failed to convert register form to user")
		return
	}

	token, err := CreateJWTToken(user.ID)
	if err != nil {
		log.Print("failed to create token")
		return
	}
	userToken := tokenJSON{
		Token: token,
	}
	userTokenJSON, err := json.Marshal(userToken)
	if err != nil {
		log.Print("failed to marshal token")
		return
	}

	db.UsersDB.CreateUser(user)
	c.JSON(http.StatusOK, userTokenJSON)
}
