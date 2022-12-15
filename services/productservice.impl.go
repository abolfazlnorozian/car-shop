package services

import (
	"context"
	"errors"
	"gologin/abolfazl-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductServiceImpl struct {
	ShopCollection *mongo.Collection
	ctx            context.Context
}

func NewProductServiceImpl(shopCollection *mongo.Collection, ctx context.Context) Products {
	return &ProductServiceImpl{
		ShopCollection: shopCollection,
		ctx:            ctx,
	}

}

func (c *ProductServiceImpl) CreateCar(car *models.Car) error {

	_, err := c.ShopCollection.InsertOne(c.ctx, car)

	return err

}
func (c *ProductServiceImpl) GetCar(name *string) (*models.Car, error) {
	var car *models.Car
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := c.ShopCollection.FindOne(c.ctx, query).Decode(&car)
	return car, err

}
func (c *ProductServiceImpl) GetAll() ([]*models.Car, error) {
	var cars []*models.Car

	cursor, err := c.ShopCollection.Find(c.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(c.ctx) {
		var car models.Car
		err := cursor.Decode(&car)
		if err != nil {
			return nil, err
		}
		cars = append(cars, &car)
	}
	if err := cursor.Err(); err != nil {
		if err != nil {
			return nil, err
		}
		cursor.Close(c.ctx)
		if len(cars) == 0 {
			return nil, errors.New("document not found")
		}

	}

	return cars, nil

}
func (c *ProductServiceImpl) UpdateCar(car *models.Car) error {
	filter := bson.D{bson.E{Key: "carId", Value: car.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "carId", Value: car.ID}, bson.E{Key: "carName", Value: car.Name}, bson.E{Key: "carColor", Value: car.Color}, bson.E{Key: "carModel", Value: car.Model}, bson.E{Key: "carPrice", Value: car.Price}, bson.E{Key: "carInsurance", Value: car.Insurance}, bson.E{Key: "carCount", Value: car.Count}}}}
	result, _ := c.ShopCollection.UpdateOne(c.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no match document found for update")
	}
	return nil

}
func (c *ProductServiceImpl) DeleteCar(name *string) error {
	filter := bson.D{bson.E{Key: "carName", Value: name}}
	result, _ := c.ShopCollection.DeleteOne(c.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no match document found for update")
	}
	return nil

}
