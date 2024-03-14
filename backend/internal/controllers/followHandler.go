package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
)

func FollowUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	// Parse request body to get the user ID of the user to follow
	var followRequest struct {
		FolloweeID int `json:"followee_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&followRequest)
	if err != nil {
		http.Error(w, "Failed to decode request body to get the user ID of the user to follow", http.StatusBadRequest)
		return
	}

	// Check if the user is trying to follow themselves
	if followRequest.FolloweeID == userID {
		http.Error(w, "Cannot follow yourself", http.StatusBadRequest)
		return
	}

	// Check if the user is already following the target user
	isFollowing, err := sqlite.IsFollowing(userID, followRequest.FolloweeID)
	if err != nil {
		http.Error(w, "Failed to check if the user is already following the target user", http.StatusInternalServerError)
		return
	}
	if isFollowing {
		http.Error(w, "You are already following this user", http.StatusBadRequest)
		return
	}

	// Follow the user
	err = sqlite.FollowUser(userID, followRequest.FolloweeID)
	if err != nil {
		http.Error(w, "Failed to follow user", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User followed successfully"})
}

func UnfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
	}
	
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	// Parse request body to get the user ID to unfollow
	var unfollowRequest struct {
		FolloweeID int `json:"followee_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&unfollowRequest)
	if err != nil {
		http.Error(w, "Failed to decode request body to get the user ID to unfollow", http.StatusBadRequest)
		return
	}

	// Check if the user is trying to unfollow themselves
	if unfollowRequest.FolloweeID == userID {
		http.Error(w, "Cannot unfollow yourself", http.StatusBadRequest)
		return
	}

	// Check if the user is not following the target user
	isfollowing, err := sqlite.IsFollowing(userID, unfollowRequest.FolloweeID)
	if err != nil {
		http.Error(w, "Failed to check if the user is following", http.StatusInternalServerError)
		return
	}

	if !isfollowing {
		http.Error(w, "You are not following this user", http.StatusBadRequest)
		return
	}

	// Unfollow the user
	err = sqlite.UnfollowUser(userID, unfollowRequest.FolloweeID)
	if err != nil {
		http.Error(w, "Failed to unfollow user", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User unfollowed successfully"})
}

func AcceptFollowRequestHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	// Parse request body to get the user ID of the follower to accept
	var acceptRequest struct {
		FollowerID int `json:"follower_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&acceptRequest)
	if err != nil {
		http.Error(w, "Failed to decode request body to get the user ID of the follower to accept", http.StatusBadRequest)
		return
	}

	// Check if the follower's user ID is valid
	if acceptRequest.FollowerID == 0 {
		http.Error(w, "Invalid follower ID", http.StatusBadRequest)
		return
	}

	// Check if the requester is the intended recipient of the follow request
	if acceptRequest.FollowerID == userID {
		http.Error(w, "You cannot accept your own follow request", http.StatusBadRequest)
		return
	}

	// Accept the follow request
	err = sqlite.AcceptFollowRequest(userID, acceptRequest.FollowerID)
	if err != nil {
		http.Error(w, "Failed to accept follow request", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Follow request accepted successfully"})
}

func DeclineFollowRequestHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	// Parse request body to get the user ID of the follower to decline
	var declineRequest struct {
		FollowerID int `json:"follower_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&declineRequest)
	if err != nil {
		http.Error(w, "Failed to decode request body of the follower to decline", http.StatusBadRequest)
		return
	}

	// Decline the follow request
	err = sqlite.DeclineFollowRequest(userID, declineRequest.FollowerID)
	if err != nil {
		http.Error(w, "Failed to decline follow request", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Follow request declined successfully"})
}