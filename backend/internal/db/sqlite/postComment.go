package sqlite

import (
	"fmt"

	"github.com/OumarLAM/SocialFace/internal/models"
)

func CreatePost(userID int, post models.Post) error {
	db, err := ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Post (user_id, content, privacy, image_gif) VALUES (?, ?, ?, ?)",
		userID, post.Content, post.Privacy, post.ImageGIF)
	if err != nil {
		return fmt.Errorf("failed to insert post into database: %v", err)
	}

	return nil
}

func CreateComment(userID, postID int, comment models.Comment) error {
	db, err := ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Comment (user_id, post_id, content, image_gif) VALUES (?, ?, ?, ?)",
		userID, postID, comment.Content, comment.ImageGIF)
	if err != nil {
		return fmt.Errorf("failed to insert comment into database: %v", err)
	}

	return nil
}

func PostExists(postID int) (bool, error) {
	db, err := ConnectDB()
    if err != nil {
        return false, fmt.Errorf("failed to connect to database: %v", err)
    }
    defer db.Close()

    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM Post WHERE post_id =?", postID).Scan(&count)
    if err != nil {
        return false, fmt.Errorf("failed to query database: %v", err)
    }

    return count > 0, nil
}
