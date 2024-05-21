package db

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

func (d *Database) NewChat(chatName string, chatType ChatType, userTable []UserTable) (string, error) {
	name := chatName
	if chatName == "" {
		for _, user := range userTable {
			name += "" + user.Username
		}

	}
	chatID := hashDB(chatName)

	chat := Chat{
		Name:        name,
		Type:        chatType,
		CreatedTime: time.Now(),
	}
	chatTable := ConvertChatToChatTable(chat)
	if chatType == Direct {

		directChatID, err := Mysql.CheckRepeatedDirectChat(userTable)
		if err != nil {
			return "", fmt.Errorf("error checking for repeated chat: %v", err)

		}

		if directChatID != "" {
			log.Printf("direct chat already exists: %v", directChatID)
			return directChatID, nil
		}
		chatID, err = chatIDForDirectChat(userTable)
		if err != nil {
			return "", fmt.Errorf("error generating chat id for chat: %v", err)
		}
		chatTable.ID = chatID
		d.db.Create(&chatTable)
		if err := d.generateChatMemberForChat(userTable, chatTable); err != nil {
			return "", fmt.Errorf("error generating chat member for chat : %v", err)
		}

		if err != nil {
			return "", fmt.Errorf("error generating ID for direct chat : %v", err)
		}
	}
	if chatType == Group {

		chatTable.ID = chatID
		d.db.Create(&chatTable)
		if err := d.generateChatMemberForChat(userTable, chatTable); err != nil {
			return "", fmt.Errorf("error generating chat member for chat : %v", err)
		}
	}

	return chatID, nil
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

func (d *Database) CheckRepeatedDirectChat(userTable []UserTable) (string, error) {

	hashedChatID, err := chatIDForDirectChat(userTable)
	if err != nil {
		log.Print(err)
		return "", err
	}
	var chatTable ChatTable
	err = d.db.Model(&chatTable).Where("id = ?", hashedChatID).First(&chatTable).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("no direct chat found: %v ", err)
			return "", nil
		}
		log.Printf("failed to query for repeated direct chat : %v ", err)
		return "", err
	}
	return chatTable.ID, nil

}

func chatIDForDirectChat(userTable []UserTable) (string, error) {
	if len(userTable) != 2 {
		return "", errors.New("wrong number of users")
	}
	currentID := userTable[0].ID
	secondID := userTable[1].ID
	chatID := currentID + secondID
	hashedChatID := hashDB(chatID)
	return hashedChatID, nil

}

func (d *Database) generateChatMemberForChat(userTable []UserTable, chatTable ChatTable) error {
	var chatMembers []ChatMember
	for _, u := range userTable {

		chatMember := ChatMember{
			UserTable:   u,
			ChatTable:   chatTable,
			JoinedTime:  time.Now(),
			ChatTableID: chatTable.ID,
			UserTableID: u.ID,
		}
		chatMembers = append(chatMembers, chatMember)

	}
	if err := d.db.Create(&chatMembers).Error; err != nil {

		return fmt.Errorf("cannot create chat member: %w", err)

	}
	return nil

}
