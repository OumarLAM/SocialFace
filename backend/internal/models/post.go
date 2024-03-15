package models

import (
	"fmt"

	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
)

type PrivacyType int

const (
	Public PrivacyType = iota + 1
	Private
	AlmostPrivate
)

type Post struct {
	PostID    int         `json:"post_id"`
	UserID    int         `json:"user_id"`
	Content   string      `json:"content"`
	CreatedAt string      `json:"created_at,omitempty"`
	Privacy   PrivacyType `json:"privacy"`
	ImageGIF  string      `json:"image_gif,omitempty"`
}

func CreatePost(userID int, post Post) error {
	db, err := sqlite.ConnectDB()
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

func PostExists(postID int) (bool, error) {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return false, fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Post WHERE post_id = ?", postID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to query database: %v", err)
	}

	return count > 0, nil
}

func GetPostsByUserID(userID int) ([]Post, error) {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var posts []Post
	rows, err := db.Query("SELECT post_id, content, created_at, privacy, image_gif FROM Post WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		if err = rows.Scan(&post.PostID, &post.Content, &post.CreatedAt, &post.Privacy, &post.ImageGIF); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
