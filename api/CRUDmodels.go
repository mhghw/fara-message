package CRUD_User

import (
	"time"
)

type User struct {
	ID       int
	Username      string    `json:"username"`
	Firstname     string    `json:"firstname"`
	Lastname      string    `json:"lastname"`
	Password      string    `json:"passowrd"`
	Gender        string    `json:"gender"`
	Date_Of_Birth time.Time `json:"date_of_birthday"`
	Created_Time  time.Time `json:"created_time"`
}

type newInformation struct{
	Firstname     string    `json:"firstname"`
	Lastname      string    `json:"lastname"`
	Gender        string    `json:"gender"`
	Date_Of_Birth time.Time `json:"date_of_birthday"`
}
