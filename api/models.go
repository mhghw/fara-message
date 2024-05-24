package api

import (
	"github.com/mhghw/fara-message/db"
)

type HTTPError struct {
	Message string `json:"message"`
}

func convertUserTableToRegisterForm(user db.UserTable) RegisterForm {
	var gender string
	switch user.Gender {
	case 0:
		gender = "Male"
	case 1:
		gender = "Female"
	case 2:
		gender = "NonBinary"
	}
	layout := "2006-01-02 15:04:05"
	formattedTime := user.DateOfBirth.Format(layout)
	result := RegisterForm{
		Username:        user.Username,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Password:        user.Password,
		ConfirmPassword: user.Password,
		Gender:          gender,
		DateOfBirth:     formattedTime,
	}
	return result
}
