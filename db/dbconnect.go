package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"api.dardasha.com/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(utils.EnvMongoURL()))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conneted to mongodb")
	return client
}

var DB *mongo.Client = ConnectToDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("dardashaDB").Collection(collectionName)
	return collection
}
