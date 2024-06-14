package player

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlayerData struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Tour      string             `json:"tour"`
	Country   string             `json:"country"`
}
