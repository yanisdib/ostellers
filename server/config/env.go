package config

import (
	"log"
	"os"

	"github.com/subosito/gotenv"

	"yanisdib/ostellers/errors"
)

func GetMongoURI() string {

	err := gotenv.Load()
	if err != nil {
		log.Fatal(errors.ERR_ENV_FILE_LOADING)
	}

	// dbUser := os.Getenv("ATLAS_DB_USER")
	// dbPassword := os.Getenv("ATLAS_DB_PASSWORD")
	// dbHost := os.Getenv("ATLAS_DB_HOST")

	uri := "mongodb+srv://yanisdib:5fMFdvBsZZ26qsX5@eva-01.bg9f1tx.mongodb.net/"

	// fmt.Sprintf("mongodb+srv://%s:%s@%s", dbUser, dbPassword, dbHost)

	return uri

}

func GetDatabaseName() string {

	err := gotenv.Load()
	if err != nil {
		log.Fatal(errors.ERR_ENV_FILE_LOADING)
	}

	return os.Getenv("DB_NAME")

}
