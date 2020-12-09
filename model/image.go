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

func (image Image) Insert() (err error) {
	now := time.Now()

	image.ID = primitive.NewObjectIDFromTimestamp(now)
	image.DateTime = primitive.NewDateTimeFromTime(now)

	// Insert the document.
	_, err = db.Images.InsertOne(db.DefaultContext(), image)
	return
}

func (image Image) FindMatch() (matches []Image, err error) {
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
		return
	}

	err = cursor.All(ctx, &matches)
	return
}
