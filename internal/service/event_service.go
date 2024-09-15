package service

import (
	"errors"
	"time"

	"github.com/GonzaloC17/event-management-api/internal/model"
)

var events []model.Event
var idCounter = 1

func GetActiveEvents() []model.Event {
	var activeEvents []model.Event
	for _, event := range events {
		if event.Status == model.Published && event.DateTime.After(time.Now()) {
			activeEvents = append(activeEvents, event)
		}
	}
	return activeEvents
}

func GetCompletedEvents() []model.Event {
	var completedEvents []model.Event
	for _, event := range events {
		if event.Status == model.Published && event.DateTime.Before(time.Now()) {
			completedEvents = append(completedEvents, event)
		}
	}
	return completedEvents
}

func GetAllEvents() []model.Event {
	return events
}

func CreateEvent(event model.Event) error {
	if event.Title == "" {
		return errors.New("Title cannot be empty")
	}
	if event.DateTime.Before(time.Now()) {
		return errors.New("Event date must be in the future")
	}
	event.ID = idCounter
	idCounter++
	events = append(events, event)
	return nil
}

func GetEventByID(id int) (model.Event, error) {
	for _, event := range events {
		if event.ID == id {
			return event, nil
		}
	}
	return model.Event{}, errors.New("Event not found")
}

func UpdateEvent(updatedEvent model.Event) (model.Event, error) {
	for i, event := range events {
		if event.ID == updatedEvent.ID {
			events[i] = updatedEvent
			return updatedEvent, nil
		}
	}
	return model.Event{}, errors.New("Event not found")
}

func DeleteEvent(id int) error {
	for i, event := range events {
		if event.ID == id {
			events = append(events[:i], events[i+1:]...)
			return nil
		}
	}
	return errors.New("Event not found")
}
