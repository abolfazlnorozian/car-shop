package models

type Car struct {
	ID        int    `json:"id"  bson:"car_id"`
	Name      string `json:"name"  bson:"car_name"`
	Color     string `json:"color"  bson:"car_color"`
	Model     string `json:"model" bson:"car_model"`
	Price     int    `json:"price" bson:"car_price"`
	Insurance string `json:"insurance" bson:"car_insurance"`
	Count     int    `json:"count" bson:"car_count"`
}
