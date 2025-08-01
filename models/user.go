package models

import (
	"example.com/m/v2/db"
	"example.com/m/v2/utils"
)

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

	hashedPassword, err := utils.HashPassword(u.password)
	if err != nil {
		return err
	}

	_, err = statement.Exec(u.email, hashedPassword)
	return err
}
