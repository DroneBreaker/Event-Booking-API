package db

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "event.db")

	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		location TEXT NOT NULL,
		dateTime TEXT NOT NULL,
		user_id INTEGER NOT NULL
	)`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create tables: " + err.Error())
	}
}
