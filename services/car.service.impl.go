package services

import (
	"context"
	"errors"

	"example.com/cars-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CarServiceImpl struct {
	carcollection *mongo.Collection
	ctx           context.Context
}

func NewCarService(carcollection *mongo.Collection, ctx context.Context) CarService {
	return &CarServiceImpl{
		carcollection: carcollection,
		ctx:           ctx,
	}

}
func (c *CarServiceImpl) CreateCar(car *models.Car) error {
	_, err := c.carcollection.InsertOne(c.ctx, car)
	return err

}
func (c *CarServiceImpl) GetCar(name *string) (*models.Car, error) {
	var car *models.Car
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := c.carcollection.FindOne(c.ctx, query).Decode(&car)
	return car, err

}
func (c *CarServiceImpl) GetAll() ([]*models.Car, error) {
	var cars []*models.Car
	cursor, err := c.carcollection.Find(c.ctx, bson.D{{}})
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
func (c *CarServiceImpl) UpdateCar(car *models.Car) error {
	filter := bson.D{bson.E{Key: "car_name", Value: car.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "car_id", Value: car.ID}, bson.E{Key: "car_name", Value: car.Name}, bson.E{Key: "car_color", Value: car.Color}, bson.E{Key: "car_model", Value: car.Model}, bson.E{Key: "car_price", Value: car.Price}, bson.E{Key: "car_insurance", Value: car.Insurance}, bson.E{Key: "car_count", Value: car.Count}}}}
	result, _ := c.carcollection.UpdateOne(c.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no match document found for update")
	}
	return nil

}
func (c *CarServiceImpl) DeleteCar(name *string) error {
	filter := bson.D{bson.E{Key: "car_name", Value: name}}
	result, _ := c.carcollection.DeleteOne(c.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no match document found for update")
	}
	return nil

}
