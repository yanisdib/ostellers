package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func openMongoDBConnection(env Env) mongo.Client {
	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://@%s:%s", dbHost, dbPort)
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongodbURI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var result bson.M
	if err := client.Database(env.DBName).RunCommand(context.TODO(), bson.D{{Key: "Ping", Value: 1}}).Decode(result); err != nil {
		panic(err)
	}

	fmt.Println("You successfully connected to MongoDB.")

	return *client
}

func closeMongoDBConnection(client *mongo.Client) {
	if client != nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}

	log.Println("Connection to MongoDB successfully closed.")
}
