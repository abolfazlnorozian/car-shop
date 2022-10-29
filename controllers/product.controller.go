package controllers

import (
	"gologin/abolfazl-api/middleware"
	"gologin/abolfazl-api/models"
	"gologin/abolfazl-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService services.Products
	UserService    services.UserLogin
}

func NewProductService(productservice services.Products, UserService services.UserLogin) ProductController {
	return ProductController{
		ProductService: productservice,
		UserService:    UserService,
	}

}

func (uc *ProductController) CreateCar(ctx *gin.Context) {

	var car models.Car

	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"massage": err.Error()})
		return

	}
	if err := middleware.CheckUserType(ctx, "ADMIN"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	err := uc.ProductService.CreateCar(&car)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"massage": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"massage": "success"})

}
func (uc *ProductController) GetCar(ctx *gin.Context) {
	carname := ctx.Param("name")
	car, err := uc.ProductService.GetCar(&carname)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"massage": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, car)

}
func (uc *ProductController) GetAll(ctx *gin.Context) {
	cars, err := uc.ProductService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"massage": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cars)

}
func (uc *ProductController) UpdateCar(ctx *gin.Context) {
	var car models.Car
	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"massage": err.Error()})
		return
	}
	err := uc.ProductService.UpdateCar(&car)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"massage": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"massage": "success"})

}
func (uc *ProductController) DeleteCar(ctx *gin.Context) {
	username := ctx.Param("name")
	err := uc.ProductService.DeleteCar(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"massage": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"massage": "success"})

}
func (uc *ProductController) ProductRoutes(cg *gin.RouterGroup) {
	carroute := cg.Group("/car")
	carroute.Use(middleware.Authenticate())
	carroute.POST("/createCar", uc.CreateCar)
	carroute.GET("/getCar/:name", uc.GetCar)
	carroute.GET("/all", uc.GetAll)
	carroute.PATCH("/updateCar", uc.UpdateCar)
	carroute.DELETE("/deleteCar/:name", uc.DeleteCar)

}
