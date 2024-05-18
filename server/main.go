package main

import (
	"sync"

	"github.com/gin-gonic/gin"

	"yanisdib/ostellers/artbook"
	"yanisdib/ostellers/config"
)

func main() {

	var once sync.Once

	// running this function only once
	once.Do(func() {

		router := gin.Default()
		// establishing connection to MongoDB
		config.OpenDBConnection()
		// creating artbook routes
		artbook.NewArtbookRoutes(router)
		// running server
		router.Run("localhost:6000")

	})

}
