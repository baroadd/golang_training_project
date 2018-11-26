package models 

import (
	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	Limit       int `bson:"limit" json:"limit"`
	Avaliable int `bson:"avaliable" json:"avaliable"`
	Speaker string `bson:"speaker" json:"speaker"`
	Date        string `bson:"date" json:"date"`
}