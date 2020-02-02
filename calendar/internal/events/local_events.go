package events

import "time"

type LocalStorage struct {
	Events map[int]Event
	nextId int
}

func (les *LocalStorage) AddEvent(e Event) error {
	if !les.IsFree(e.startDate, e.endDate) {
		return &ErrDateBusy{}
	}
	les.Events[les.nextId] = e
	les.nextId++
	return nil
}
func (les *LocalStorage) RemoveEvent(eventId int) error {
	if _, ok := les.Events[eventId]; ok {
		delete(les.Events, eventId)
		return nil
	}
	return &ErrNotSuchId{}
}
func (les *LocalStorage) ReplaceEvent(eventId int, e Event) error {
	tempEvent, ok := les.Events[eventId]
	if ok {
		delete(les.Events, eventId)
	} else {
		return &ErrNotSuchId{}
	}
	if les.IsFree(e.startDate, e.endDate) {
		les.Events[eventId] = e
		return nil
	}
	les.Events[eventId] = tempEvent
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
