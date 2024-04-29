package db

import (
	"time"
)

type Message struct {
	ID      int `gorm:"primary_key"`
	UserID  int
	User    User
	ChatID  int
	Chat    Chat
	Content string
}

type Gender struct {
	gender int
}

var (
	Male      = Gender{gender: 0}
	Female    = Gender{gender: 1}
	NonBinary = Gender{gender: 2}
)

type User struct {
	ID          string `gorm:"primary_key"`
	Username    string
	FirstName   string
	LastName    string
	Password    string
	Gender      Gender
	DateOfBirth time.Time
	CreatedTime time.Time
	DeletedTime time.Time
}
type UserTable struct {
	ID          string `gorm:"primary_key"`
	Username    string
	FirstName   string
	LastName    string
	Password    string
	Gender      string
	DateOfBirth time.Time
	CreatedTime time.Time
	DeletedTime time.Time
}

type UserInfo struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Gender      Gender    `json:"gender"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	CreatedTime time.Time `json:"createdTime"`
}

func ConvertUserToUserInfo(user User) UserInfo {
	return UserInfo{
		ID:          user.ID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Gender:      user.Gender,
		DateOfBirth: user.DateOfBirth,
		CreatedTime: user.CreatedTime,
	}
}

func ConvertUserToUserTable(user User) UserTable {
	var gender string
	switch user.Gender.gender {
	case 0:
		gender = "Male"
	case 1:
		gender = "Female"
	case 2:
		gender = "Non binary"
	}
	userTable := UserTable{
		ID:          user.ID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Password:    user.Password,
		Gender:      gender,
		DateOfBirth: user.DateOfBirth,
		CreatedTime: user.CreatedTime,
		DeletedTime: user.DeletedTime,
	}
	return userTable
}
