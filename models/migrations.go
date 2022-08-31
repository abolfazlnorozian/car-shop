package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	userCollection := client.Database("logindb").Collection("admins")

	userCollection.UpdateMany(
		context.TODO(),
		bson.D{},

		bson.D{{Key: "$set", Value: bson.D{{Key: "age", Value: "null"}}}},
	)

}
