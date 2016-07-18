package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	BaseDBmodel
	Id       bson.ObjectId `bson:"_id" form:"-" `
	Name     string        `form:"name" valid:"Required"`
	Premises []Premise
	Results  []Result
}
