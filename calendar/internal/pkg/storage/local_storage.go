package local_storage

import (
	"time"

	"github.com/tima-fey/otus-golang/calendar/internal/pkg/errors"
	events "github.com/tima-fey/otus-golang/calendar/internal/pkg/types"
)

type LocalStorage struct {
	Events map[int32]events.Event
	NextID int32
}

func (les *LocalStorage) AddEvent(e events.Event) (int32, error) {
	if !les.IsFree(e.StartDate, e.EndDate) {
		return 0, &errors.ErrDateBusy{}
	}
	e.EventID = les.NextID
	les.Events[les.NextID] = e
	les.NextID++
	return e.EventID, nil
}
func (les *LocalStorage) RemoveEvent(eventID int32) error {
	if _, ok := les.Events[eventID]; ok {
		delete(les.Events, eventID)
		return nil
	}
	return &errors.ErrNotSuchID{}
}
func (les *LocalStorage) GetEvent(eventID int32) (events.Event, error) {
	var dummyEvent events.Event
	if _, ok := les.Events[eventID]; ok {
		return les.Events[eventID], nil
	}
	return dummyEvent, &errors.ErrNotSuchID{}
}
func (les *LocalStorage) ReplaceEvent(eventID int32, e events.Event) error {
	tempEvent, ok := les.Events[eventID]
	if ok {
		delete(les.Events, eventID)
	} else {
		return &errors.ErrNotSuchID{}
	}
	if les.IsFree(e.StartDate, e.EndDate) {
		les.Events[eventID] = e
		return nil
	}
	e.EventID = eventID
	les.Events[eventID] = tempEvent
	return &errors.ErrDateBusy{}
}
func (les *LocalStorage) IsFree(timeStart, timeEnd time.Time) bool {
	for _, event := range les.Events {
		if timeStart.Before(event.StartDate) && timeEnd.After(event.StartDate) {
			return false
		}
		if timeStart.After(event.StartDate) && timeEnd.Before(event.EndDate) {
			return false
		}
		if timeStart.Before(event.EndDate) && timeEnd.After(event.EndDate) {
			return false
		}
	}
	return true
}

func (les LocalStorage) EventList() map[int32]events.Event {
	return les.Events
}
