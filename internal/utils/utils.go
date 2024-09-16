package utils

import (
	"strings"
	"time"

	"github.com/GonzaloC17/event-management-api/internal/domain"
)

func ContainsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

func MatchesStatus(eventStatus domain.EventStatus, filterStatus string) bool {
	return strings.EqualFold(string(eventStatus), filterStatus)
}

func MatchesDate(eventDate time.Time, filterDate string) bool {
	parsedDate, err := time.Parse("2006-01-02", filterDate)
	if err != nil {
		return false
	}
	return eventDate.Format("2006-01-02") == parsedDate.Format("2006-01-02")
}
