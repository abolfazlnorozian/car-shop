package services

import "gologin/abolfazl-api/models"

type Products interface {
	CreateCar(*models.Car) error
	GetCar(*string) (*models.Car, error)
	GetAll() ([]*models.Car, error)
	UpdateCar(*models.Car) error
	DeleteCar(*string) error
}
type UserLogin interface {
	CreateUser(*models.User) error
	LoginUser(*models.User) error
}
type Admins interface {
	CreateAdmin(*models.Admin) error
	LoginAdmin(*models.Admin) error
}
type Download interface {
	DownloadFile(*string) error
}
