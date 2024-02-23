package models

type User struct {
	UserId      int    `json:"user_id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	DateOfBirth string `json:"date_of_birth"`
	AvatarImage string `json:"avatar_image,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	AboutMe     string `json:"about_me,omitempty"`
	ProfileType string `json:"profile_type"`
}


// The omitempty tag in Go's JSON serialization indicates that
// if a field has its zero value for its type (e.g., an empty string ""
//  for a string type), the field will be omitted from the JSON output.