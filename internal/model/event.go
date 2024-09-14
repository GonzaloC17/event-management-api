package model

import "time"

type EventStatus string

const (
	Draft     EventStatus = "draft"
	Published EventStatus = "published"
)

type Event struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	ShortDesc   string      `json:"short_desc"` // Etiqueta JSON correcta
	LongDesc    string      `json:"long_desc"`  // Etiqueta JSON correcta
	DateTime    time.Time   `json:"date_time"`
	Organizer   string      `json:"organizer"`
	Location    string      `json:"location"`
	Status      EventStatus `json:"status"`
	Subscribers []string    `json:"subscribers"`
}
