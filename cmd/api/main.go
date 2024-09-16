package main

import (
	"github.com/GonzaloC17/event-management-api/internal/handler"
	"github.com/GonzaloC17/event-management-api/internal/infrastructure/repository"
	"github.com/GonzaloC17/event-management-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	eventRepo := repository.NewInMemoryEventRepository()
	userRepo := repository.NewInMemoryUserRepository()

	eventService := usecase.NewEventService(eventRepo)
	userService := usecase.NewUserService(userRepo)

	eventHandler := handler.NewEventHandler(eventService)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	eventRoutes := r.Group("/events")
	{
		eventRoutes.GET("/", eventHandler.GetEvents)

		eventRoutes.POST("/", eventHandler.CreateEvent)

		eventRoutes.POST("/:eventID/subscribe", eventHandler.SubscribeToEvent)

		eventRoutes.GET("/active", eventHandler.GetActiveEvents)

		eventRoutes.GET("/completed", eventHandler.GetCompletedEvents)

		eventRoutes.GET("/subscribed", eventHandler.GetSubscribedEvents)

		eventRoutes.PUT("/:eventID", eventHandler.UpdateEvent)

		eventRoutes.DELETE("/:eventID", eventHandler.DeleteEvent)
	}

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)

		userRoutes.GET("/", userHandler.GetAllUsers)

		userRoutes.GET("/:userID", userHandler.GetUserByID)
	}
	r.Run(":8080")
}
