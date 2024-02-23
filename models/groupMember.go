package models

type GroupMember struct {
	GroupID int `json:"group_id"`
	UserID  int `json:"user_id"`
}
