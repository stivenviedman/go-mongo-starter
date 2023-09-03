package database

import (
	"context"

	"github.com/stivenviedman/go-mongo-starter/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init() {
	if mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Env.MONGO_URI)); err != nil {
		panic(err)
	} else {
		client = mongoClient
	}
}

func Close() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
