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

var events []Event

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
	// insert, update and delete use Exec
	result, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	// defer db.DB.Close()

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id

	return err
}

func GetEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		var dateTimeStr string

		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserID)
		if err != nil {
			return nil, err
		}

		// Parse the datetime string into time.Time
		// Parse the datetime string - try different formats
		layouts := []string{
			"2006-01-02 15:04:05-07:00",
			"2006-01-02 15:04:05",
			time.RFC3339,
		}

		var parseErr error
		for _, layout := range layouts {
			event.DateTime, parseErr = time.Parse(layout, dateTimeStr)
			if parseErr == nil {
				break
			}
		}
		if parseErr != nil {
			return nil, parseErr
		}

		events = append(events, event)
	}

	return events, nil
}
