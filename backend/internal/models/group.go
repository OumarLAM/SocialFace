package models

import (
	"database/sql"
	"fmt"

	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
)

type UserGroup struct {
	GroupID     int    `json:"group_id"`
	CreatorID   int    `json:"creator_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateGroup(group UserGroup) error {
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

	result, err := tx.Exec("INSERT INTO UserGroup (creator_id, title, description) VALUES (?, ?, ?)", group.CreatorID, group.Title, group.Description)
	if err != nil {
		return fmt.Errorf("failed to insert group into database: %v", err)
	}

	// Retrieve the auto-generated group ID
	groupID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last inserted group ID: %v", err)
	}

	_, err = tx.Exec("INSERT INTO GroupMember (group_id, user_id) VALUES (?, ?)", groupID, group.CreatorID)
	if err != nil {
		return fmt.Errorf("failed to add creator as member of the group: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func BrowseGroups() ([]UserGroup, error) {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT group_id, creator_id, title, description FROM UserGroup")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve groups: %v", err)
	}
	defer rows.Close()

	var groups []UserGroup
	for rows.Next() {
		var group UserGroup
		if err := rows.Scan(&group.GroupID, &group.CreatorID, &group.Title, &group.Description); err != nil {
			return nil, fmt.Errorf("failed to scan group row: %v", err)
		}
		groups = append(groups, group)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed iterating over group rows: %v", err)
	}

	return groups, nil
}

func IsGroupCreator(userID, groupID int) (bool, error) {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return false, fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	var creatorID int
	err = db.QueryRow("SELECT creator_id FROM UserGroup WHERE group_id = ?", groupID).Scan(&creatorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("group not found: %v", err)
		}
		return false, fmt.Errorf("failed to get group creator ID: %v", err)
	}

	return userID == creatorID, nil
}
