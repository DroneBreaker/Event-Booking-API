package models

import (
	"time"

	"github.com/DroneBreaker/Event-Booking-API.git/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	// q, err := db.DB.Prepare(query)
	// if err != nil {
	// 	return err
	// }

	// result, err := q.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	// if err != nil {
	// 	return err
	// }
	result, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id

	return nil
}

func GetEvents() []Event {
	return events
}
