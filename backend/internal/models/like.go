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

type Dislike struct {
	DislikeID int    `json:"dislike_id"`
	UserID    int    `json:"user_id"`
	PostID    int    `json:"post_id"`
	CreatedAt string `json:"created_at,omitempty"`
}

func LikePost(userID, postID int) error {
	
	if disliked, err := isDisliked(userID, postID); err != nil {
		return err
	} else if disliked {
		if err := UndislikePost(userID, postID); err != nil {
			return err
		}
	}

	if liked, err := isLiked(userID, postID); err!= nil {
		return err
	} else if liked {
		return nil
	}

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

func DislikePost(userID, postID int) error {
	if liked, err := isLiked(userID, postID); err!= nil {
		return err
	} else if liked {
		if err := UnlikePost(userID, postID); err!= nil {
            return err
        }
	}

	if disliked, err := isDisliked(userID, postID); err!= nil {
        return err
    } else if disliked {
        return nil
    }

	db, err := sqlite.ConnectDB()
    if err != nil {
        return fmt.Errorf("failed to connect to database: %v", err)
    }
    defer db.Close()
	
    _, err = db.Exec("INSERT INTO Dislike (user_id, post_id) VALUES (?, ?)", userID, postID)
    if err != nil {
		return fmt.Errorf("failed to dislike post: %v", err)
    }
	
    return nil
}

func isLiked(userID, postID int) (bool, error) {
	db, err := sqlite.ConnectDB()
    if err != nil {
        return false, fmt.Errorf("failed to connect to database: %v", err)
    }
    defer db.Close()

    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM Like WHERE user_id = ? AND post_id = ?", userID, postID).Scan(&count)
	if err != nil {
        return false, fmt.Errorf("failed to query database: %v", err)
    }

	return count > 0, nil   
}

func isDisliked(userID, postID int) (bool, error) {
	db, err := sqlite.ConnectDB()
    if err != nil {
        return false, fmt.Errorf("failed to connect to database: %v", err)
    }
    defer db.Close()

    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM Dislike WHERE user_id = ? AND post_id = ?", userID, postID).Scan(&count)
    if err != nil {
        return false, fmt.Errorf("failed to query database: %v", err)
    }

    return count > 0, nil
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

func UndislikePost(userID, postID int) error {
	db, err := sqlite.ConnectDB()
    if err != nil {
        return fmt.Errorf("failed to connect to database: %v", err)
    }
    defer db.Close()

    _, err = db.Exec("DELETE FROM Dislike WHERE user_id = ? AND post_id = ?", userID, postID)
    if err != nil {
        return fmt.Errorf("failed to undislike post: %v", err)
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
