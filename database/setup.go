package database

import (
	"context"
	"log"
	"time"

	"bean.com/web-service/env"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var collection *mongo.Collection

func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(env.MongoEnv("mongo")))
	//connect cluster
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer cancel()
	//defer client.Disconnect(ctx)

	//ping to db
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected")
	}
	collection = client.Database("fet_kltn").Collection("users")
}
