package db

import "fmt"

func (d *Database) SendMessage(senderID int, chatID int, content string) error {
	var message Message
	message.UserTableID = senderID
	message.ChatTableID = chatID
	message.Content = content
	result := d.db.Create(&message)
	if result.Error != nil {
		return fmt.Errorf("error sending message: %w", result.Error)
	}
	return nil
}

func (d *Database) DeleteMessage(messageID int) error {
	var message Message
	result := d.db.Where("ID=?", messageID).Delete(&message)
	if result.Error != nil {
		return fmt.Errorf("error deleting message: %w", result.Error)
	}
	return nil
}
