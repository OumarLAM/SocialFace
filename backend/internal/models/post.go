package models

type Post struct {
	PostID    int    `json:"post_id"`
	UserID    int    `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at,omitempty"`
	Privacy   string `json:"privacy"`
	ImageGIF  string `json:"image_gif,omitempty"`
}
