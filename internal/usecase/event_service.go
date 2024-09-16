package usecase

import (
	"strings"
	"time"

	"github.com/GonzaloC17/event-management-api/internal/domain"
	"github.com/GonzaloC17/event-management-api/internal/utils"
)

type EventService struct {
	repo domain.EventRepository
}

func NewEventService(repo domain.EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) GetActiveEvents() []domain.Event {
	var activeEvents []domain.Event
	for _, event := range s.repo.GetAll() {
		if event.Status == domain.Published && event.DateTime.After(time.Now()) {
			activeEvents = append(activeEvents, event)
		}
	}
	return activeEvents
}

func (s *EventService) GetCompletedEvents() []domain.Event {
	var completedEvents []domain.Event
	for _, event := range s.repo.GetAll() {
		if event.Status == domain.Published && event.DateTime.Before(time.Now()) {
			completedEvents = append(completedEvents, event)
		}
	}
	return completedEvents
}

func (s *EventService) GetSubscribedEvents(userEmail string) []domain.Event {
	var subscribedEvents []domain.Event
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

func (s *EventService) GetAllEventsFiltered(userRole, titleFilter, statusFilter, dateFilter string) []domain.Event {
	events := s.repo.GetAll()
	var filteredEvents []domain.Event

	for _, event := range events {
		if event.Status == domain.Draft && userRole != "admin" {
			continue
		}

		if titleFilter != "" && !strings.Contains(strings.ToLower(event.Title), strings.ToLower(titleFilter)) {
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
	return filteredEvents
}

func (s *EventService) CreateEvent(event domain.Event) error {
	return s.repo.Create(event)
}

func (s *EventService) GetEventByID(id int) (domain.Event, error) {
	return s.repo.GetByID(id)
}

func (s *EventService) UpdateEvent(updatedEvent domain.Event) (domain.Event, error) {
	return s.repo.Update(updatedEvent)
}

func (s *EventService) DeleteEvent(id int) error {
	return s.repo.Delete(id)
}
