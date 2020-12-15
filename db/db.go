package db

import (
	"context"
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

func Init() {
	opts := options.Client().ApplyURI(uri)

	var err error
	client, err = mongo.NewClient(opts)
	if err != nil {
		panic(err)
	}

	ctx := DefaultContext()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, opts.ReadPreference)
	if err != nil {
		panic(err)
	}

	db := client.Database(name)

	Images = db.Collection("images")
	Channels = db.Collection("channels")

	log.Println("Connected to database.")

	setupIndexes()
}

func Close() {
	err := client.Disconnect(DefaultContext())
	if err != nil {
		panic(err)
	}

	log.Println("Disconnected from database.")
}

func DefaultContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return
}
