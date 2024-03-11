package send_message

type MessageInformation struct{
	SenderID int `json:"senderID"`
	ChatID int  `json:"chatID"`
	Content string `json:"content"`
}
