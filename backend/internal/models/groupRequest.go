package models

import (
	"fmt"

	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
)

type GroupRequest struct {
	RequestID int    `json:"request_id"`
	UserID    int    `json:"user_id"`
	GroupID   int    `json:"group_id"`
	Status    string `json:"status"`
}

func AddGroupRequest(userID, groupID int) error {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO GroupRequest (user_id, group_id, status) VALUES (?, ?, ?)", userID, groupID, "pending")
	if err != nil {
		return fmt.Errorf("failed to add group request: %v", err)
	}

	return nil
}

func AcceptGroupRequest(userID, groupID, requestID int) error {
	isCreator, err := IsGroupCreator(userID, groupID)
	if err != nil {
		return fmt.Errorf("failed to check group creator: %v", err)
	}
	if !isCreator {
		return fmt.Errorf("User is not the creator of the group")
	}

	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.Exec("UPDATE groupRequest SET status = 'accepted' WHERE request_id = ?", requestID)
	if err != nil {
		return fmt.Errorf("failed to accept group request: %v", err)
	}

	_, err = tx.Exec("INSERT INTO GroupMember (user_id, group_id) VALUES (?, ?)", userID, groupID)
	if err != nil {
		return fmt.Errorf("failed to add user to group: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func DeclineGroupRequest(userID, groupID, requestID int) error {
	isCreator, err := IsGroupCreator(userID, groupID)
	if err != nil {
		return fmt.Errorf("failed to check group creator: %v", err)
	}
	if !isCreator {
		return fmt.Errorf("User is not the creator of the group")
	}

	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE GroupRequest SET status = 'declined' WHERE request_id = ?", requestID)
	if err != nil {
		return fmt.Errorf("failed to decline group request: %v", err)
	}

	return nil
}
