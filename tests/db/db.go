package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DB, error := sql.Open("sqlite3", "api.db")

	if error != nil {
		panic("couldnt connect to DB")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createEventsTable()
}

func createEventsTable() {
	createEventsTableQuery := `
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name STRING NOT NULL,
	description STRING NOT NULL,
	date_time DATETIME NOT NULL,
	user_id INTEGER
	)`

	_, error := DB.Exec(createEventsTableQuery)

	if error != nil {
		panic("error creating events table")
	}
}
