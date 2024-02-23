package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/OumarLAM/SocialFace/config"
	"github.com/OumarLAM/SocialFace/migrations"
)

func main() {
	_, err := config.InitializeDB()
	if err != nil {
		log.Println("Driver creation failed: ", err.Error())
	} else {
		// Run all migrations
		migrations.Run()

		router := gin.Default()

		router.Run(":8000")
		log.Println("Server started on port 8000")
	}
}