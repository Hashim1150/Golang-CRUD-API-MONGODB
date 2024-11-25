package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Age    int                `bson:"age,omitempty"`
	Active bool               `bson:"active,omitempty"`
}
