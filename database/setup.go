package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"bean.com/web-service/env"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client, err = mongo.NewClient(options.Client().ApplyURI(env.MongoEnv("mongo")))
var Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

func ConnectDB() {

	//connect cluster
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(Ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(Ctx)

	//ping to db
	err = client.Ping(Ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	//list database
	databases, err := client.ListDatabaseNames(Ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}
