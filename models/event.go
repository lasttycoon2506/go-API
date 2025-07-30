package models

import (
	"time"

	"example.com/m/v2/db"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (e Event) Save() error {
	insertQuery := `
		INSERT INTO events (name, description, date_time, user_id)
		VALUES (?, ?, ?, ?)
	`

	statement, err := db.DB.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer statement.Close()

	statement.Exec(e.Name, e.Description, e.DateTime, e.UserId)
}

func GetAllEvents() []Event {
	return events
}
