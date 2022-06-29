package services

import (
	"context"
	"gologin/abolfazl-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminServiceImpl struct {
	shopcollection *mongo.Collection
	ctx            context.Context
}

func NewAdminServiceImpl(shopcollection *mongo.Collection, ctx context.Context) Admins {
	return &AdminServiceImpl{
		shopcollection: shopcollection,
		ctx:            ctx,
	}

}

func (u *AdminServiceImpl) CreateAdmin(admin *models.Admin) error {
	_, err := u.shopcollection.InsertOne(u.ctx, admin)
	return err

}
func (u *AdminServiceImpl) LoginAdmin(admin *models.Admin) error {

	err := u.shopcollection.FindOne(u.ctx, bson.D{bson.E{Key: "admin_username", Value: &admin.Username}, bson.E{Key: "admin_password", Value: &admin.Password}, bson.E{Key: "admin_role", Value: &admin.Role}}).Decode(&admin)

	return err

}
