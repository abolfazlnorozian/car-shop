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
	server          *gin.Engine
	carservice      services.CarShopService
	LogController   controllers.LogController
	ctx             context.Context
	mongocollection *mongo.Collection
	mongoclient     *mongo.Client
	err             error
)

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

	mongocollection = mongoclient.Database("logindb").Collection("logins")
	carservice = services.NewServiceImpl(mongocollection, ctx)

	LogController = controllers.New(carservice)
	server = gin.Default()
}

func main() {

	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("v1")
	LogController.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":3000"))

}
