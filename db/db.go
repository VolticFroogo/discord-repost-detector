package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client

	Images *mongo.Collection
)

func Init() (err error) {
	uri := os.Getenv("DB_URI")
	opts := options.Client().ApplyURI(uri)

	client, err = mongo.NewClient(opts)
	if err != nil {
		return fmt.Errorf("creating client with options: %s", err)
	}

	ctx := DefaultContext()
	err = client.Connect(ctx)
	if err != nil {
		return fmt.Errorf("connecting to server: %s", err)
	}

	err = client.Ping(ctx, opts.ReadPreference)
	if err != nil {
		return fmt.Errorf("pinging server: %s", err)
	}

	dbName := os.Getenv("DB_NAME")
	db := client.Database(dbName)

	Images = db.Collection("images")

	log.Println("Connected to database.")

	err = setupIndexes()
	return
}

func Close() (err error) {
	err = client.Disconnect(DefaultContext())
	if err != nil {
		return fmt.Errorf("disconnecting: %s", err)
	}

	return
}

func DefaultContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return
}
