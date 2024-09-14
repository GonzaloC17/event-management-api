package service

import (
	"errors"

	"github.com/GonzaloC17/even-management-api/internal/model"
)

var events []model.Event
var idCounter = 1

func GetAllEvents() []model.Event {
	return events
}

func CreateEvent(event model.Event) model.Event {
	event.ID = idCounter
	idCounter++
	events = append(events, event)
	return event
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
