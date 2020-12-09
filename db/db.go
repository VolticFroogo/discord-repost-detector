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
	uri  = os.Getenv("DB_URI")
	name = os.Getenv("DB_NAME")

	client *mongo.Client

	Images   *mongo.Collection
	Channels *mongo.Collection
)

func Init() (err error) {
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

	db := client.Database(name)

	Images = db.Collection("images")
	Channels = db.Collection("channels")

	log.Println("Connected to database.")

	err = setupIndexes()
	return
}

func Close() (err error) {
	err = client.Disconnect(DefaultContext())
	if err != nil {
		return fmt.Errorf("disconnecting: %s", err)
	}

	log.Println("Disconnected from database.")
	return
}

func DefaultContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return
}
