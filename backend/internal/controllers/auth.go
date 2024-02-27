package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
	"github.com/OumarLAM/SocialFace/internal/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if user.Email == "" || user.Password == "" || user.Firstname == "" || user.Lastname == "" || user.DateOfBirth == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	// Save user to database
	db, err := sqlite.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO User (email, password, firstname, lastname, date_of_birth, avatar_image, nickname, about_me, profile_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		user.Email, user.Password, user.Firstname, user.Lastname, user.DateOfBirth, user.AvatarImage, user.Nickname, user.AboutMe, user.ProfileType)
	if err != nil {
		http.Error(w, "Failed to save user to database", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var loginCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&loginCredentials)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Fetch user from database
	db, err := sqlite.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user models.User
	err = db.QueryRow(`SELECT * FROM User WHERE email =?`, loginCredentials.Email).Scan(
		&user.UserId, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.DateOfBirth, &user.AvatarImage, &user.Nickname, &user.AboutMe, &user.ProfileType)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Compare password
	err = ComparePassword(user.Password, loginCredentials.Password)
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Generate session token
	sessionToken := GenerateUUID()

	// Store session token in the database
	if err := sqlite.StoreSessionToken(user.UserId, sessionToken); err != nil {
		http.Error(w, "Failed to store session token", http.StatusInternalServerError)
		return
	}

	// Store session token in cookie
	cookie := http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true, // Cookie cannot be accessed by javascript
	}
	http.SetCookie(w, &cookie)

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged in successfully"})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Clear session token from cookie
	cookie := http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
	}
	http.SetCookie(w, &cookie)

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateUUID() string {
	return uuid.New().String()
}
