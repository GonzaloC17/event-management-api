package utils

import (
	"strings"
	"time"
)

func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

func matchesStatus(eventStatus, filterStatus string) bool {
	return strings.ToLower(eventStatus) == strings.ToLower(filterStatus)
}

func matchesDate(eventDate time.Time, filterDate string) bool {
	parsedDate, err := time.Parse("2006-01-02", filterDate)
	if err != nil {
		return false
	}
	return eventDate.Format("2006-01-02") == parsedDate.Format("2006-01-02")
}
