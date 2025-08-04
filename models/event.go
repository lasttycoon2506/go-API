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

func (e *Event) Save() error {
	insertEventQuery := `
		INSERT INTO events (name, description, date_time, user_id)
		VALUES (?, ?, ?, ?)
	`
	insertUsersEventsIntersectionQuery := `
		INSERT INTO usersevents (event_id, user_id)
		VALUES (?, ?)
	`

	insertEventStatement, err := db.DB.Prepare(insertEventQuery)
	if err != nil {
		return err
	}
	defer insertEventStatement.Close()

	result, err := insertEventStatement.Exec(e.Name, e.Description, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	eventId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = eventId

	insertUsersEventsStatement, err := db.DB.Prepare(insertUsersEventsIntersectionQuery)
	if err != nil {
		return err
	}
	defer insertUsersEventsStatement.Close()

	_, err = insertUsersEventsStatement.Exec(eventId, e.UserId)
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
	deleteEventQuery := `DELETE FROM events WHERE id = ?`
	deleteUsersEventsIntersectionQuery := `DELETE FROM usersevents WHERE event_id = ?`

	deleteEventsStatement, err := db.DB.Prepare(deleteEventQuery)
	if err != nil {
		return err
	}

	defer deleteEventsStatement.Close()

	_, err = deleteEventsStatement.Exec(e.ID)
	if err != nil {
		return err
	}

	deleteUsersEventsStatement, err := db.DB.Prepare(deleteUsersEventsIntersectionQuery)
	if err != nil {
		return err
	}

	defer deleteUsersEventsStatement.Close()

	_, err = deleteUsersEventsStatement.Exec(e.ID)

	return err
}
