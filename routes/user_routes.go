package routes

import (
	"github.com/DroneBreaker/Event-Booking-API.git/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(server *gin.Engine) {
	server.POST("/signup", handlers.Signup)
}
