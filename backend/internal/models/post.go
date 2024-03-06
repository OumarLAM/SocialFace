package models

import "github.com/OumarLAM/SocialFace/internal/db/sqlite"

type Post struct {
	PostID    int    `json:"post_id"`
	UserID    int    `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at,omitempty"`
	Privacy   string `json:"privacy"`
	ImageGIF  string `json:"image_gif,omitempty"`
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