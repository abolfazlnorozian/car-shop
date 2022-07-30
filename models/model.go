package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password      *string            `json:"password" validate:"required,min=6"`
	Email         *string            `json:"email" validate:"email,required"`
	Phone         *string            `json:"phone" validate:"required"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
type Car struct {
	ID        int    `json:"id"  bson:"car_id"`
	Name      string `json:"name"  bson:"car_name"`
	Color     string `json:"color"  bson:"car_color"`
	Model     string `json:"model" bson:"car_model"`
	Price     int    `json:"price" bson:"car_price"`
	Insurance string `json:"insurance" bson:"car_insurance"`
	Count     int    `json:"count" bson:"car_count"`
}
type Admin struct {
	ID        int       `json:"id" bson:"admin_id"`
	Username  string    `json:"username" bson:"admin_username"`
	Password  string    `json:"password" bson:"admin_password"`
	Role      string    `json:"role" bson:"admin_role"`
	CreatedAt time.Time `json:"created_at" bson:"admin_createdAt"`
	UpdatedAt time.Time `json:"updated_at" bson:"admin_updatedAt"`
}
