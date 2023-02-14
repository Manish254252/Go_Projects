package models

import (
	"gopkg.in/mgo.v2/bson"
)

type user struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   int           `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    uint          `json:"age" bson:"age"`
}
