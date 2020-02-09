package events

import "time"

type Event struct {
	Name      string
	EventID   int
	StartDate time.Time
	EndDate   time.Time
}
