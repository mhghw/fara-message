package api

import "github.com/mhghw/fara-message/db"

type loginBody struct {
	Username string `json:username`
	Password string `json:password`
}

type HTTPError struct {
	Message string `json:message`
}

type NewChat struct {
	ChatName string    `json:"chatName"`
	ChatType string    `json:"chatType"`
	Users    []db.User `json:"users"`
}

type Chat struct {
	ID int
}
