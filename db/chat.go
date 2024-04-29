package db

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Chat struct {
	ID          int    `gorm:"primary_key"`
	Name        string `gorm:"chat_name;default:' '"`
	CreatedTime time.Time
	DeletedTime time.Time
	Type        ChatType
}

type ChatMember struct {
	UserID     int `gorm:"foreign_key"`
	User       User
	ChatID     int `gorm:"foreign_key"`
	Chat       Chat
	JoinedTime time.Time
	LeftTime   time.Time
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
			UserID:     int(userID),
		}

		if err := d.db.Create(&chatMembers).Error; err != nil {
			d.db.Delete(&chatMembers)
			return errors.New("cannot create chat member")

		}
	}

	d.db.Create(&chat)
	return nil
}

func (d *Database) GetChatMessages(ChatID int64) ([]Message, error) {
	var messages []Message
	if err := d.db.Where("chat_id = ?", ChatID).Find(&messages).Error; err != nil {
		return nil, fmt.Errorf("no  message found for chat %w", err)
	}
	return messages, nil
}

func (d *Database) GetUsersChatMembers(userID int) ([]ChatMember, error) {
	var usersChats []ChatMember
	if err := d.db.Where("user_id = ?", userID).Find(&usersChats).Error; err != nil {
		return nil, fmt.Errorf("no  chat found for user %w", err)
	}
	return usersChats, nil
}

func (d *Database) GetUsersChatIDAndChatName(chatMember []ChatMember) ([]ChatIDAndChatName, error) {
	var result []ChatIDAndChatName
	for _, chat := range chatMember {
		result = append(result, ChatIDAndChatName{
			ChatID:   chat.ChatID,
			ChatName: chat.Chat.Name,
		})
	}
	if len(result) == 0 {
		return result, errors.New("no chat found for user ")
	}
	return result, nil
}
