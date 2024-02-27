package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	// Get the absolute path to the migrations directory
	absPath, err := filepath.Abs("../../internal/db/migrations/sqlite")
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %v", err)
	}

	// Check if the migrations directory exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("migrations directory does not exist: %v", err)
	}

	// Open the sqlite database
	db, err := sql.Open("sqlite3", "../../internal/db/socialface.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Ping the database to verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}

func MigrateDB(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %v", err)
	}

	// Get the absolute path to the migrations directory
	absPath, err := filepath.Abs("../../internal/db/migrations/sqlite")
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	// Run migrations
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+absPath,
		"sqlite3", driver)

	if err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("an error occurred while syncing the database: %v", err)
	}

	log.Println("Migration completed successfully")
	return nil
}

// Function to store session token in the database
func StoreSessionToken(userId int, sessionToken string) error {
	db, err := ConnectDB()
    if err!= nil {
        return fmt.Errorf("failed to connect to database: %v", err)
    }
    defer db.Close()

    // Prepare sql statement
	stmt, err := db.Prepare("UPDATE user SET session_token = ? WHERE user_id = ?")
    if err!= nil {
        return fmt.Errorf("failed to prepare statement: %v", err)
    }
    defer stmt.Close()

    // Execute sql statement
    _, err = stmt.Exec(sessionToken, userId)
    if err!= nil {
        return fmt.Errorf("failed to execute statement: %v", err)
    }

    return nil
}
