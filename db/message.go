package db

func (d *Database) DeleteMessage(messageID int) error {
	var message Message
	message.ID = messageID
	result := d.db.Where("ID=?", message.ID).Delete(&Message{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *Database) SendMessage(messageID int, senderID int, chatID int, content string) error {
	var message Message
	message.ID = messageID
	message.SenderID = senderID
	message.ChatID = chatID
	message.Content = content
	result := d.db.Create(&message)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
