package models

import "github.com/OumarLAM/SocialFace/internal/db/sqlite"

type Comment struct {
	CommentID int    `json:"comment_id"`
	UserID    int    `json:"user_id"`
	PostID    int    `json:"post_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at,omitempty"`
	ImageGIF  string `json:"image_gif,omitempty"`
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
