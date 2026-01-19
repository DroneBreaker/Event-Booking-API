package routes

import (
	"github.com/DroneBreaker/Event-Booking-API.git/handlers"
	"github.com/gin-gonic/gin"
)

func EventRoutes(server *gin.Engine) {
	server.GET("/events", handlers.GetEvents)
	server.POST("/events", handlers.CreateEvent)
	server.GET("/events/:id", handlers.GetEventByID)
	server.PUT("/events/:id", handlers.UpdateEvent)
	server.DELETE("/events/:id", handlers.DeleteEvent)
}
