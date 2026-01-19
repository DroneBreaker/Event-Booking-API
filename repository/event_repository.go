package repository

import (
	"time"

	"github.com/DroneBreaker/Event-Booking-API.git/db"
	"github.com/DroneBreaker/Event-Booking-API.git/models"
)

func Save(e *models.Event) error {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	result, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id

	return err
}

func GetEvents() ([]models.Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
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

func GetEventByID(id int64) (*models.Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event models.Event
	var dateTimeStr string

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserID)
	if err != nil {
		return nil, err
	}

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

	return &event, nil
}

func Update(e models.Event) error {
	query := `
	UPDATE events 
	SET name = ?, description = ?, location = ?, datetime = ?, user_id = ?
	WHERE id = ?`

	_, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func Delete(id int64) error {
	query := `DELETE FROM events WHERE id = ?`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
