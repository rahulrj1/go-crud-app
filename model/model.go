package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string
	Body  string
}

type User struct {
	Name string `json:"name"`
}
