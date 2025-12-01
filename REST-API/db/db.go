package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver - underscore is a blank to force inclusion
)

var DB *sql.DB // global variable to hold the database connection pool

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // driver name, database path & file name
	if err != nil {
		panic("Could not connect to database") // crashes the app if there is no DB
	}
	DB.SetMaxOpenConns(10) // configure connection pooling - how many simultaneously (auto managed)
	DB.SetMaxIdleConns(5)  // how many connections stay open if no one is using it

	createTables()
}

func createTables() {

	// Create USERS table
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable) // execute SQL
	if err != nil {
		fmt.Println("error", err)
		panic("Could not create users table")
	}

	// Create EVENTS table
	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		userid INTEGER,
		FOREIGN KEY(userid) REFERENCES users(id)
		)
		`
	//FOREIGN KEY (userid) REFERENCES users(id)

	_, err = DB.Exec(createEventsTable) // execute the SQL statement
	if err != nil {
		fmt.Println("error", err)
		panic("Could not create events table")
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  eventid INTEGER,
	  userid INTEGER,
	  FOREIGN KEY(eventid) REFERENCES events(id),
	  FOREIGN KEY(userid) REFERENCES users(id)
	  )
	  `
	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		fmt.Println("error", err)
		panic("Cound not create registrations table")
	}
}
