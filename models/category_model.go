package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	Id     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name" validate:"required"`
	ArName string             `bson:"arName" validate:"required"`
}
