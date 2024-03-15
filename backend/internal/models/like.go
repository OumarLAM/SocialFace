package models

import (
	"fmt"

	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
)

type Like struct {
	LikeID    int    `json:"like_id"`
	UserID    int    `json:"user_id"`
	PostID    int    `json:"post_id"`
	CreatedAt string `json:"created_at,omitempty"`
}

func LikePost(userID, postID int) error {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Like (user_id, post_id) VALUES (?, ?)", userID, postID)
	if err != nil {
		return fmt.Errorf("failed to like post: %v", err)
	}

	return nil
}

func UnlikePost(userID, postID int) error {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Like WHERE user_id = ? AND post_id = ?", userID, postID)
	if err != nil {
		return fmt.Errorf("failed to unlike post: %v", err)
	}

	return nil
}

func GetLikesByUserID(userID int) ([]Like, error) {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT like_id, post_id, created_at FROM Like WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []Like
	for rows.Next() {
		var like Like
		if err = rows.Scan(&like.LikeID, &like.PostID, &like.CreatedAt); err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return likes, nil
}
