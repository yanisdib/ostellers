package config

import (
	"fmt"
	"log"
	"os"

	"github.com/subosito/gotenv"
)

func GetMongoURI() string {

	err := gotenv.Load()
	if err != nil {
		log.Fatal("Error occured while loading .env file")
	}

	dbUser := os.Getenv("ATLAS_DB_USER")
	dbPassword := os.Getenv("ATLAS_DB_PASSWORD")
	dbHost := os.Getenv("ATLAS_HOST")

	return fmt.Sprintf("mongodb+srv://%s:%s@%s", dbUser, dbPassword, dbHost)

}

func GetDatabaseName() string {

	err := gotenv.Load()
	if err != nil {
		log.Fatal("Error occured while loading .env file")
	}

	return os.Getenv("DB_NAME")

}
