package models

import (
	"database/sql"
	"fmt"

	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
)

type GroupMember struct {
	GroupID int `json:"group_id"`
	UserID  int `json:"user_id"`
}

func AddGroupMember(groupID, userID int) error {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO GroupMember (group_id, user_id) VALUES (?, ?)", groupID, userID)
	if err != nil {
		return fmt.Errorf("failed to add user to group: %v", err)
	}

	return nil
}

func IsGroupMember(userID, groupID int) (bool, error) {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return false, fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM GroupMember WHERE user_id = ? AND group_id = ?", userID, groupID).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to query database: %v", err)
	}

	return count > 0, nil
}

func LeaveGroup(userID, groupID int) error {
	db, err := sqlite.ConnectDB()
    if err != nil {
        return fmt.Errorf("failed to connect to database: %v", err)
    }
    defer db.Close()

    _, err = db.Exec("DELETE FROM GroupMember WHERE user_id = ? AND group_id = ?", userID, groupID)
    if err != nil {
        return fmt.Errorf("failed to remove user from group: %v", err)
    }

    return nil
}
