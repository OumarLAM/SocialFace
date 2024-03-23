package models

type Event struct {
	EventID     int      `json:"event_id"`
	GroupID     int      `json:"group_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	DateTime    string   `json:"date_time,omitempty"`
	Options     []string `json:"options"`
}
