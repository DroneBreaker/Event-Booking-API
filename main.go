package main

import (
	"github.com/DroneBreaker/Event-Booking-API.git/db"
	"github.com/DroneBreaker/Event-Booking-API.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db.InitDB()

	// Setup engine
	server := gin.Default()

	// Routes
	routes.EventRoutes(server)
	routes.UserRoutes(server)

	server.Run(":8080")
}
