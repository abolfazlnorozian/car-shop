package services

import "gologin/abolfazl-api/models"

type CarShopService interface {
	CreateUser(*models.User) error
	LoginUser(*models.User) error
	CreateCar(*models.Car) error
	GetCar(*string) (*models.Car, error)
	GetAll() ([]*models.Car, error)
	UpdateCar(*models.Car) error
	DeleteCar(*string) error
}
