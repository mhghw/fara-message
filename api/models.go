package api

import "github.com/mhghw/fara-message/db"

// type loginBody struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

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
