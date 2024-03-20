package models

import (
	"database/sql"
	"fmt"

	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
)

type User struct {
	UserId            int    `json:"user_id"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Firstname         string `json:"firstname"`
	Lastname          string `json:"lastname"`
	DateOfBirth       string `json:"date_of_birth"`
	AvatarImage       string `json:"avatar_image,omitempty"`
	Nickname          string `json:"nickname,omitempty"`
	AboutMe           string `json:"about_me,omitempty"`
	PublicProfile     bool   `json:"public_profile"`
	SessionToken      string `json:"session_token,omitempty"`
	SessionExpiration string `json:"session_expiration,omitempty"`
}

func (u *User) IsProfilePublic() bool {
	return u.PublicProfile
}

func (u *User) SetProfilePrivacy(publicProfile bool) {
	u.PublicProfile = publicProfile
}

// GetUserByID retrieves user information from the database based on the user ID.
func GetUserByID(userID int) (*User, error) {
	// Connect to the database
	db, err := sqlite.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query the database for user information
	var user User
	err = db.QueryRow(`SELECT user_id, email, firstname, lastname, date_of_birth, avatar_image, nickname, about_me, public_profile FROM User WHERE user_id = ?`, userID).Scan(&user.UserId, &user.Email, &user.Firstname, &user.Lastname, &user.DateOfBirth, &user.AvatarImage, &user.Nickname, &user.AboutMe, &user.PublicProfile)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateProfilePrivacy updates the profile privacy settings for a user in the database.
func UpdateProfilePrivacy(userID int, publicProfile bool) error {
	// Connect to the database
	db, err := sqlite.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Update profile privacy settings in the database
	_, err = db.Exec(`UPDATE User SET public_profile = ? WHERE user_id = ?`, publicProfile, userID)
	if err != nil {
		return err
	}

	return nil
}

func FetchFollowers(userID int) ([]User, error) {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT u.user_id, u.email, u.firstname, u.lastname FROM User u JOIN Follow f ON u.user_id = f.follower_id WHERE f.followee_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []User
	for rows.Next() {
		var follower User
		err := rows.Scan(&follower.UserId, &follower.Email, &follower.Firstname, &follower.Lastname)
		if err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return followers, nil
}

func FetchFollowing(userID int) ([]User, error) {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT u.user_id, u.email, u.firstname, u.lastname FROM User u JOIN Follow f ON u.user_id = f.followee_id WHERE f.follower_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var following []User
	for rows.Next() {
		var followingUser User
		err := rows.Scan(&followingUser.UserId, &followingUser.Email, &followingUser.Firstname, &followingUser.Lastname)
		if err != nil {
			return nil, err
		}
		following = append(following, followingUser)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return following, nil
}

func UserExists(userID int) bool {
	db, err := sqlite.ConnectDB()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v", err)
		return false
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM User WHERE user_id = ?)", userID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("Failed to check if user exists: %v", err)
		return false
	}

	return exists
}