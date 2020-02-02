package events

import "time"

type ErrDateBusy struct {
}

func (e *ErrDateBusy) Error() string {
	return "Date is busy"
}

type ErrNotSuchID struct {
}

func (e *ErrNotSuchID) Error() string {
	return "There is no such id"
}

type Storage interface {
	AddEvent(e Event) error
	RemoveEvent(eventID int) error
	ReplaceEvent(eventID int, e Event) error
	EventList() []Event
	IsFree(time.Time) bool
}
type Event struct {
	name      string
	startDate time.Time
	endDate   time.Time
}
