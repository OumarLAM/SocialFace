package sqlite

import (
	"fmt"
)

func FollowUser(followerID, followeeID int) error {
	db, err := ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Follow (follower_id, followee_id) VALUES (?,?)", followerID, followeeID)
	if err != nil {
		return fmt.Errorf("failed to follow user: %v", err)
	}

	return nil
}

func UnfollowUser(followerID, followeeID int) error {
	db, err := ConnectDB()
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

func AcceptFollowRequest(userID, followerID int) error {
	db, err := ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Follow (follower_id, followee_id) VALUES (?, ?)", followerID, userID)
	if err != nil {
		return fmt.Errorf("failed to accept follow request: %v", err)
	}

	return nil
}

func DeclineFollowRequest(userID, followerID int) error {
	following, err := IsFollowing(userID, followerID)
	if err != nil {
		return fmt.Errorf("failed to check if the user is following: %v", err)
	}

	if !following {
		return fmt.Errorf("user is not following the follower")
	}

	// Unfollow the user (remov the follow relationship)
	err = UnfollowUser(userID, followerID)
	if err != nil {
		return fmt.Errorf("failed to unfollow user: %v", err)
	}

	return nil
}

func IsFollowing(followerID, followeeID int) (bool, error) {
	db, err := ConnectDB()
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