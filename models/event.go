package models

import (
	"time"

	"example.com/m/v2/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

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

	_, err = statement.Exec(e.Name, e.Description, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	return err
}

func GetEvent(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`

	dbRow := db.DB.QueryRow(query, id)
	var event Event

	err := dbRow.Scan(&event.ID, &event.Name, &event.Description, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	dbRows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer dbRows.Close()

	var events []Event
	for dbRows.Next() {
		var event Event
		err := dbRows.Scan(&event.ID, &event.Name, &event.Description, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, err
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, date_time = ?
	WHERE id = ?
	`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(e.Name, e.Description, e.DateTime, e.ID)

	return err
}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(e.ID)

	return err
}
