package models

import (
	"encoding/json"
	"fmt"

	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
)

const (
	OptionGoing    = "Going"
	OptionNotGoing = "Not going"
)

type Event struct {
	EventID     int      `json:"event_id"`
	GroupID     int      `json:"group_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	DateTime    string   `json:"date_time,omitempty"`
	Options     []string `json:"options"`
}

func CreateEvent(event *Event) error {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	event.Options = []string{OptionGoing, OptionNotGoing}

	// Convert options slice to JSON
	optionsJSON, err := json.Marshal(event.Options)
	if err != nil {
		return fmt.Errorf("failed to marshal options: %v", err)
	}

	_, err = db.Exec("INSERT INTO Event (group_id, title, description, date_time, options) VALUES (?, ?, ?, ?, ?)", event.GroupID, event.Title, event.Description, event.DateTime, string(optionsJSON))
	if err != nil {
		return fmt.Errorf("failed to create event: %v", err)
	}

	return nil
}
