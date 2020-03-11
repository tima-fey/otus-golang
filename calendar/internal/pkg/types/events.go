package events

import "time"

type Event struct {
	Name      string
	EventID   int32
	StartDate time.Time
	EndDate   time.Time
}
