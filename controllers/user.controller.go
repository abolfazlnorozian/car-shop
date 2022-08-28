package controllers

import (
	"fmt"
	"log"
	"time"

	"gologin/abolfazl-api/models"
	"gologin/abolfazl-api/services"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	UserService services.UserLogin
}

func NewUserService(userservice services.UserLogin) Controller {
	return Controller{
		UserService: userservice,
	}
}

var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)

}

func (uc *Controller) RegistrUser(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	validationErr := validate.Struct(user) //errorha ra be kharej hedayat mikonad
	if validationErr != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		// return
		fmt.Println(validationErr)
	}
	password := HashPassword(*user.Password)
	user.Password = &password
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	token, refreshToken, err := services.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, user.User_id) //token ra ersal konim barye user
	user.Token = &token
	user.Refresh_token = &refreshToken

	err = uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (uc *Controller) LoginUser(ctx *gin.Context) {
	//var _, cancle = context.WithTimeout(context.Background(), 100*time.Second)

	// var foundUser models.User
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}

	foundUser, err := uc.UserService.LoginUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "this email or password is incorrect"})
		fmt.Println(err)
		// fmt.Println(foundUser)
		return

	}

	ctx.JSON(http.StatusOK, foundUser)

	// 	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
	// 		"token":         foundUser.Token,
	// 		"refresh_token": foundUser.Refresh_token,
	// 		"message":       "return successfully",
	// 	}})

}

func (uc *Controller) UserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")

	userroute.POST("/register", uc.RegistrUser)
	userroute.POST("/login", uc.LoginUser)
}
