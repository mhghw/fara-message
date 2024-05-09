package api

import (
	"log"
	"strconv"
	"time"

	"github.com/mhghw/fara-message/db"
)

type HTTPError struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	ID       int        `json:"chatId"`
	Name     string     `json:"chatName"`
	Messages []Message  `json:"messages"`
	Users    []UserInfo `json:"users"`
}

type AnotherUserInfo struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type UserInfo struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Gender      db.Gender `json:"gender"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	CreatedTime time.Time `json:"createdTime"`
}

func convertUserInfo(newInfo UserInfo) db.UserInfo {

	return db.UserInfo{
		Username:    newInfo.Username,
		FirstName:   newInfo.FirstName,
		LastName:    newInfo.LastName,
		Gender:      newInfo.Gender,
		DateOfBirth: newInfo.DateOfBirth,
		CreatedTime: newInfo.CreatedTime,
	}
}

func ConvertUserTableToUserInfo(userTable db.UserTable) UserInfo {
	var gender db.Gender
	switch userTable.Gender {
	case 0:
		gender = db.Male
	case 1:
		gender = db.Female
	case 2:
		gender = db.NonBinary

	}
	userInfo := UserInfo{
		ID:          string(userTable.ID),
		Username:    userTable.Username,
		FirstName:   userTable.FirstName,
		LastName:    userTable.LastName,
		Gender:      gender,
		DateOfBirth: userTable.DateOfBirth,
		CreatedTime: userTable.CreatedTime,
	}
	return userInfo
}
func ConvertUserInfoToUserTable(user UserInfo) db.UserTable {
	var gender int8
	switch user.Gender {
	case db.Male:
		gender = 0
	case db.Female:
		gender = 1
	case db.NonBinary:
		gender = 2
	}
	userID, err := strconv.Atoi(user.ID)
	if err != nil {
		log.Printf("Error converting user: %v", err)
	}
	userTable := db.UserTable{
		ID:          userID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Gender:      gender,
		DateOfBirth: user.DateOfBirth,
		CreatedTime: user.CreatedTime,
	}
	return userTable
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
