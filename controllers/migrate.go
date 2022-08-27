package controllers

import (
	"gologin/abolfazl-api/models"
	"gologin/abolfazl-api/services"
)

type MigrateController struct {
	MigrateService services.Migrations
}

func NewMigrateService(migrateservice services.Migrations) MigrateController {
	return MigrateController{
		MigrateService: migrateservice,
	}

}
func (uc *MigrateController) Migrations() {
	var field models.Car

	err := uc.MigrateService.Migrations(field)
	if err != nil {

		return err
	}
}
