package main

import (
	"log"
	"net/http"

	"github.com/OumarLAM/SocialFace/internal/controllers"
	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
	"github.com/OumarLAM/SocialFace/internal/middlewares"
)

func main() {
	// Connect to the database
	db, err := sqlite.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Apply migrations
	if err := sqlite.MigrateDB(db); err != nil {
		log.Fatalf("failed to apply migrations: %v", err)
	}

	// Initialize router
	router := http.NewServeMux()

	// Register authentication endpoints
	router.HandleFunc("/register", controllers.RegisterHandler)
	router.HandleFunc("/login", controllers.LoginHandler)
	router.HandleFunc("/logout", controllers.LogoutHandler)

	// Register profile endpoints
	router.HandleFunc("/profile/info", middlewares.AuthMiddleware(controllers.ProfileInfoHandler))
	router.HandleFunc("/profile/privacy", middlewares.AuthMiddleware(controllers.UpdateProfilePrivacyHandler))

	// Start server
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}
