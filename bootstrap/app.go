package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/solomon04/go-docker/src/db"
	"log"
)

func Boot() {
	// Initialize the Environment
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Init database
	db.InitDB()

	// Init Mailer
	// https://github.com/sirupsen/logrus
	// https://github.com/getsentry/sentry-go
}
