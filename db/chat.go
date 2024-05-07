package db

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func (d *Database) NewChat(chatName string, chatType ChatType, userTable []UserTable) error {
	var name string
	if chatName == "" {
		for _, user := range userTable {
			name += " " + user.Username
		}

	}
	chat := Chat{
		Name:        name,
		Type:        chatType,
		CreatedTime: time.Now(),
	}
	chatTable := ConvertChatToChatTable(chat)
	chatID, err := Mysql.CheckRepeatedDirectChat(userTable)
	if err != nil {
		log.Printf("Error checking for repeated chat: %v", err)
		return err
	}
	log.Println(chatID)
	if chatID == 0 {

		d.db.Create(&chatTable)
		var chatMembers []ChatMember
		for _, u := range userTable {
			userID, _ := strconv.Atoi(u.ID)
			chatMember := ChatMember{
				UserTable:   u,
				ChatTable:   chatTable,
				JoinedTime:  time.Now(),
				ChatTableID: chatTable.ID,
				UserTableID: userID,
				LeftTime:    time.Date(1, time.January, 1, 1, 1, 1, 0, time.UTC),
			}
			chatMembers = append(chatMembers, chatMember)
			// log.Println("List of chatMembers", chatMembers)

		}
		if err := d.db.Create(&chatMembers).Error; err != nil {

			return errors.New("cannot create chat member")

		}

	}
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

func (d *Database) CheckRepeatedDirectChat(userTable []UserTable) (int, error) {
	var userIDs []string
	for _, user := range userTable {
		userIDs = append(userIDs, user.ID)
	}
	var chatTable ChatTable
	err := d.db.Model(&ChatTable{}).Joins("JOIN chat_members ON chat_tables.id=chat_members.chat_table_id").Where("chat_tables.type=?", 0).Where("chat_members.user_table_id IN ?", userIDs).Find(&chatTable).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("no direct chat found: %v ", err)
			return 0, nil
		}
		log.Printf("failed to query for repeated direct chat : %v ", err)
		return 0, err
	}
	return chatTable.ID, nil
}
