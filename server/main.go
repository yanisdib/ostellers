package main

import (
	"log"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"yanisdib/ostellers/artbook"
	"yanisdib/ostellers/config"
)

func main() {

	var once sync.Once

	// init and run server only once
	once.Do(func() {

		router := gin.Default()

		// Setting up CORS parameters
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "https://github.com"
			},
			MaxAge: 12 * time.Hour,
		}))

		config.OpenDBConnection()
		artbook.NewArtbookRoutes(router)

		if err := router.Run("localhost:6060"); err != nil {
			log.Fatal("Oops! An error occured while runnning the server at localhost:6060")
		}

	})

}
