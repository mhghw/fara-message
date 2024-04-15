package db

import (
	"time"
)

type ChatType struct {
	chatType int8
}

func (c *ChatType) Int() int8 {
	return c.chatType
}

var (
	Direct  = ChatType{0}
	Group   = ChatType{1}
	Unknown = ChatType{-1}
)

// func FromInt(i int8) (ChatType, error) {
// 	switch i {
// 	case Direct.chatType:
// 		return Direct, nil
// 	case Group.chatType:
// 		return Group, nil
// 	}
// 	return Unknown, errors.New("invalid chat type")

// }

type Chat struct {
	ID          int64  `gorm:"primary_key"`
	Name        string `gorm:"chat_name;default:' '"`
	CreatedTime time.Time
	DeletedTime time.Time
	Type        ChatType
}

type ChatMember struct {
	UserID     int64 `gorm:"foreign_key"`
	User       User
	ChatID     int64 `gorm:"foreign_key"`
	Chat       Chat
	JoinedTime time.Time
	LeftTime   time.Time
}
type MessageInformation struct {
	SenderID int    `json:"senderID"`
	ChatID   int    `json:"chatID"`
	Content  string `json:"content"`
}
