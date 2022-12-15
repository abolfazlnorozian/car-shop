package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"firstName" validate:"required,min=2,max=100"`
	LastName     *string            `json:"lastName" validate:"required,min=2,max=100"`
	Password     *string            `json:"password" validate:"required,min=6"`
	Email        *string            `json:"email" validate:"required"`
	Phone        *string            `json:"phone" validate:"required"`
	Token        *string            `json:"token"`
	UserType     *string            `json:"userType" validate:"required,eq=ADMIN"`
	RefreshToken *string            `json:"refreshToken"`
	CreatedAt    time.Time          `json:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"`
	UserId       string             `json:"userId"`
}
type Car struct {
	ID        int    `json:"id"  bson:"carId"`
	Name      string `json:"name"  bson:"carName"`
	Color     string `json:"color"  bson:"carColor"`
	Model     string `json:"model" bson:"carModel"`
	Price     int    `json:"price" bson:"carPrice"`
	Insurance string `json:"insurance" bson:"carInsurance"`
	Count     int    `json:"count" bson:"carCount"`
}
type Admin struct {
	ID        int       `json:"id" bson:"adminId"`
	Username  string    `json:"username" bson:"adminUsername"`
	Password  string    `json:"password" bson:"adminPassword"`
	Role      string    `json:"role" bson:"adminRole"`
	CreatedAt time.Time `json:"created_at" bson:"adminCreatedAt"`
	UpdatedAt time.Time `json:"updated_at" bson:"adminUpdatedAt"`
}
