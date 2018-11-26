package dao

import (
	"log"
	. "reserve-service/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type EventsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "events"
)

func (e *EventsDAO) Connect() {
	session, err := mgo.Dial(e.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(e.Database)
}

func (e *EventsDAO) FindAll() ([]Event, error) {
	var event []Event
	err := db.C(COLLECTION).Find(bson.M{}).All(&event)
	return event, err
}

func (e *EventsDAO) Insert(event Event) error {
	err := db.C(COLLECTION).Insert(&event)
	return err
}

func (e *EventsDAO) FindById(id string) (Event, error) {
	var event Event
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&event)
	return event, err
}

func (e *EventsDAO) Update(event Event) error {
	err := db.C(COLLECTION).UpdateId(event.ID, &event)
	return err
}

func (e *EventsDAO) Delete(event Event) error {
	err := db.C(COLLECTION).Remove(&event)
	return err
}