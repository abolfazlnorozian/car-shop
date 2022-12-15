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
	ShopCollection *mongo.Collection
	ctx            context.Context
}

func NewUserServiceImpl(shopCollection *mongo.Collection, ctx context.Context) UserLogin {
	return &UserServiceImpl{
		ShopCollection: shopCollection,
		ctx:            ctx,
	}

}

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Uid       string
	UserType  string
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Uid:       uid,
		UserType:  userType,
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
	updateObj = append(updateObj, bson.E{Key: "refreshToken", Value: signedRefreshToken})
	UpdatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updatedAt", Value: UpdatedAt})
	upsert := true
	filter := bson.M{"userId": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := uc.ShopCollection.UpdateOne(
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
	count, err := uc.ShopCollection.CountDocuments(uc.ctx, bson.M{"email": user.Email})

	if err != nil {
		log.Panic(err)

	}
	count, err = uc.ShopCollection.CountDocuments(uc.ctx, bson.M{"phone": user.Phone})

	if err != nil {
		log.Panic(err)

	}

	if count > 0 {
		log.Fatal(err)
	}

	_, err = uc.ShopCollection.InsertOne(uc.ctx, user)
	return err

}
func (uc *UserServiceImpl) LoginUser(user *models.User) (*models.User, error) {

	var foundUser *models.User

	err := uc.ShopCollection.FindOne(uc.ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		return nil, err
	}

	err = uc.ShopCollection.FindOne(uc.ctx, bson.M{"password": user.Password}).Decode(&foundUser)
	passwordIsValid, _ := VerifyPassword(*user.Password, *foundUser.Password)

	if passwordIsValid != true {

		return nil, err

	}
	if foundUser.Email == nil {

		fmt.Println("user not found")
		return nil, err

	}
	token, refreshToken, err := GenerateAllTokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, *foundUser.UserType, foundUser.UserId)
	uc.UpdateAllTokens(token, refreshToken, foundUser.UserId)
	err = uc.ShopCollection.FindOne(uc.ctx, bson.M{"userId": foundUser.UserId}).Decode(&foundUser)

	if err != nil {
		return nil, err
	}

	return foundUser, err
}
