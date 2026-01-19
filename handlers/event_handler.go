package handlers

import (
	"net/http"
	"strconv"

	"github.com/DroneBreaker/Event-Booking-API.git/models"
	"github.com/DroneBreaker/Event-Booking-API.git/repository"
	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	events, err := repository.GetEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch events",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, events)
}

func CreateEvent(c *gin.Context) {
	var event models.Event

	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})

		return
	}

	event.ID = 1
	event.UserID = 1

	err = repository.Save(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create event",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}

func GetEventByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid event ID",
			"details": err.Error(),
		})
		return
	}

	event, err := repository.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch event",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, event)
}

func UpdateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid event ID",
			"details": err.Error(),
		})
		return
	}

	_, err = repository.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch event",
			"details": err.Error(),
		})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
		return
	}

	updatedEvent.ID = id
	err = repository.Update(updatedEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update event",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"event":   updatedEvent,
	})
}

func DeleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid event ID",
			"details": err.Error(),
		})
		return
	}

	_, err = repository.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch event",
			"details": err.Error(),
		})
		return
	}

	err = repository.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete event",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}
