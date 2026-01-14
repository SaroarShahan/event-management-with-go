package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "events.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT NOT NULL UNIQUE,
		"email" TEXT NOT NULL UNIQUE,
		"password" TEXT NOT NULL
	);`

	if _, err := DB.Exec(createUsersTable); err != nil {
		panic("Could not create users table.")
	}

	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT NOT NULL,
		"description" TEXT NOT NULL,
		"location" TEXT NOT NULL,
		"datetime" DATETIME NOT NULL,
		"user_id" INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	if _, err := DB.Exec(createEventsTable); err != nil {
		panic("Could not create events table.")
	}

	createRegistrationsTable := `CREATE TABLE IF NOT EXISTS registrations (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"user_id" INTEGER,
		"event_id" INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(event_id) REFERENCES events(id)
	);`

	if _, err := DB.Exec(createRegistrationsTable); err != nil {
		panic("Could not create registrations table.")
	}
}