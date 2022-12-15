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
	UserService    services.UserLogin
	ProductService services.Products
	AdminService   services.Admins

	Controller        controllers.Controller
	ProductController controllers.ProductController
	AdminController   controllers.AdminControllers

	ctx               context.Context
	UserCollection    *mongo.Collection
	ProductCollection *mongo.Collection
	AdminCollection   *mongo.Collection
	BrandCollection   *mongo.Collection

	MongoClient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")

	MongoClient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongo connection established.")

	ProductCollection = MongoClient.Database("logindb").Collection("products")
	UserCollection = MongoClient.Database("logindb").Collection("users")
	AdminCollection = MongoClient.Database("logindb").Collection("admins")
	BrandCollection = MongoClient.Database("logindb").Collection("brands")

	UserService = services.NewUserServiceImpl(UserCollection, ctx)
	ProductService = services.NewProductServiceImpl(ProductCollection, ctx)
	AdminService = services.NewAdminServiceImpl(AdminCollection, ctx)

	Controller = controllers.NewUserService(UserService)
	ProductController = controllers.NewProductService(ProductService, UserService)
	AdminController = controllers.NewAdminService(AdminService)

	server = gin.Default()
	server.MaxMultipartMemory = 8 << 20
}

func main() {

	defer MongoClient.Disconnect(ctx)
	basepath := server.Group("v1")
	Controller.UserRoutes(basepath)

	ProductController.ProductRoutes(basepath)
	AdminController.AdminRoutes(basepath)

	log.Fatal(server.Run(":3000"))

}
