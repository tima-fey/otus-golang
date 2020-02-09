package calendar

import (
	"time"

	events "github.com/tima-fey/otus-golang/calendar/internal/pkg/types"
)

type Calendar struct {
	Events Storage
	User   string
	UserID int
}
type Storage interface {
	AddEvent(e events.Event) (int, error)
	RemoveEvent(eventID int) error
	ReplaceEvent(eventID int, e events.Event) error
	EventList() map[int]events.Event
	IsFree(start, end time.Time) bool
	GetEvent(eventID int) (events.Event, error)
}
