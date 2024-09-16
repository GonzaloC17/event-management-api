package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/GonzaloC17/event-management-api/internal/domain"
	"github.com/GonzaloC17/event-management-api/internal/usecase"
	"github.com/GonzaloC17/event-management-api/internal/utils"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventService *usecase.EventService
}

func NewEventHandler(eventService *usecase.EventService) *EventHandler {
	return &EventHandler{eventService: eventService}
}

func (h *EventHandler) SubscribeToEvent(c *gin.Context) {
	eventID, err := strconv.Atoi(c.Param("eventID"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid event ID")
		return
	}
	event, err := h.eventService.GetEventByID(eventID)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "Event ID not found")
		return
	}

	if event.Status != domain.Published || event.DateTime.Before(time.Now()) {
		utils.SendError(c, http.StatusForbidden, "Cannot subscribe to this event")
		return
	}

	// Simulación de autenticación de usuario
	userEmail := c.GetHeader("email")
	event.Subscribers = append(event.Subscribers, userEmail)
	_, err = h.eventService.UpdateEvent(event)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to update event")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully subscribed to the event"})
}

func (h *EventHandler) GetEvents(c *gin.Context) {
	userRole := c.GetHeader("role") // Simulación de roles
	titleFilter := c.Query("title")
	statusFilter := c.Query("status")
	dateFilter := c.Query("date")

	events := h.eventService.GetAllEventsFiltered(userRole, titleFilter, statusFilter, dateFilter)
	c.JSON(http.StatusOK, events)
}

func (h *EventHandler) GetActiveEvents(c *gin.Context) {
	events := h.eventService.GetActiveEvents()
	c.JSON(http.StatusOK, events)
}

func (h *EventHandler) GetCompletedEvents(c *gin.Context) {
	events := h.eventService.GetCompletedEvents()
	c.JSON(http.StatusOK, events)
}

func (h *EventHandler) GetSubscribedEvents(c *gin.Context) {
	userEmail := c.GetHeader("email") // Simulación de autenticación de usuario
	if userEmail == "" {
		utils.SendError(c, http.StatusBadRequest, "User email is required")
		return
	}

	events := h.eventService.GetSubscribedEvents(userEmail)
	c.JSON(http.StatusOK, events)
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var newEvent domain.Event
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if newEvent.Title == "" {
		utils.SendError(c, http.StatusBadRequest, "Title is required")
		return
	}

	err := h.eventService.CreateEvent(newEvent)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create event")
		return
	}

	c.JSON(http.StatusCreated, newEvent)
}

func (h *EventHandler) UpdateEvent(c *gin.Context) {
	userRole := c.GetHeader("role") // Simulación de roles
	if userRole != "admin" {
		utils.SendError(c, http.StatusForbidden, "Only admins can update events")
		return
	}

	eventID, err := strconv.Atoi(c.Param("eventID"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid event ID")
		return
	}

	var updatedEvent domain.Event
	if err := c.ShouldBindJSON(&updatedEvent); err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if updatedEvent.DateTime.Before(time.Now()) {
		utils.SendError(c, http.StatusBadRequest, "Event date must be in the future")
		return
	}

	updatedEvent.ID = eventID
	updatedEvent, err = h.eventService.UpdateEvent(updatedEvent)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "Event not found")
		return
	}

	c.JSON(http.StatusOK, updatedEvent)
}

func (h *EventHandler) DeleteEvent(c *gin.Context) {
	userRole := c.GetHeader("role") // Simulación de roles
	if userRole != "admin" {
		utils.SendError(c, http.StatusForbidden, "Only admins can delete events")
		return
	}

	eventID, err := strconv.Atoi(c.Param("eventID"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid event ID")
		return
	}

	err = h.eventService.DeleteEvent(eventID)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "Event not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
