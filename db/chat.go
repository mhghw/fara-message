package db

import (
	"errors"
	"log"
	"time"
)

type Chat struct {
	ChatID      int64 `gorm:"primary_key"`
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

func NewChat(chatName string, chatType string, user []User) error {
	chat := Chat{
		ChatName:    chatName,
		ChatType:    chatType,
		CreatedTime: time.Now(),
	}

	for _, u := range user {
		chatMember := ChatMember{
			JoinedTime: time.Now(),
			ChatID:     chat.ChatID,
			UserID:     u.ID,
		}

		if err := DB.Create(&chatMember).Error; err != nil {
			DB.Delete(&chatMember)
			return errors.New("cannot create chat member")

		}
	}

	DB.Create(&chat)
	return nil
}
