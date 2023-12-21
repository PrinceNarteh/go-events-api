package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (evt Event) save() {
	events = append(events, evt)
}

func GetAllEvents() []Event {
	return events
}
