package main

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"

	"yanisdib/ostellers/artbook"
	"yanisdib/ostellers/config"
)

func main() {

	var once sync.Once

	// init and run server only once
	once.Do(func() {

		router := gin.Default()

		config.OpenDBConnection()
		artbook.NewArtbookRoutes(router)

		if err := router.Run("localhost:6000"); err != nil {
			log.Fatal("Oops! An error occured while runnning the server at localhost:6000")
		}

	})

}
