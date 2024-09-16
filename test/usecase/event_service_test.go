package usecase_test

import (
	"testing"
	"time"

	"github.com/GonzaloC17/event-management-api/internal/domain"
	"github.com/GonzaloC17/event-management-api/internal/infrastructure/repository"
	"github.com/GonzaloC17/event-management-api/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetActiveEvents(t *testing.T) {
	repo := repository.NewInMemoryEventRepository()
	service := usecase.NewEventService(repo)

	event1 := domain.Event{
		Title:     "Active Event 1",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	event2 := domain.Event{
		Title:     "Inactive Event 1",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(-time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	repo.Create(event1)
	repo.Create(event2)

	activeEvents := service.GetActiveEvents()
	assert.Len(t, activeEvents, 1)
	assert.Contains(t, activeEvents, event1)
}

func TestGetCompletedEvents(t *testing.T) {
	repo := repository.NewInMemoryEventRepository()
	service := usecase.NewEventService(repo)

	event1 := domain.Event{
		Title:     "Completed Event 1",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(-time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	repo.Create(event1)

	event2 := domain.Event{
		Title:     "Active Event 2",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	repo.Create(event2)

	completedEvents := service.GetCompletedEvents()
	assert.Len(t, completedEvents, 1)
	assert.Contains(t, completedEvents, event1)
}

func TestGetSubscribedEvents(t *testing.T) {
	repo := repository.NewInMemoryEventRepository()
	service := usecase.NewEventService(repo)

	event := domain.Event{
		Title:       "Subscribed Event",
		ShortDesc:   "Short description",
		LongDesc:    "Long description",
		DateTime:    time.Now().Add(time.Hour),
		Organizer:   "Organizer",
		Location:    "Location",
		Status:      domain.Published,
		Subscribers: []string{"user1@example.com"},
	}
	repo.Create(event)

	subscribedEvents := service.GetSubscribedEvents("user1@example.com")
	assert.Len(t, subscribedEvents, 1)
	assert.Contains(t, subscribedEvents, event)

	unsubscribedEvents := service.GetSubscribedEvents("user2@example.com")
	assert.Len(t, unsubscribedEvents, 0)
}

func TestGetAllEventsFiltered(t *testing.T) {
	repo := repository.NewInMemoryEventRepository()
	service := usecase.NewEventService(repo)

	event1 := domain.Event{
		Title:     "Event with Filter",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	event2 := domain.Event{
		Title:     "Another Event",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(-time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Draft,
	}
	repo.Create(event1)
	repo.Create(event2)

	events := service.GetAllEventsFiltered("user", "Filter", "", "")
	assert.Len(t, events, 1)
	assert.Contains(t, events, event1)

	events = service.GetAllEventsFiltered("admin", "", "", "")
	assert.Len(t, events, 2)
}

func TestCreateEvent(t *testing.T) {
	repo := repository.NewInMemoryEventRepository()
	service := usecase.NewEventService(repo)

	event := domain.Event{
		Title:     "Event 7",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	err := service.CreateEvent(event)
	assert.NoError(t, err)

	createdEvent, err := repo.GetByID(0)
	assert.NoError(t, err)
	assert.Equal(t, event.Title, createdEvent.Title)
}

func TestUpdateEvent(t *testing.T) {
	repo := repository.NewInMemoryEventRepository()
	service := usecase.NewEventService(repo)

	event := domain.Event{
		Title:     "Event 8",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	repo.Create(event)

	updatedEvent := domain.Event{
		ID:        0,
		Title:     "Updated Event 8",
		ShortDesc: "Updated description",
		LongDesc:  "Updated long description",
		DateTime:  time.Now().Add(time.Hour * 2),
		Organizer: "Updated Organizer",
		Location:  "Updated Location",
		Status:    domain.Published,
	}
	_, err := service.UpdateEvent(updatedEvent)
	assert.NoError(t, err)

	event, err = repo.GetByID(0)
	assert.NoError(t, err)
	assert.Equal(t, updatedEvent.Title, event.Title)
}

func TestDeleteEvent(t *testing.T) {
	repo := repository.NewInMemoryEventRepository()
	service := usecase.NewEventService(repo)

	event := domain.Event{
		Title:     "Event 9",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	repo.Create(event)

	err := service.DeleteEvent(0)
	assert.NoError(t, err)

	_, err = repo.GetByID(0)
	assert.Error(t, err)
	assert.Equal(t, "event not found", err.Error())
}
