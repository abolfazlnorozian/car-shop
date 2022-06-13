package services

import (
	"context"
	"gologin/abolfazl-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserServiceImpl(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}

}
func (u *UserServiceImpl) RegistrUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err

}
func (u *UserServiceImpl) LoginUser(user *models.User) error {

	err := u.usercollection.FindOne(u.ctx, bson.D{bson.E{Key: "user_email", Value: &user.Email}, bson.E{Key: "user_password", Value: user.Password}}).Decode(&user)

	return err

}
