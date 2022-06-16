package controllers

import (
	"net/http"

	"example.com/cars-api/models"
	"example.com/cars-api/services"
	"github.com/gin-gonic/gin"
)

type CarController struct {
	CarService services.CarService
}

func New(carservice services.CarService) CarController {
	return CarController{
		CarService: carservice,
	}

}
func (ca *CarController) CreateCar(ctx *gin.Context) {
	var car models.Car
	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"massage": err.Error()})
		return

	}
	err := ca.CarService.CreateCar(&car)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"massage": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"massage": "success"})

}
func (ca *CarController) GetCar(ctx *gin.Context) {
	carname := ctx.Param("name")
	car, err := ca.CarService.GetCar(&carname)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"massage": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, car)

}
func (ca *CarController) GetAll(ctx *gin.Context) {
	cars, err := ca.CarService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"massage": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cars)

}
func (ca *CarController) UpdateCar(ctx *gin.Context) {
	var car models.Car
	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"massage": err.Error()})
		return
	}
	err := ca.CarService.UpdateCar(&car)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"massage": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"massage": "success"})

}
func (ca *CarController) DeleteCar(ctx *gin.Context) {
	username := ctx.Param("name")
	err := ca.CarService.DeleteCar(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"massage": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"massage": "success"})

}
func (ca *CarController) RegisterCarRoutes(rg *gin.RouterGroup) {
	CarRoute := rg.Group("/car")
	CarRoute.POST("/create", ca.CreateCar)
	CarRoute.GET("/get/:name", ca.GetCar)
	CarRoute.GET("/getall", ca.GetAll)
	CarRoute.PATCH("/update", ca.UpdateCar)
	CarRoute.DELETE("/delete/:name", ca.DeleteCar)

}
