package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/GonzaloC17/event-management-api/internal/domain"
	"github.com/GonzaloC17/event-management-api/internal/handler"
	"github.com/GonzaloC17/event-management-api/internal/infrastructure/repository"
	"github.com/GonzaloC17/event-management-api/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupEventRouter() *gin.Engine {
	repo := repository.NewInMemoryEventRepository()
	service := usecase.NewEventService(repo)
	handler := handler.NewEventHandler(service)

	router := gin.Default()
	router.POST("/events/:eventID/subscribe", handler.SubscribeToEvent)
	router.GET("/events", handler.GetEvents)
	router.GET("/events/active", handler.GetActiveEvents)
	router.GET("/events/completed", handler.GetCompletedEvents)
	router.GET("/events/subscribed", handler.GetSubscribedEvents)
	router.POST("/events", handler.CreateEvent)
	router.PUT("/events/:eventID", handler.UpdateEvent)
	router.DELETE("/events/:eventID", handler.DeleteEvent)

	return router
}

func TestSubscribeToEvent(t *testing.T) {
	router := setupEventRouter()

	event := domain.Event{
		Title:     "Event to Subscribe",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	repo := repository.NewInMemoryEventRepository()
	repo.Create(event)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/events/1/subscribe", nil)
	req.Header.Set("email", "user@example.com")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Successfully subscribed to the event")
}

func TestGetEvents(t *testing.T) {
	router := setupEventRouter()

	repo := repository.NewInMemoryEventRepository()
	event := domain.Event{
		Title:     "Event 10",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	repo.Create(event)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/events", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Event 10")
}

func TestCreateEvent(t *testing.T) {
	router := setupEventRouter()

	eventPayload := `{
		"title": "New Event",
		"short_desc": "Short description",
		"long_desc": "Long description",
		"date_time": "` + time.Now().Add(time.Hour).Format(time.RFC3339) + `",
		"organizer": "Organizer",
		"location": "Location",
		"status": "Published"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/events", bytes.NewBufferString(eventPayload))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "New Event")
}

func TestUpdateEvent(t *testing.T) {
	router := setupEventRouter()

	// Create an event
	event := domain.Event{
		Title:     "Event 11",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	repo := repository.NewInMemoryEventRepository()
	repo.Create(event)

	updatedEventPayload := `{
		"title": "Updated Event",
		"short_desc": "Updated description",
		"long_desc": "Updated long description",
		"date_time": "` + time.Now().Add(time.Hour*2).Format(time.RFC3339) + `",
		"organizer": "Updated Organizer",
		"location": "Updated Location",
		"status": "Published"
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/events/0", bytes.NewBufferString(updatedEventPayload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("role", "admin")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated Event")
}

func TestDeleteEvent(t *testing.T) {
	router := setupEventRouter()

	event := domain.Event{
		Title:     "Event 12",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	repo := repository.NewInMemoryEventRepository()
	repo.Create(event)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/events/0", nil)
	req.Header.Set("role", "admin")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Event deleted successfully")
}
