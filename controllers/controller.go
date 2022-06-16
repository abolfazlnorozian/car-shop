package controllers

import (
	"gologin/abolfazl-api/models"
	"gologin/abolfazl-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type UserLogin struct {
// 	Email string `json:"email" bson:"user_email"`
// }

type LogController struct {
	Carservice services.CarShopService
}

func New(carservice services.CarShopService) LogController {
	return LogController{
		Carservice: carservice,
	}
}
func (uc *LogController) CreateCar(ctx *gin.Context) {
	var car models.Car
	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"massage": err.Error()})
		return

	}
	err := uc.Carservice.CreateCar(&car)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"massage": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"massage": "success"})

}
func (uc *LogController) GetCar(ctx *gin.Context) {
	carname := ctx.Param("name")
	car, err := uc.Carservice.GetCar(&carname)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"massage": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, car)

}
func (uc *LogController) GetAll(ctx *gin.Context) {
	cars, err := uc.Carservice.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"massage": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cars)

}
func (uc *LogController) UpdateCar(ctx *gin.Context) {
	var car models.Car
	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"massage": err.Error()})
		return
	}
	err := uc.Carservice.UpdateCar(&car)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"massage": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"massage": "success"})

}
func (uc *LogController) DeleteCar(ctx *gin.Context) {
	username := ctx.Param("name")
	err := uc.Carservice.DeleteCar(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"massage": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"massage": "success"})

}

func (uc *LogController) RegistrUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.Carservice.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (uc *LogController) LoginUser(ctx *gin.Context) {

	var user models.User
	//var foundUser models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this email not exist"})
		return

	}
	err := uc.Carservice.LoginUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "this email not exist"})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (uc *LogController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/createCar", uc.CreateCar)
	userroute.GET("/getCar/:name", uc.GetCar)
	userroute.GET("/getAll", uc.GetAll)
	userroute.PATCH("/updateCar", uc.UpdateCar)
	userroute.DELETE("/deleteCar/:name", uc.DeleteCar)
	userroute.POST("/register", uc.RegistrUser)
	userroute.POST("/login", uc.LoginUser)
}
