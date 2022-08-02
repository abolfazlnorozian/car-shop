package controllers

import (
	"fmt"

	"gologin/abolfazl-api/models"
	"gologin/abolfazl-api/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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

var validate = validator.New()

func (uc *UserController) Authenticate(c *gin.Context) {

	clientToken := c.Request.Header.Get("token")
	if clientToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("no Authorization header provided")})
		c.Abort() //baraye amniat
		return
	}
	//claims tamami etelaate dorosti ke ma migirim hastan
	claims, err := services.ValidateToken(clientToken) //token haei ke ok hastan baraye userha
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}
	c.Set("email", claims.Email)
	c.Set("first_name", claims.First_name)
	c.Set("last_name", claims.Last_name)
	c.Set("uid", claims.Uid)
	c.Set("user_type", claims.User_type)
	c.Next()

}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)

}

func (uc *UserController) RegistrUser(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	validationErr := validate.Struct(user) //errorha ra be kharej hedayat mikonad
	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
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

func (uc *UserController) LoginUser(ctx *gin.Context) {
	//var _, cancle = context.WithTimeout(context.Background(), 100*time.Second)

	var user models.User
	//var foundUser models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

	}

	// err := uc.UserService.LoginUser(&user)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "this email not exist"})
	// 	return

	// }

	err := uc.UserService.LoginUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "this email not exist"})
		return

	}

	ctx.JSON(http.StatusOK, user)

	// ctx.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
	// 	"token":         foundUser.Token,
	// 	"refresh_token": foundUser.Refresh_token,
	// 	"message":       "return successfully",
	// }})

}

func (uc UserController) UserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")

	userroute.POST("/register", uc.RegistrUser)
	userroute.POST("/login", uc.LoginUser)
}
