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

type UserController struct {
	UserService services.UserLogin
}

func NewUserService(userservice services.UserLogin) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) RegistrUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (uc *UserController) LoginUser(ctx *gin.Context) {

	var user models.User
	//var foundUser models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "this email not exist"})
		return

	}
	err := uc.UserService.LoginUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "this email not exist"})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (uc *UserController) UserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")

	userroute.POST("/register", uc.RegistrUser)
	userroute.POST("/login", uc.LoginUser)
}
