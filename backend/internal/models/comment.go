package models

import (
	"fmt"

	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
)

type Comment struct {
	CommentID int    `json:"comment_id"`
	UserID    int    `json:"user_id"`
	PostID    int    `json:"post_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at,omitempty"`
	ImageGIF  string `json:"image_gif,omitempty"`
}

func CreateComment(userID, postID int, comment Comment) error {
	db, err := sqlite.ConnectDB()
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

func GetCommentsByUserID(userID int) ([]Comment, error) {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var comments []Comment
	rows, err := db.Query("SELECT comment_id, post_id, content, created_at, image_gif FROM Comment WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		if err = rows.Scan(&comment.CommentID, &comment.PostID, &comment.Content, &comment.CreatedAt, &comment.ImageGIF); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func GetCommentsForPost(postID int) ([]Comment, error) {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT comment_id, post_id, user_id, content, created_at, image_gif FROM Comment WHERE post_id = ?", postID)
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %v", err)
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
        if err := rows.Scan(&comment.CommentID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.ImageGIF); err != nil {
            return nil, fmt.Errorf("failed to scan comment row: %v", err)
        }
        comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("failed iterating over comment rows: %v", err)
    }

	return comments, nil
}
