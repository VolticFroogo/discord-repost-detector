package model

import (
	"github.com/VolticFroogo/discord-repost-detector/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Image struct {
	ID       primitive.ObjectID `bson:"_id"`
	DateTime primitive.DateTime `bson:"datetime"`
	Hash     string             `bson:"hash"`
	Guild    uint64             `bson:"guild"`
	Channel  uint64             `bson:"channel"`
	Message  uint64             `bson:"message"`
	User     uint64             `bson:"user"`
}

func (image Image) Insert() {
	now := time.Now()

	image.ID = primitive.NewObjectIDFromTimestamp(now)
	image.DateTime = primitive.NewDateTimeFromTime(now)

	// Insert the document.
	_, err := db.Images.InsertOne(db.DefaultContext(), image)
	if err != nil {
		panic(err)
	}
}

func (image Image) FindMatch() (matches []Image) {
	ctx := db.DefaultContext()

	cursor, err := db.Images.Aggregate(ctx, []bson.M{
		{
			"$match": bson.M{
				"hash":  image.Hash,
				"guild": image.Guild,
			},
		},
		{
			"$sort": bson.M{
				"datetime": 1,
			},
		},
	})
	if err != nil {
		panic(err)
	}

	err = cursor.All(ctx, &matches)
	if err != nil {
		panic(err)
	}

	return
}
