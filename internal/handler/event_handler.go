package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/GonzaloC17/event-management-api/internal/model"
	"github.com/GonzaloC17/event-management-api/internal/service"
	"github.com/gin-gonic/gin"
)

func SubscribeToEvent(c *gin.Context) {
	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := service.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event ID not found"})
		return
	}

	if event.Status != model.Published || event.DateTime.Before(time.Now()) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot subscribe to this event"})
		return
	}

	//simulacion
	userEmail := c.GetHeader("email")
	event.Subscribers = append(event.Subscribers, userEmail)
	service.UpdateEvent(event)
	c.JSON(http.StatusOK, gin.H{"message": "Succesfully subscribed to the event"})
}

func GetEvents(c *gin.Context) {
	userRole := c.GetHeader("role") //simulacion
	titleFilter := c.Query("title")

	events := service.GetAllEvents()
	var filteredEvents []model.Event
	for _, event := range events {
		if event.Status == model.Draft && userRole != "admin" {
			continue
		}

		if titleFilter != "" && !utils.containsIgnoreCase(event.Title, titleFilter) {
			continue
		}
		filteredEvents = append(filteredEvents, event)
	}
	c.JSON(http.StatusOK, filteredEvents)
}

func GetActiveEvents(c *gin.Context) {
	events := service.GetActiveEvents()
	c.JSON(http.StatusOK, events)
}

func GetCompletedEvents(c *gin.Context) {
	events := service.GetCompletedEvents()
	c.JSON(http.StatusOK, events)
}

func CreateEvent(c *gin.Context) {
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

func UpdateEvent(c *gin.Context) {
	userRole := c.GetHeader("role") //simulacion
	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can update events"})
		return
	}

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var updatedEvent model.Event
	if err := c.ShouldBindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updatedEvent.DateTime.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event date must be in the future"})
		return
	}

	updatedEvent.ID = eventID
	updatedEvent, err = service.UpdateEvent(updatedEvent)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, updatedEvent)
}

func DeteleEvent(c *gin.Context) {
	userRole := c.GetHeader("role") //simulacion
	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admmins can delete events"})
		return
	}

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	err = service.DeleteEvent(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted succesfully"})
}
