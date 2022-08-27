package services

import (
	"gologin/abolfazl-api/models"
)

type Products interface {
	CreateCar(*models.Car) error
	GetCar(*string) (*models.Car, error)
	GetAll() ([]*models.Car, error)
	UpdateCar(*models.Car) error
	DeleteCar(*string) error
}
type Migrations interface {
	Migrations()
}
type UserLogin interface {
	CreateUser(*models.User) error
	LoginUser(*models.User) (*models.User, error)
	//GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error)

	UpdateAllTokens(string, string, string)
}

type Admins interface {
	CreateAdmin(*models.Admin) error
	LoginAdmin(*models.Admin) error
}
