package migrations

import (
	"database/sql"
	"github.com/OumarLAM/SocialFace/config"
	"log"
)

func Run() {
	// Migrate our tables
	migrate(config.DB, User)
	migrate(config.DB, Follow)
	migrate(config.DB, Post)
	migrate(config.DB, Comment)
	migrate(config.DB, Group)
	migrate(config.DB, GroupMember)
	migrate(config.DB, Event)
	migrate(config.DB, EventResponse)
	migrate(config.DB, Chat)
	migrate(config.DB, Notification)
}

func migrate(dbDriver *sql.DB, query string) {
	statement, err := dbDriver.Prepare(query)
	if err == nil {
		_, creationErr := statement.Exec()
		if creationErr == nil {
			log.Println("Table created successfully")
		} else {
			log.Println(creationErr.Error())
		}
	} else {
		log.Println(err.Error())
	}
}
