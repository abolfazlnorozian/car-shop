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
type UserLogin interface {
	CreateUser(*models.User) error
	LoginUser(*models.User) error
	// GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error)
	// ValidateToken(signedToken string) (claims *SignedDetails, msg string)
	UpdateAllTokens(signedToken string, signedRefreshToken string, userId string)
}
type Admins interface {
	CreateAdmin(*models.Admin) error
	LoginAdmin(*models.Admin) error
}
