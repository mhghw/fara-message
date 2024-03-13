package delete_message

type Message struct {
	ID       int    `json:"id`
	SenderID int    `json:"senderID"`
	ChatID   int    `json:"chatID"`
	Content  string `json:"content`
}
