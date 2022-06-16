package main

import (
	"context"
	"fmt"
	"log"

	"example.com/cars-api/controllers"
	"example.com/cars-api/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server        *gin.Engine
	carservice    services.CarService
	CarController controllers.CarController
	ctx           context.Context
	carcollection *mongo.Collection
	mongoclient   *mongo.Client
	err           error
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
	fmt.Println("mongo connect established")
	carcollection = mongoclient.Database("cardb").Collection("cras")
	carservice = services.NewCarService(carcollection, ctx)
	CarController = controllers.New(carservice)
	server = gin.Default()

}

func main() {

	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("v1")
	CarController.RegisterCarRoutes(basepath)
	log.Fatal(server.Run(":3000"))

}
