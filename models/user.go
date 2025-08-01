package models

import (
	"errors"
	"fmt"

	"example.com/m/v2/db"
	"example.com/m/v2/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
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

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, err = statement.Exec(u.Email, hashedPassword)
	return err
}

func (u User) Verify() error {
	query := `SELECT password FROM users WHERE email = ?`
	dbRow := db.DB.QueryRow(query, u.Email)

	var hashedPasswordInDb string
	err := dbRow.Scan(&hashedPasswordInDb)
	fmt.Println("Error:", err)
	if err != nil {
		return errors.New("invalid creds")
	}

	passwordIsValid := utils.CheckHashedPassword(u.Password, hashedPasswordInDb)
	// fmt.Println("Error:", "invalid pw")
	if !passwordIsValid {
		return errors.New("invalid creds")
	}

	return nil
}
