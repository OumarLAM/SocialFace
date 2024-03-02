package models

import (
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

// The omitempty tag in Go's JSON serialization indicates that
// if a field has its zero value for its type (e.g., an empty string ""
//  for a string type), the field will be omitted from the JSON output.
