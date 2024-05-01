package db

import (
	"time"
)

type Message struct {
	ID          int
	UserTableID int
	UserTable   UserTable
	ChatTableID int
	ChatTable   ChatTable
	Content     string
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
	Gender      int8
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

type Chat struct {
	ID          int
	Name        string
	CreatedTime time.Time
	DeletedTime time.Time
	Type        ChatType
}
type ChatTable struct {
	ID          int
	Name        string
	CreatedTime time.Time
	DeletedTime time.Time
	Type        int8
}
type ChatMember struct {
	UserTableID int
	UserTable   UserTable
	ChatTableID int
	ChatTable   ChatTable
	JoinedTime  time.Time
	LeftTime    time.Time
}
type ChatType struct {
	chatType int
}
type ChatIDAndChatName struct {
	ChatID   int
	ChatName string
}

func (c *ChatType) Int() int {
	return c.chatType
}

var (
	Direct  = ChatType{chatType: 0}
	Group   = ChatType{chatType: 1}
	Unknown = ChatType{chatType: -1}
)

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
	var gender int8
	switch user.Gender.gender {
	case 0:
		gender = 0
	case 1:
		gender = 1
	case 2:
		gender = 2
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

func ConvertChatToChatTable(chat Chat) ChatTable {
	var chatType int8
	switch chat.Type {
	case Direct:
		chatType = 0
	case Group:
		chatType = 1
	default:
		chatType = -1
	}

	result := ChatTable{
		ID:          chat.ID,
		Name:        chat.Name,
		CreatedTime: chat.CreatedTime,
		DeletedTime: chat.DeletedTime,
		Type:        chatType,
	}
	return result
}
