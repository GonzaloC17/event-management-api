package repository_test

import (
	"testing"
	"time"

	"github.com/GonzaloC17/event-management-api/internal/domain"
	"github.com/GonzaloC17/event-management-api/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
	repo := repository.NewInMemoryEventRepository()

	event := domain.Event{
		Title:     "Event 1",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour), // Future date
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	err := repo.Create(event)
	assert.NoError(t, err)

	// Test creacion de evento sin titulo
	invalidEvent := domain.Event{
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	err = repo.Create(invalidEvent)
	assert.Error(t, err)
	assert.Equal(t, "title cannot be empty", err.Error())

	// Test creacion de evento con fecha pasada
	pastEvent := domain.Event{
		Title:     "Past Event",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(-time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	err = repo.Create(pastEvent)
	assert.Error(t, err)
	assert.Equal(t, "event date must be in the future", err.Error())
}

func TestGetEventByID(t *testing.T) {
	repo := repository.NewInMemoryEventRepository()

	event := domain.Event{
		Title:     "Event 2",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	repo.Create(event)

	retrievedEvent, err := repo.GetByID(0)
	assert.NoError(t, err)
	assert.Equal(t, event.Title, retrievedEvent.Title)

	_, err = repo.GetByID(1)
	assert.Error(t, err)
	assert.Equal(t, "event not found", err.Error())
}

func TestUpdateEvent(t *testing.T) {
	repo := repository.NewInMemoryEventRepository()

	event := domain.Event{
		Title:     "Event 3",
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
		Title:     "Updated Event",
		ShortDesc: "Updated description",
		LongDesc:  "Updated long description",
		DateTime:  time.Now().Add(time.Hour * 2),
		Organizer: "Updated Organizer",
		Location:  "Updated Location",
		Status:    domain.Published,
	}
	event, err := repo.Update(updatedEvent)
	assert.NoError(t, err)
	assert.Equal(t, updatedEvent.Title, event.Title)

	invalidEvent := domain.Event{
		ID:        1,
		Title:     "Non-existent Event",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	_, err = repo.Update(invalidEvent)
	assert.Error(t, err)
	assert.Equal(t, "event not found", err.Error())
}

func TestDeleteEvent(t *testing.T) {
	repo := repository.NewInMemoryEventRepository()

	event := domain.Event{
		Title:     "Event 4",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	repo.Create(event)

	err := repo.Delete(0)
	assert.NoError(t, err)

	_, err = repo.GetByID(0)
	assert.Error(t, err)
	assert.Equal(t, "event not found", err.Error())
}

func TestGetAllEvents(t *testing.T) {
	repo := repository.NewInMemoryEventRepository()

	event1 := domain.Event{
		Title:     "Event 5",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	repo.Create(event1)
	event2 := domain.Event{
		Title:     "Event 6",
		ShortDesc: "Short description",
		LongDesc:  "Long description",
		DateTime:  time.Now().Add(time.Hour),
		Organizer: "Organizer",
		Location:  "Location",
		Status:    domain.Published,
	}
	repo.Create(event2)

	events := repo.GetAll()
	assert.Len(t, events, 2)
	assert.Contains(t, events, event1)
	assert.Contains(t, events, event2)
}
