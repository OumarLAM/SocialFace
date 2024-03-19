package models

import (
	"fmt"

	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
)

type Follow struct {
	FollowerID int `json:"follower_id"`
	FolloweeID int `json:"followee_id"`
}

func FollowUser(followerID, followeeID int) error {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Follow (follower_id, followee_id) VALUES (?, ?)", followerID, followeeID)
	if err != nil {
		return fmt.Errorf("failed to follow user: %v", err)
	}

	return nil
}

func UnfollowUser(followerID, followeeID int) error {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Follow WHERE follower_id = ? AND followee_id = ?", followerID, followeeID)
	if err != nil {
		return fmt.Errorf("failed to unfollow user: %v", err)
	}

	return nil
}

func SendFollowRequest(followerID, followeeID int) error {
	// Check if the follow request already exists
	exists, err := FollowRequestExists(followerID, followeeID)
	if err != nil {
		return fmt.Errorf("failed to check if the follow request already exists: %v", err)
	}
	if exists {
		return fmt.Errorf("follow request already send")
	}

	// Create the follow request in the database
	err = CreateFollowRequest(followerID, followeeID)
	if err != nil {
		return fmt.Errorf("failed to create follow request: %v", err)
	}

	return nil
}

func FollowRequestExists(followerID, followeeID int) (bool, error) {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return false, fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM FollowRequest WHERE follower_id = ? AND followee_id = ?", followerID, followeeID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to query database: %v", err)
	}

	return count > 0, nil
}

func CreateFollowRequest(followerID, followeeID int) error {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO FollowRequest (follower_id, followee_id) VALUES (?, ?)", followerID, followeeID)
	if err != nil {
		return fmt.Errorf("failed to create follow request: %v", err)
	}

	return nil
}

func AcceptFollowRequest(userID, followerID int) error {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Start a transaction to ensure atomicity
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	_, err = tx.Exec("INSERT INTO Follow (follower_id, followee_id) VALUES (?, ?)", followerID, userID)
	if err != nil {
		return fmt.Errorf("failed to accept follow request: %v", err)
	}

	_, err = tx.Exec("DELETE FROM FollowRequest WHERE follower_id = ? AND followee_id = ?", followerID, userID)
	if err != nil {
		return fmt.Errorf("failed to delete follow request: %v", err)
	}

	return nil
}

func DeclineFollowRequest(userID, followerID int) error {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Delete the corresponding record from the followeRequest table
	_, err = db.Exec("DELETE FROM FollowRequest WHERE follower_id = ? AND followee_id = ?", followerID, userID)
    if err != nil {
        return fmt.Errorf("failed to delete follow request: %v", err)
    }

    return nil
}

func IsFollowing(followerID, followeeID int) (bool, error) {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return false, fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Follow WHERE follower_id = ? AND followee_id = ?", followerID, followeeID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to query database: %v", err)
	}

	return count > 0, nil
}
