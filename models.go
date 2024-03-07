package CRUD_User

import (
	"time"
)

type User struct {
	User_ID       int
	Username      string    `json:"username"`
	Firstname     string    `json:"firstname"`
	Lastname      string    `json:"lastname"`
	Password      string    `json:"passowrd"`
	Gender        string    `json:"gender"`
	Date_Of_Birth time.Time `json:"date_of_birthday"`
	Created_Time  time.Time `json:"created_time"`
}
