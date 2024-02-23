package models

type Comment struct {
	CommentID int    `json:"comment_id"`
	UserID    int    `json:"user_id"`
	PostID    int    `json:"post_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at,omitempty"`
	ImageGIF  string `json:"image_gif,omitempty"`
}
