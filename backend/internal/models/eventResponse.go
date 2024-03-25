package models

import (
	"fmt"

	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
)

type EventResponse struct {
	EventResponseID int    `json:"event_response_id"`
	EventID         int    `json:"event_id"`
	UserID          int    `json:"user_id"`
	Response        string `json:"response"`
}

func RespondToEvent(eventResponse *EventResponse) error {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO EventResponse (event_id, user_id, response) VALUES (?, ?, ?)", eventResponse.EventID, eventResponse.UserID, eventResponse.Response)
	if err != nil {
		return fmt.Errorf("failed to create event response: %v", err)
	}

	return nil
}

func GetEventResponses(eventID int) ([]EventResponse, error) {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT event_response_id, event_id, user_id, response FROM EventResponse WHERE event_id = ?", eventID)
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %v", err)
	}
	defer rows.Close()

	var eventResponses []EventResponse
	for rows.Next() {
		var eventResponse EventResponse
		if err := rows.Scan(&eventResponse.EventResponseID, &eventResponse.EventID, &eventResponse.UserID, &eventResponse.Response); err != nil {
			return nil, fmt.Errorf("failed to scan event response row: %v", err)
		}
		eventResponses = append(eventResponses, eventResponse)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed iterating over event response rows: %v", err)
	}

	return eventResponses, nil
}
