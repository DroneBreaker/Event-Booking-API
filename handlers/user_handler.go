package handlers

import (
	"net/http"

	"github.com/DroneBreaker/Event-Booking-API.git/models"
	"github.com/DroneBreaker/Event-Booking-API.git/repository"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})

		return
	}

	err = repository.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create user",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    user,
	})

}
