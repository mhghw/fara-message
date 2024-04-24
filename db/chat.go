package db

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

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
type ChatType struct {
	chatType int
}

func (c *ChatType) Int() int {
	return c.chatType
}

var (
	Direct  = ChatType{chatType: 0}
	Group   = ChatType{chatType: 1}
	Unknown = ChatType{chatType: -1}
)

func (d *Database) NewChat(chatName string, chatType ChatType, user []User) error {
	chat := Chat{
		Name:        chatName,
		Type:        chatType,
		CreatedTime: time.Now(),
	}
	var chatMembers []ChatMember
	for i, u := range user {
		userID, _ := strconv.Atoi(u.ID)
		chatMembers[i] = ChatMember{
			JoinedTime: time.Now(),
			ChatID:     chat.ID,
			UserID:     int64(userID),
		}

		if err := Mysql.db.Create(&chatMembers).Error; err != nil {
			Mysql.db.Delete(&chatMembers)
			return errors.New("cannot create chat member")

		}
	}

	Mysql.db.Create(&chat)
	return nil
}

func (d *Database) GetChatMessages(ChatID int64) ([]Message, error) {
	var messages []Message
	if err := Mysql.db.Where("chat_id = ?", ChatID).Find(&messages).Error; err != nil {
		return nil, fmt.Errorf("no  message found for chat %w", err)
	}
	return messages, nil
}

func (d *Database) GetUsersChatMembers(userID int) ([]ChatMember, error) {
	var usersChats []ChatMember
	if err := Mysql.db.Where("user_id = ?", userID).Find(&usersChats).Error; err != nil {
		return nil, fmt.Errorf("no  chat found for user %w", err)
	}
	return usersChats, nil
}
