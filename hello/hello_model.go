package hello

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HelloWorldData struct {
	Message string `json:"message"`
}

type HelloData struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Message string             `json:"message"`
}
