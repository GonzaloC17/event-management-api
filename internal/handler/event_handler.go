package handler

import (
	"net/http"
	"time"

	"github.com/GonzaloC17/even-management-api/internal/model"
	"github.com/GonzaloC17/even-management-api/internal/service"
	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	events := service.GetAllEvents()
	c.JSON(http.StatusOK, events)
}

func CreateEvents(c *gin.Context) {
	var event model.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if event.DateTime.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event date must be in the future"})
		return
	}

	newEvent := service.CreateEvent(event)
	c.JSON(http.StatusCreated, newEvent)
}
