package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/solomon04/go-docker/database"
	"log"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Init database
	database.InitDB()

	// Init Redis

	// Init Queue
}
