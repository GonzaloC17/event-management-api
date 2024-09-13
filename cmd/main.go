package main

import (
	"github.com/GonzaloC17/even-management-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/events", handler.GetEvents)
	router.POST("/events", handler.CreateEvents)
	router.Run(":8080")
}
