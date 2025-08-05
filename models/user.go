package models

import (
	"errors"

	"example.com/m/v2/db"
	"example.com/m/v2/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, err = statement.Exec(u.Email, hashedPassword)

	return err
}

func (u *User) Verify() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	dbRow := db.DB.QueryRow(query, u.Email)

	var hashedPasswordInDb string
	err := dbRow.Scan(&u.ID, &hashedPasswordInDb)
	if err != nil {
		return errors.New("invalid creds")
	}

	passwordIsValid := utils.CheckHashedPassword(u.Password, hashedPasswordInDb)
	if !passwordIsValid {
		return errors.New("invalid creds")
	}

	return nil
}

func GetUserEvents(userId int64) ([]Event, error) {
	query := `SELECT * FROM events WHERE user_id = ? ORDER BY date_time DESC`

	dbRows, err := db.DB.Query(query, userId)
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

func (u *User) UpdatePassword() error {
	newHashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	query := `
	UPDATE users
	SET password = ?
	WHERE email = ?
	`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()
	_, err = statement.Exec(newHashedPassword, u.Email)

	return err
}
