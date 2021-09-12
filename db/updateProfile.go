package db

import (
	"context"
	"time"

	"github.com/lmSeryi/golang-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* UpdateProfile */
func UpdateProfile(u models.User, Id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	regist := make(map[string]interface{})
	if len(u.Name) > 0 {
		regist["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		regist["lastName"] = u.LastName
	}
	regist["birthday"] = u.Birthday
	if len(u.Avatar) > 0 {
		regist["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		regist["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		regist["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		regist["location"] = u.Location
	}
	if len(u.Email) > 0 {
		regist["email"] = u.Email
	}
	if len(u.WebSite) > 0 {
		regist["webSite"] = u.WebSite
	}
	updtString := bson.M{
		"$set": regist,
	}
	objId, _ := primitive.ObjectIDFromHex(Id)
	filter := bson.M{"_id": bson.M{"$eq": objId}}

	_, err := col.UpdateOne(ctx, filter, updtString)

	if err != nil {
		return false, err
	}
	return true, nil
}
