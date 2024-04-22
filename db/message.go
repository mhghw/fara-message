package db

func DeleteMessage(messageID int) error {
	var message Message
	message.ID = messageID
	result := DB.Where("ID=?", message.ID).Delete(&Message{})
	if result.Error != nil { /////check!!
		return result.Error
	}
	return nil
}

func SendMessage(messageID int, senderID int, chatID int, content string) error {
	var message Message
	message.ID = messageID
	message.SenderID = senderID
	message.ChatID = chatID
	message.Content = content
	result := DB.Create(&message)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
