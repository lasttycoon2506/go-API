package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	DateTime    time.Time
	UserId      int
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}
