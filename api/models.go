package api

import (
	"time"

	"github.com/mhghw/fara-message/db"
)

type HTTPError struct {
	Message string `json:"message"`
}

type NewGroupChatRequest struct {
	ChatName string    `json:"chatName"`
	Users    []db.User `json:"users"`
}
type NewDirectChatRequest struct {
	Users []db.User `json:"users"`
}

type Chat struct {
	ID int
}

type AnotherUserInfo struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Gender int8

const (
	Male Gender = iota
	Female
)

type UserInfo struct {
	Username    string    `json:"username"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Gender      Gender    `json:"gender"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	CreatedTime time.Time `json:"createdTime"`
}
