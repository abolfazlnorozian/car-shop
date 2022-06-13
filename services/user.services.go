package services

import "gologin/abolfazl-api/models"

type UserService interface {
	RegistrUser(*models.User) error
	LoginUser(*models.User) error
}
