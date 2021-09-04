package db

import (
	"context"
	"log"
	"time"

	"github.com/lmSeryi/golang-twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRegist(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjId, _ := result.InsertedID.(primitive.ObjectID)
	log.Println(result, ObjId)
	return ObjId.String(), true, nil
}
