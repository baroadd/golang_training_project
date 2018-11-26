package models 

import (
	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	Avaliable int `bson:"avaliable" json:"avaliable"`
	Speaker string `bson:"speaker" json:"speaker"`
	Date        string `bson:"date" json:"date"`
	Round string`bson:"round" json:"round"`
	User [] string `bson:"user" json:"user"`
}