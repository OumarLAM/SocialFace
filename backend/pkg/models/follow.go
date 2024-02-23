package models

type Follow struct {
	FollowerID int `json:"follower_id"`
	FolloweeID int `json:"followee_id"`
}
