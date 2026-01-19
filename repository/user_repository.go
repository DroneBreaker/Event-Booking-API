package repository

import (
	"github.com/DroneBreaker/Event-Booking-API.git/db"
	"github.com/DroneBreaker/Event-Booking-API.git/models"
	"github.com/DroneBreaker/Event-Booking-API.git/utils"
)

func CreateUser(u *models.User) error {
	query := `
	INSERT INTO users (email, password) 
	VALUES (?, ?)`

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	// u.Password = hashedPassword

	result, err := db.DB.Exec(query, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id

	return err
}
