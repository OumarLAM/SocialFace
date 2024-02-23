package models

type Chat struct {
	ChatID      int    `json:"chat_id"`
	SenderID    int    `json:"sender_id"`
	RecipientID int    `json:"recipient_id"`
	Message     string `json:"message"`
	CreatedAt   string `json:"created_at,omitempty"`
}
