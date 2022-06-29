package services

import (
	"context"
	"gologin/abolfazl-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.shopcollection.InsertOne(u.ctx, user)
	return err

}
func (u *UserServiceImpl) LoginUser(user *models.User) error {

	err := u.shopcollection.FindOne(u.ctx, bson.D{bson.E{Key: "user_email", Value: &user.Email}, bson.E{Key: "user_password", Value: user.Password}}).Decode(&user)

	return err

}

//********************************ADMIN*************************************************
