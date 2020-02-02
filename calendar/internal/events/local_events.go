package events

import "time"

type LocalStorage struct {
	Events map[int]Event
	nextID int
}

func (les *LocalStorage) AddEvent(e Event) error {
	if !les.IsFree(e.startDate, e.endDate) {
		return &ErrDateBusy{}
	}
	les.Events[les.nextID] = e
	les.nextID++
	return nil
}
func (les *LocalStorage) RemoveEvent(eventID int) error {
	if _, ok := les.Events[eventID]; ok {
		delete(les.Events, eventID)
		return nil
	}
	return &ErrNotSuchID{}
}
func (les *LocalStorage) ReplaceEvent(eventID int, e Event) error {
	tempEvent, ok := les.Events[eventID]
	if ok {
		delete(les.Events, eventID)
	} else {
		return &ErrNotSuchID{}
	}
	if les.IsFree(e.startDate, e.endDate) {
		les.Events[eventID] = e
		return nil
	}
	les.Events[eventID] = tempEvent
	return &ErrDateBusy{}
}
func (les *LocalStorage) IsFree(timeStart, timeEnd time.Time) bool {
	for _, event := range les.Events {
		if timeStart.Before(event.startDate) && timeEnd.After(event.startDate) {
			return false
		}
		if timeStart.After(event.startDate) && timeEnd.Before(event.endDate) {
			return false
		}
		if timeStart.Before(event.endDate) && timeEnd.After(event.endDate) {
			return false
		}
	}
	return true
}

func (les LocalStorage) EventList() map[int]Event {
	return les.Events
}
