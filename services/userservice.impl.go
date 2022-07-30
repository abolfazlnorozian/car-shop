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

func (uc *UserServiceImpl) GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {

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
func (uc *UserServiceImpl) ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
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

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	count, err := u.shopcollection.CountDocuments(u.ctx, bson.M{"email": user.Email})

	if err != nil {
		log.Panic(err)

	}
	count, err = u.shopcollection.CountDocuments(u.ctx, bson.M{"phone": user.Phone})

	if err != nil {
		log.Panic(err)

	}
	if count > 0 {
		log.Fatal(err)
	}

	_, err = u.shopcollection.InsertOne(u.ctx, user)
	return err

}
func (u *UserServiceImpl) LoginUser(user *models.User) error {
	// var _, cancle = context.WithTimeout(context.Background(), 100*time.Second)
	//var foundUser models.User

	// err := u.shopcollection.FindOne(u.ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	// defer cancle()
	// if err != nil {

	// 	return err
	// }

	// err = u.shopcollection.FindOne(u.ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)
	// if err != nil {

	// 	return err
	// }

	err := u.shopcollection.FindOne(u.ctx, bson.D{bson.E{Key: "email", Value: &user.Email}, bson.E{Key: "user_id", Value: &user.User_id}}).Decode(&user)

	return err

}

//********************************ADMIN*************************************************
