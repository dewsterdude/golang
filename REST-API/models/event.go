package models

import (
	"fmt"
	"time"

	"example.com/REST-API/db"
)

type Event struct {
	ID          int64
	Name        string    `binding: "required"`
	Description string    `binding: "required"`
	Location    string    `binding: "required"`
	DateTime    time.Time `binding: "required"`
	UserID      int64
}

// var events = []Event{} // slice initialized to empty

func (e *Event) Save() error {
	// later add to a DB
	// 	events = append(events, &e)

	query := `INSERT INTO events (Name, Description, Location, DateTime, Userid) 
			  VALUES (?, ?, ?, ?, ?)` // ? placeholders for parameters to avoid SQL injection
	stmt, err := db.DB.Prepare(query) // prepare the SQL statement
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID) // execute the query with the event data
	if err != nil {
		return err
	}
	id, err := result.LastInsertId() // get the ID of the newly inserted row if needed
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query) // typically used if you want a bunch of rows, vs exec for insert/update/delete
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() { // for as long as there are rows to read
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID) // scan the columns into the struct fields
		if err != nil {
			return nil, err // nil for events slice and the error
		}
		events = append(events, event) // add to the slice
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id) // query for a single row

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e Event) Update() error {
	query := `UPDATE events
			  SET name = ?, description = ?, location = ?, datetime = ?, userid = ?
			  WHERE id = ?
			  `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("update() prepare query failed:", err)
		return err
	}
	fmt.Println("update() query:", stmt)

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)

	fmt.Println("update() failed stmt.exec:", stmt)

	return err // nil if everything worked or actual error if something was wrong
}

func (event Event) Delete() error {
	query := "DELETE FROM events where id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registration(eventid, userid) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registration WHERE eventid = ? AND userid = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	return err
}
