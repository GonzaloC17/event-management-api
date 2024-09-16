package repository

import (
	"errors"
	"time"

	"github.com/GonzaloC17/event-management-api/internal/domain"
)

type InMemoryEventRepository struct {
	events    []domain.Event
	idCounter int
}

func NewInMemoryEventRepository() *InMemoryEventRepository {
	return &InMemoryEventRepository{}
}

func (r *InMemoryEventRepository) Create(event domain.Event) error {
	if event.Title == "" {
		return errors.New("title cannot be empty")
	}
	if event.DateTime.Before(time.Now()) {
		return errors.New("event date must be in the future")
	}
	event.ID = r.idCounter
	r.idCounter++
	r.events = append(r.events, event)
	return nil
}

func (r *InMemoryEventRepository) GetByID(id int) (domain.Event, error) {
	for _, event := range r.events {
		if event.ID == id {
			return event, nil
		}
	}
	return domain.Event{}, errors.New("event not found")
}

func (r *InMemoryEventRepository) Update(updatedEvent domain.Event) (domain.Event, error) {
	for i, event := range r.events {
		if event.ID == updatedEvent.ID {
			r.events[i] = updatedEvent
			return updatedEvent, nil
		}
	}
	return domain.Event{}, errors.New("event not found")
}

func (r *InMemoryEventRepository) Delete(id int) error {
	for i, event := range r.events {
		if event.ID == id {
			r.events = append(r.events[:i], r.events[i+1:]...)
			return nil
		}
	}
	return errors.New("event not found")
}

func (r *InMemoryEventRepository) GetAll() []domain.Event {
	return r.events
}
