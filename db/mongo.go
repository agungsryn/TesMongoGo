package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func ConnectMongo() error {
	opts := options.ClientOptions{
		// ConnectTimeout: helper.GetPointer(60 * time.Second),
	}

	opts.SetDirect(true)
	opts.SetMaxPoolSize(uint64(100))
	opts.SetMinPoolSize(10)
	opts.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), &opts)
	if err != nil {
		log.Fatalf("error while connecting to db: %v", err)
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	db = client.Database("goTest")
	InitCollections()
	log.Println("2. Connected to mongodb at:", "local", "environment")
	return nil
}

func InitCollections() {
	UserCollection = db.Collection("users")
	FormulirCollection = db.Collection("formulirs")
}