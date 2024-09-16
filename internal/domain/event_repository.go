package domain

type EventRepository interface {
	Create(event Event) error
	GetByID(id int) (Event, error)
	Update(updatedEvent Event) (Event, error)
	Delete(id int) error
	GetAll() []Event
}
