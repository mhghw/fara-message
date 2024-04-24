package api

import (
	"github.com/mhghw/fara-message/db"
	"time"
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

type AnotherUserInfo struct{
	Username    string
	FirstName   string
	LastName    string
}

type Gender int8
const (
	Male Gender = iota
	Female
)
type UserInfo struct{
	Username    string
	FirstName   string
	LastName    string
	Gender      Gender
	DateOfBirth time.Time
	CreatedTime time.Time
}
