package main

import (
	"log"
	"net/http"

	// "github.com/OumarLAM/SocialFace/internal/db/sqlite"
	"github.com/OumarLAM/SocialFace/internal/controllers"
	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
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


	// Start server
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", router)
	// if err != nil {
	// 	log.Fatalf("failed to start server: %v", err)
	// }
}
