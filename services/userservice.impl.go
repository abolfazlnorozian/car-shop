package services

import (
	"context"
	"fmt"

	"gologin/abolfazl-api/models"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	shopcollection *mongo.Collection
	ctx            context.Context
}

func NewUserServiceImpl(shopcollection *mongo.Collection, ctx context.Context) UserLogin {
	return &UserServiceImpl{
		shopcollection: shopcollection,
		ctx:            ctx,
	}

}

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		User_type:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil

		},
	)
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}
	return claims, msg

}
func (uc *UserServiceImpl) UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	var ctx, cancle = context.WithTimeout(context.Background(), 100*time.Second)
	var updateObj primitive.D
	//append token and refreshtoken in update object
	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedRefreshToken})
	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: Updated_at})
	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := uc.shopcollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{Key: "$set", Value: updateObj},
		},
		&opt,
	)
	defer cancle()
	if err != nil {
		log.Panic(err)
		return
	}
	return

}
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		msg = fmt.Sprintf("email of password is incorrect")
		check = false
	}
	return check, msg

}

func (uc *UserServiceImpl) CreateUser(user *models.User) error {
	count, err := uc.shopcollection.CountDocuments(uc.ctx, bson.M{"email": user.Email})

	if err != nil {
		log.Panic(err)

	}
	count, err = uc.shopcollection.CountDocuments(uc.ctx, bson.M{"phone": user.Phone})

	if err != nil {
		log.Panic(err)

	}
	if count > 0 {
		log.Fatal(err)
	}

	_, err = uc.shopcollection.InsertOne(uc.ctx, user)
	return err

}
func (uc *UserServiceImpl) LoginUser(user *models.User) (*models.User, error) {

	var foundUser *models.User

	err := uc.shopcollection.FindOne(uc.ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		return nil, err
	}

	passwordIsValid, _ := VerifyPassword(*user.Password, *foundUser.Password)
	err = uc.shopcollection.FindOne(uc.ctx, bson.M{"password": user.Password}).Decode(&foundUser)

	if passwordIsValid != true {

		return nil, err

	}
	if foundUser.Email == nil {

		fmt.Println("user not found")
		return nil, err

	}
	token, refreshToken, err := GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, foundUser.User_id)
	uc.UpdateAllTokens(token, refreshToken, foundUser.User_id)
	err = uc.shopcollection.FindOne(uc.ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)

	if err != nil {
		return nil, err
	}

	return foundUser, err
}

//********************************ADMIN*************************************************
