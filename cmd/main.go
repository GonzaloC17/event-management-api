package main

import (
	"github.com/GonzaloC17/event-management-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eventRoutes := r.Group("/events")
	{
		eventRoutes.GET("/", handler.GetEvents)

		eventRoutes.POST("/", handler.CreateEvent)

		eventRoutes.POST("/:eventID/subscribe", handler.SubscribeToEvent)

		eventRoutes.GET("/active", handler.GetActiveEvents)

		eventRoutes.GET("/completed", handler.GetCompletedEvents)

		eventRoutes.PUT("/:eventID", handler.UpdateEvent)

		eventRoutes.DELETE("/:eventID", handler.DeleteEvent)
	}

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", handler.CreateUser)

		userRoutes.GET("/", handler.GetAllUsers)

		userRoutes.GET("/:userID", handler.GetUserByID)
	}
	r.Run(":8080")
}
