package services

import (
	"context"
	"gologin/abolfazl-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminServiceImpl struct {
	ShopCollection *mongo.Collection
	ctx            context.Context
}

func NewAdminServiceImpl(shopCollection *mongo.Collection, ctx context.Context) Admins {
	return &AdminServiceImpl{
		ShopCollection: shopCollection,
		ctx:            ctx,
	}

}

func (u *AdminServiceImpl) CreateAdmin(admin *models.Admin) error {
	_, err := u.ShopCollection.InsertOne(u.ctx, admin)
	return err

}
func (u *AdminServiceImpl) LoginAdmin(admin *models.Admin) error {

	err := u.ShopCollection.FindOne(u.ctx, bson.D{bson.E{Key: "adminUsername", Value: &admin.Username}, bson.E{Key: "adminPassword", Value: &admin.Password}, bson.E{Key: "adminRole", Value: &admin.Role}}).Decode(&admin)

	return err

}
