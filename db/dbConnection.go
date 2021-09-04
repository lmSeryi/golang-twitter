package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN Object connection */
var MongoCN = ConnectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://adm:kLbfr7B8FzkB2im@cluster0.ibcua.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* ConnectDB Doing DB Connection */
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Connection success")
	return client
}

/* CheckConnection sends ping to DB */
func CheckConnection() int {
	err := MongoCN.lient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	log.Println("Connection success")
	return 1
}
