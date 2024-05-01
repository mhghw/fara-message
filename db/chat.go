package db

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func (d *Database) NewChat(chatName string, chatType ChatType, userTable []UserTable) error {
	chat := Chat{
		Name:        chatName,
		Type:        chatType,
		CreatedTime: time.Now(),
	}
	var chatMembers []ChatMember
	for i, u := range userTable {
		userID, _ := strconv.Atoi(u.ID)
		chatMembers[i] = ChatMember{
			JoinedTime:  time.Now(),
			ChatTableID: chat.ID,
			UserTableID: int(userID),
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
	var usersChatMembers []ChatMember
	if err := d.db.Where("user_id = ?", userID).Find(&usersChatMembers).Error; err != nil {
		return nil, fmt.Errorf("no  chat found for user %w", err)
	}
	return usersChatMembers, nil
}

func (d *Database) GetUsersChatIDAndChatName(chatMember []ChatMember) ([]ChatIDAndChatName, error) {
	var result []ChatIDAndChatName
	for _, chat := range chatMember {
		result = append(result, ChatIDAndChatName{
			ChatID:   chat.ChatTableID,
			ChatName: chat.ChatTable.Name,
		})
	}
	if len(result) == 0 {
		return result, errors.New("no chat found for user ")
	}
	return result, nil
}
