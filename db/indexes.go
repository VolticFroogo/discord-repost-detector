package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func setupIndexes() (err error) {
	ctx := DefaultContext()

	indexName, err := Images.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{"datetime", 1},
			{"hash", 1},
			{"guild", 1},
		},
	})
	if err != nil {
		return
	} else {
		log.Printf("Created index %s", indexName)
	}

	log.Println("Setup/verified database indexes.")

	return
}
