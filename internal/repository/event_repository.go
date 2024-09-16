package repository

import (
	"github.com/GonzaloC17/event-management-api/internal/model"
)

type EventRepository interface {
	Create(event model.Event) error
	GetByID(id int) (model.Event, error)
	Update(updatedEvent model.Event) (model.Event, error)
	Delete(id int) error
	GetAll() []model.Event
}
