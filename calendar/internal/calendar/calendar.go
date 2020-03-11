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
	AddEvent(e events.Event) (int32, error)
	RemoveEvent(eventID int32) error
	ReplaceEvent(eventID int32, e events.Event) error
	EventList() map[int32]events.Event
	IsFree(start, end time.Time) bool
	GetEvent(eventID int32) (events.Event, error)
}
