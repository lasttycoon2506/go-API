package models

import "example.com/m/v2/db"

type User struct {
	ID       int64
	email    string `binding:"required"`
	password string `binding:"required"`
}

func (u User) Save() error {
	query := `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(u.email, u.password)
	return err
}
