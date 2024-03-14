package db

import (
	"errors"
	"fmt"
	"time"
)

func NewChat(chatName string, chatType ChatType, user []User) error {
	chat := Chat{
		ChatName:    chatName,
		Type:        chatType,
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

func GetChatMessages(ChatID int64) ([]Message, error) {
	var messages []Message
	if err := DB.Where("chat_id = ?", ChatID).Find(&messages).Error; err != nil {
		return nil, fmt.Errorf("no  message found for chat %w", err)
	}
	return messages, nil
}

func GetUserChats(UserID int64) ([]Chat, error) {
	var chats []Chat

}
