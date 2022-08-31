package main

import (
	"context"
	"fmt"
	"log"

	"gologin/abolfazl-api/controllers"

	"gologin/abolfazl-api/services"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userservice    services.UserLogin
	productservice services.Products
	adminservice   services.Admins

	controller        controllers.Controller
	productcontroller controllers.ProductController
	admincontroller   controllers.AdminControllers

	ctx               context.Context
	usercollection    *mongo.Collection
	productcollection *mongo.Collection
	admincollection   *mongo.Collection
	brandcollection   *mongo.Collection

	mongoclient *mongo.Client
	err         error
)

// type E struct {
// 	Key   string
// 	Value interface{}
// }

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")

	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongo connection established.")

	productcollection = mongoclient.Database("logindb").Collection("products")
	usercollection = mongoclient.Database("logindb").Collection("users")
	admincollection = mongoclient.Database("logindb").Collection("admins")
	brandcollection = mongoclient.Database("logindb").Collection("brands")

	userservice = services.NewUserServiceImpl(usercollection, ctx)
	productservice = services.NewProductServiceImpl(productcollection, ctx)
	adminservice = services.NewAdminServiceImpl(admincollection, ctx)

	controller = controllers.NewUserService(userservice)
	productcontroller = controllers.NewProductService(productservice)
	admincontroller = controllers.NewAdminService(adminservice)

	server = gin.Default()
	server.MaxMultipartMemory = 8 << 20
}

func main() {

	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("v1")
	controller.UserRoutes(basepath)

	productcontroller.ProductRoutes(basepath)
	admincontroller.AdminRoutes(basepath)

	log.Fatal(server.Run(":3000"))

}
