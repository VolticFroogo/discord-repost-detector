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
}

func setupIndexes() {
	ctx := DefaultContext()

	for i := range indexes {
		indexName, err := (*indexes[i].Collection).Indexes().CreateOne(ctx, indexes[i].Model)
		if err != nil {
			panic(err)
		}

		log.Printf("Created index %s in collection %s", indexName, (*indexes[i].Collection).Name())
	}

	log.Println("Setup/verified database indexes.")
}
