package models

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

// The omitempty tag in Go's JSON serialization indicates that
// if a field has its zero value for its type (e.g., an empty string ""
//  for a string type), the field will be omitted from the JSON output.
