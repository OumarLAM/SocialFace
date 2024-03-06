package sqlite

import "fmt"

func LikePost(userID, postID int) error {
	db, err := ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Like (user_id, post_id) VALUES (?,?)", userID, postID)
	if err != nil {
		return fmt.Errorf("failed to like post: %v", err)
	}

	return nil
}

func UnlikePost(userID, postID int) error {
	db, err := ConnectDB()
    if err != nil {
        return fmt.Errorf("failed to connect to database: %v", err)
    }
    defer db.Close()

    _, err = db.Exec("DELETE FROM Like WHERE user_id =? AND post_id =?", userID, postID)
    if err != nil {
        return fmt.Errorf("failed to unlike post: %v", err)
    }

    return nil
}