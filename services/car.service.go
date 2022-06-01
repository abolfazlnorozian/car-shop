package services

import "example.com/cars-api/models"

type CarService interface {
	CreateCar(*models.Car) error
	GetCar(*string) (*models.Car, error)
	GetAll() ([]*models.Car, error)
	UpdateCar(*models.Car) error
	DeleteCar(*string) error
}
