package routes

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func DBInstance() *mongo.Client {
	mongoDB := "mongodb://localhost:27017/caloriesdb"

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDB))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MONGODB")
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) (*mongo.Collection) {
	var collection *mongo.Collection = client.Database("caloriesdb").Collection(collectionName)
	return collection
}