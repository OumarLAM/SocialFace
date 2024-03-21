package models

import (
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
    if err!= nil {
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
