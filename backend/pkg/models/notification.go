package models

type Notification struct {
	NotificationID int    `json:"notification_id"`
	UserID         int    `json:"user_id"`
	Content        string `json:"content"`
	CreatedAt      string `json:"created_at,omitempty"`
}
