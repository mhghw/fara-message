package api

import "github.com/mhghw/fara-message/db"

type loginBody struct {
	Username string `json:username`
	Password string `json:password`
}

type HTTPError struct {
	Message string `json:message`
}

type ChatResponse struct {
	ID       int       `json:"chatId"`
	Name     string    `json:"chatName"`
	Messages []Message `json:"messages"`
	Users    []User    `json:"users"`
}
