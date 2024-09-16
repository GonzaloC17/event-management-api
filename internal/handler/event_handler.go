package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/GonzaloC17/event-management-api/internal/model"
	"github.com/GonzaloC17/event-management-api/internal/repository"
	"github.com/GonzaloC17/event-management-api/internal/service"
	"github.com/GonzaloC17/event-management-api/internal/utils"
	"github.com/gin-gonic/gin"
)

var eventRepo = repository.NewInMemoryEventRepository()
var eventService = service.NewEventService(eventRepo)

func SubscribeToEvent(c *gin.Context) {
	eventID, err := strconv.Atoi(c.Param("eventID"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid event ID")
		return
	}
	event, err := eventService.GetEventByID(eventID)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "Event ID not found")
		return
	}

	if event.Status != model.Published || event.DateTime.Before(time.Now()) {
		utils.SendError(c, http.StatusForbidden, "Cannot subscribe to this event")
		return
	}

	// Simulación
	userEmail := c.GetHeader("email")
	event.Subscribers = append(event.Subscribers, userEmail)
	_, err = eventService.UpdateEvent(event)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to update event")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully subscribed to the event"})
}

func GetEvents(c *gin.Context) {
	userRole := c.GetHeader("role") // Simulación
	titleFilter := c.Query("title")
	statusFilter := c.Query("status")
	dateFilter := c.Query("date")

	events := eventService.GetAllEvents()
	var filteredEvents []model.Event
	for _, event := range events {
		if event.Status == model.Draft && userRole != "admin" {
			continue
		}

		if titleFilter != "" && !utils.ContainsIgnoreCase(event.Title, titleFilter) {
			continue
		}

		if statusFilter != "" && !utils.MatchesStatus(event.Status, statusFilter) {
			continue
		}

		if dateFilter != "" && !utils.MatchesDate(event.DateTime, dateFilter) {
			continue
		}

		filteredEvents = append(filteredEvents, event)
	}
	c.JSON(http.StatusOK, filteredEvents)
}

func GetActiveEvents(c *gin.Context) {
	events := eventService.GetActiveEvents()
	c.JSON(http.StatusOK, events)
}

func GetCompletedEvents(c *gin.Context) {
	events := eventService.GetCompletedEvents()
	c.JSON(http.StatusOK, events)
}

func GetSubscribedEvents(c *gin.Context) {
	userEmail := c.GetHeader("email")
	if userEmail == "" {
		utils.SendError(c, http.StatusBadRequest, "User email is required")
		return
	}

	events := eventService.GetSubscribedEvents(userEmail)
	c.JSON(http.StatusOK, events)
}

func CreateEvent(c *gin.Context) {
	var newEvent model.Event
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if newEvent.Title == "" {
		utils.SendError(c, http.StatusBadRequest, "Title is required")
		return
	}

	err := eventService.CreateEvent(newEvent)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create event")
		return
	}

	c.JSON(http.StatusCreated, newEvent)
}

func UpdateEvent(c *gin.Context) {
	userRole := c.GetHeader("role") // Simulación
	if userRole != "admin" {
		utils.SendError(c, http.StatusForbidden, "Only admins can update events")
		return
	}

	eventID, err := strconv.Atoi(c.Param("eventID"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid event ID")
		return
	}

	var updatedEvent model.Event
	if err := c.ShouldBindJSON(&updatedEvent); err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if updatedEvent.DateTime.Before(time.Now()) {
		utils.SendError(c, http.StatusBadRequest, "Event date must be in the future")
		return
	}

	updatedEvent.ID = eventID
	updatedEvent, err = eventService.UpdateEvent(updatedEvent)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "Event not found")
		return
	}

	c.JSON(http.StatusOK, updatedEvent)
}

func DeleteEvent(c *gin.Context) {
	userRole := c.GetHeader("role") // Simulación
	if userRole != "admin" {
		utils.SendError(c, http.StatusForbidden, "Only admins can delete events")
		return
	}

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid event ID")
		return
	}

	err = eventService.DeleteEvent(eventID)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "Event not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
