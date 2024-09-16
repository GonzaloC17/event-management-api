package service

import (
	"time"

	"github.com/GonzaloC17/event-management-api/internal/model"
	"github.com/GonzaloC17/event-management-api/internal/repository"
)

type EventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) GetActiveEvents() []model.Event {
	var activeEvents []model.Event
	for _, event := range s.repo.GetAll() {
		if event.Status == model.Published && event.DateTime.After(time.Now()) {
			activeEvents = append(activeEvents, event)
		}
	}
	return activeEvents
}

func (s *EventService) GetCompletedEvents() []model.Event {
	var completedEvents []model.Event
	for _, event := range s.repo.GetAll() {
		if event.Status == model.Published && event.DateTime.Before(time.Now()) {
			completedEvents = append(completedEvents, event)
		}
	}
	return completedEvents
}

func (s *EventService) GetSubscribedEvents(userEmail string) []model.Event {
	var subscribedEvents []model.Event
	for _, event := range s.repo.GetAll() {
		for _, subscriber := range event.Subscribers {
			if subscriber == userEmail {
				subscribedEvents = append(subscribedEvents, event)
				break
			}
		}
	}
	return subscribedEvents
}

func (s *EventService) GetAllEvents() []model.Event {
	return s.repo.GetAll()
}

func (s *EventService) CreateEvent(event model.Event) error {
	return s.repo.Create(event)
}

func (s *EventService) GetEventByID(id int) (model.Event, error) {
	return s.repo.GetByID(id)
}

func (s *EventService) UpdateEvent(updatedEvent model.Event) (model.Event, error) {
	return s.repo.Update(updatedEvent)
}

func (s *EventService) DeleteEvent(id int) error {
	return s.repo.Delete(id)
}
