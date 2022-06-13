package models

type User struct {
	ID        int    `json:"id" bson:"user_id"`
	Firstname string `json:"firstname" bson:"user_firstname"`
	Lastname  string `json:"lastname" bson:"user_lastname"`
	Email     string `json:"email" bson:"user_email"`
	Password  string `json:"password" bson:"user_password"`
}
