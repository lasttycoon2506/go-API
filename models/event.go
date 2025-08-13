package models

import (
	"time"

	"example.com/m/v2/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	DateTime    time.Time
	UserId      int64
	StoryId     int64 `binding:"required"`
}

func (e *Event) Save() error {
	insertEventQuery := `
		INSERT INTO events (name, description, date_time, user_id, story_id)
		VALUES (?, ?, ?, ?, ?)
	`

	insertEventStatement, err := db.DB.Prepare(insertEventQuery)
	if err != nil {
		return err
	}
	defer insertEventStatement.Close()

	e.DateTime = time.Now()
	result, err := insertEventStatement.Exec(e.Name, e.Description, e.DateTime, e.UserId, e.StoryId)
	if err != nil {
		return err
	}

	eventId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = eventId

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

func GetEventsGroupedByUser() (map[int64]map[int64][]Event, error) {
	query := `SELECT * FROM events ORDER BY date_time DESC`

	dbRows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer dbRows.Close()

	grouped := make(map[int64]map[int64][]Event)
	for dbRows.Next() {
		var event Event
		err := dbRows.Scan(&event.ID, &event.Name, &event.Description, &event.DateTime, &event.UserId, &event.StoryId)
		if err != nil {
			return nil, err
		}
		if grouped[event.UserId] == nil {
			grouped[event.UserId] = make(map[int64][]Event)
		}
		grouped[event.UserId][event.StoryId] = append(grouped[event.UserId][event.StoryId], event)
	}

	return grouped, nil
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
	if err != nil {
		return err
	}

	return err
}
