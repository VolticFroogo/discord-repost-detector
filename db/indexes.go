package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type index struct {
	Collection **mongo.Collection
	Model      mongo.IndexModel
}

var indexes = []index{
	{
		Collection: &Images,
		Model: mongo.IndexModel{
			Keys: bson.D{
				{"datetime", 1},
				{"hash", 1},
				{"guild", 1},
			},
		},
	},
	{
		Collection: &Channels,
		Model: mongo.IndexModel{
			Keys: bson.D{
				{"guild", 1},
				{"channel", 1},
			},
		},
	},
}

func setupIndexes() (err error) {
	ctx := DefaultContext()

	for i := range indexes {
		indexName, err := (*indexes[i].Collection).Indexes().CreateOne(ctx, indexes[i].Model)
		if err != nil {
			return err
		}

		log.Printf("Created index %s", indexName)
	}

	log.Println("Setup/verified database indexes.")

	return
}
