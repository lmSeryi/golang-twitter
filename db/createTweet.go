package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/lmSeryi/golang-twitter/models"
)

func CreateTweet(t models.CreateTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	payload := bson.M{
		"userId":  t.UserId,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := col.InsertOne(ctx, payload)
	if err != nil {
		return "", false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.String(), true, nil
}
