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
    if err != nil {
        return fmt.Errorf("failed to connect to database: %v", err)
    }
    defer db.Close()

    // Prepare sql statement
	stmt, err := db.Prepare("UPDATE user SET session_token = ? WHERE user_id = ?")
    if err != nil {
        return fmt.Errorf("failed to prepare statement: %v", err)
    }
    defer stmt.Close()

    // Execute sql statement
    _, err = stmt.Exec(sessionToken, userId)
    if err != nil {
        return fmt.Errorf("failed to execute statement: %v", err)
    }

    return nil
}

func IsSessionTokenValid(sessionToken string) (int, bool) {
	db, err := ConnectDB()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return 0, false
	}
	defer db.Close()

	var userId int
	var count int
	err = db.QueryRow("SELECT user_id, COUNT(*) FROM user WHERE session_token = ? AND session_expiration > datetime('now')", sessionToken).Scan(&userId, &count)
	if err != nil {
		fmt.Printf("Failed to query database: %v\n", err)
		return 0, false
	}

	return userId, count > 0
}

func ClearSessionToken(sessionToken string) error {
	db, err := ConnectDB()
    if err != nil {
        return fmt.Errorf("failed to connect to database: %v", err)
    }
    defer db.Close()

    _, err = db.Exec(`UPDATE User SET session_token = NULL, session_expiration = NULL WHERE session_token = ?`, sessionToken)
    if err != nil {
		return fmt.Errorf("failed to delete user session: %v", err)
	}

    return nil
}
