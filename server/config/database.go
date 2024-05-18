package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func OpenDBConnection() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Establishing connection to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(GetMongoURI()))
	if err != nil {
		log.Fatal("Error occured while establishing connection to MongoDB")
		return nil
	}

	// Checking MongoDB database established connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {

		log.Fatal(err)

		// Disconnecting from client if ping didn't work
		defer func() {
			err := client.Disconnect(ctx)
			if err != nil {
				log.Fatal("Error occured while disconnecting from MongoDB")
			}
		}()

		return nil

	}

	log.Println("Connection to MongoDB established successfully")

	return client

}

// Instance of mongo.Client
var MongoClient *mongo.Client = OpenDBConnection()

func GetCollection(collectionName string) *mongo.Collection {
	dbName := GetDatabaseName()
	return MongoClient.Database(dbName).Collection(collectionName)
}
