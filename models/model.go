package models

type User struct {
	ID        int    `json:"id" bson:"user_id"`
	Firstname string `json:"firstname" bson:"user_firstname"`
	Lastname  string `json:"lastname" bson:"user_lastname"`
	Email     string `json:"email" bson:"user_email"`
	Password  string `json:"password" bson:"user_password"`
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
