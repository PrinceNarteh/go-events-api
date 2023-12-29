package models

import (
	"context"
	"log"
	"time"

	"github.com/PrinceNarteh/go-events-api/mongorm"
	"go.mongodb.org/mongo-driver/bson"
)

type Event struct {
	mongorm.Model
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"user_id"`
}

var events = []Event{}
var eventsCollection = mongorm.InitDB().GetCollection("events")

func (evt Event) save() {
	events = append(events, evt)
}

func GetAllEvents() []Event {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := eventsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var events []Event
	for cursor.Next(ctx) {
		var event Event
		if err = cursor.Decode(&event); err != nil {
			log.Fatal(err)
		}
		events = append(events, event)
	}

	return events
}
