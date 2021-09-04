package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/lmSeryi/golang-twitter/models"
)

/* UserExists check is user exists */
func UserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	condition := bson.M{"email": email}
	var result models.User
	err := col.FindOne(ctx, condition).Decode(&result)
	Id := result.Id.Hex()
	if err != nil {
		return result, false, Id
	}
	return result, true, Id
}
