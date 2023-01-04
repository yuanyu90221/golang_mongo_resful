package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	City  string `json:"city,omitempty" bson:"city,omitempty"`
	State string `json:"state,omitempty" bson:"state,omitempty"`
}
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  string             `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Address   *Address           `json:"address,omitempty" bson:"address,omitempty"`
}
