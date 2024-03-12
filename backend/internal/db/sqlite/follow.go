package sqlite

import (
	"log"
)

func IsFollowing(userID, targetUserID int) bool {
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	var following bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM Follow WHERE follower_id =? AND followed_id =?) AS following", userID, targetUserID).Scan(&following)
	if err != nil {
		log.Fatalf("failed to query database: %v", err)
	}

	return following
}
