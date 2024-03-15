package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/OumarLAM/SocialFace/internal/models"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var Comment models.Comment
	var request struct {
		PostID  int    `json:"post_id"`
		Content string `json:"content"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if request.Content == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	postExists, err := models.PostExists(request.PostID)
	if err != nil {
		http.Error(w, "Failed to check if post exists", http.StatusInternalServerError)
		return
	}
	if !postExists {
		http.Error(w, "Post does not exists", http.StatusNotFound)
		return
	}

	Comment.Content = request.Content
	err = models.CreateComment(userID, request.PostID, Comment)
	if err != nil {
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Comment created successfully"})
}
