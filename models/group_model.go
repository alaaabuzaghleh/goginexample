package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name" validate:"required"`
	ArName   string             `bson:"arName" validate:"required"`
	Desc     string             `bson:"desc" validate:"required"`
	ArDesc   string             `bson:"arDesc" validate:"required"`
	SmallPic string             `bson:"smallPic" validate:"required"`
	LargePic string             `bson:"largePic" validate:"required"`
}
