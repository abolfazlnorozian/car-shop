package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MigrationServiceImpl struct {
	shopcollection *mongo.Collection
	ctx            context.Context
}

// Migrations implements Migrations

func NewMigrationServiceImpl(shopcollection *mongo.Collection, ctx context.Context) Migrations {
	return &MigrationServiceImpl{
		shopcollection: shopcollection,
		ctx:            ctx,
	}

}

func (uc *MigrationServiceImpl) Migrations() error {
	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	panic(err)
	// }
	// if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
	// 	panic(err)
	// }
	// userCollection := client.Database("logindb").Collection("products")
	filter := bson.D{}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "company", Value: "null"}}}}

	result, _ := uc.shopcollection.UpdateMany(uc.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no match document found for update")
	}

	return nil

}
