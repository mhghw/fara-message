package db

import "fmt"

func (d *Database) DeleteMessage(messageID int) error {
	var message Message
	message.ID = messageID
	result := d.db.Where("ID=?", message.ID).Delete(&message)
	if result.Error != nil {
		return fmt.Errorf("error deleting message: %w", result.Error)
	}
	return nil
}

func (d *Database) SendMessage(messageID int, senderID int, chatID int, content string) error {
	var message Message
	message.ID = messageID
	message.UserID = senderID
	message.ChatID = chatID
	message.Content = content
	result := d.db.Create(&message)
	if result.Error != nil {
		return fmt.Errorf("error sending message: %w", result.Error)
	}
	return nil
}
