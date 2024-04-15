package db

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func NewChat(chatName string, chatType ChatType, user []User) error {
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

		if err := DB.Create(&chatMembers).Error; err != nil {
			DB.Delete(&chatMembers)
			return errors.New("cannot create chat member")

		}
	}

	DB.Create(&chat)
	return nil
}

func GetChatMessages(ChatID int64) ([]MessageInformation, error) {
	var messages []MessageInformation
	if err := DB.Where("chat_id = ?", ChatID).Find(&messages).Error; err != nil {
		return nil, fmt.Errorf("no  message found for chat %w", err)
	}
	return messages, nil
}
