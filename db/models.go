package db

import "time"

type Chat struct {
	ID          int64 `gorm:"primary_key"`
	ChatName    string
	CreatedTime time.Time
	DeletedTime time.Time
	ChatType    string
}

type ChatMember struct {
	UserID     int64 `gorm:"foreign_key"`
	User       User
	ChatID     int64 `gorm:"foreign_key"`
	Chat       Chat
	JoinedTime time.Time
	LeftTime   time.Time
}
