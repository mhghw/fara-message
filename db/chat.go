package db

import (
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

func NewChat(chatName string, chatType string, user []User) {
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
			log.Print("cannot create chat member")
			return
		}
	}

	DB.Create(&chat)
}
