package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAddEvent(t *testing.T) {
	storage := new(LocalStorage)
	storage.nextId = 0
	storage.Events = make(map[int]Event)
	startTime := time.Date(2009, time.November, 10, 10, 0, 0, 0, time.UTC)
	endTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	testEvent := Event{name: "test_name", startDate: startTime, endDate: endTime}
	storage.AddEvent(testEvent)
	require.Equal(t, 1, storage.nextId)
	require.Equal(t, testEvent, storage.Events[0])
	startTime2 := time.Date(2009, time.November, 10, 14, 30, 0, 0, time.UTC)
	endTime2 := time.Date(2009, time.November, 10, 15, 30, 0, 0, time.UTC)
	testEvent2 := Event{name: "test_name", startDate: startTime2, endDate: endTime2}
	err := storage.AddEvent(testEvent2)
	require.Equal(t, &ErrDateBusy{}, err)

}

func TestRemoveEvent(t *testing.T) {
	storage := new(LocalStorage)
	storage.nextId = 0
	storage.Events = make(map[int]Event)
	startTime := time.Date(2009, time.November, 10, 22, 0, 0, 0, time.UTC)
	endTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	testEvent := Event{name: "test_name", startDate: startTime, endDate: endTime}
	storage.AddEvent(testEvent)
	storage.RemoveEvent(0)
	require.Equal(t, 1, storage.nextId)
	require.Equal(t, 0, len(storage.Events))
	err := storage.RemoveEvent(0)
	require.Equal(t, &ErrNotSuchId{}, err)
}

func TestReplaceEvent(t *testing.T) {
	storage := new(LocalStorage)
	storage.nextId = 0
	storage.Events = make(map[int]Event)
	startTime := time.Date(2009, time.November, 10, 12, 0, 0, 0, time.UTC)
	endTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	testEvent := Event{name: "test_name", startDate: startTime, endDate: endTime}
	storage.AddEvent(testEvent)
	testEvent2 := Event{name: "test_name2", startDate: startTime, endDate: endTime}
	storage.ReplaceEvent(0, testEvent2)
	require.Equal(t, 1, storage.nextId)
	require.Equal(t, testEvent2, storage.Events[0])
	err := storage.ReplaceEvent(1, testEvent2)
	require.Equal(t, &ErrNotSuchId{}, err)
	startTime3 := time.Date(2009, time.November, 11, 12, 0, 0, 0, time.UTC)
	endTime3 := time.Date(2009, time.November, 11, 23, 0, 0, 0, time.UTC)
	testEvent3 := Event{name: "test_name3", startDate: startTime3, endDate: endTime3}
	storage.AddEvent(testEvent3)
	startTime4 := time.Date(2009, time.November, 10, 10, 0, 0, 0, time.UTC)
	endTime4 := time.Date(2009, time.November, 10, 22, 0, 0, 0, time.UTC)
	testEvent4 := Event{name: "test_name4", startDate: startTime4, endDate: endTime4}
	err = storage.ReplaceEvent(1, testEvent4)
	require.Equal(t, &ErrDateBusy{}, err)
	require.Equal(t, testEvent3, storage.Events[1])
	require.Equal(t, 2, len(storage.Events))
}
func TestIsFree(t *testing.T) {
	storage := new(LocalStorage)
	storage.nextId = 0
	storage.Events = make(map[int]Event)
	startTime := time.Date(2009, time.November, 10, 12, 0, 0, 0, time.UTC)
	endTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	testEvent := Event{name: "test_name", startDate: startTime, endDate: endTime}
	storage.AddEvent(testEvent)
	answer := storage.IsFree(
		time.Date(2009, time.November, 10, 11, 0, 0, 0, time.UTC),
		time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	require.Equal(t, false, answer)
	answer = storage.IsFree(
		time.Date(2009, time.November, 10, 11, 0, 0, 0, time.UTC),
		time.Date(2009, time.November, 10, 23, 30, 0, 0, time.UTC))
	require.Equal(t, false, answer)
	answer = storage.IsFree(
		time.Date(2009, time.November, 10, 13, 0, 0, 0, time.UTC),
		time.Date(2009, time.November, 11, 23, 30, 0, 0, time.UTC))
	require.Equal(t, false, answer)
	answer = storage.IsFree(
		time.Date(2009, time.November, 11, 13, 0, 0, 0, time.UTC),
		time.Date(2009, time.November, 11, 23, 30, 0, 0, time.UTC))
	require.Equal(t, true, answer)
	answer = storage.IsFree(
		time.Date(2009, time.November, 9, 13, 0, 0, 0, time.UTC),
		time.Date(2009, time.November, 9, 23, 30, 0, 0, time.UTC))
	require.Equal(t, true, answer)
}
func TestEventList(t *testing.T) {
	storage := new(LocalStorage)
	storage.nextId = 0
	storage.Events = make(map[int]Event)
	startTime := time.Date(2009, time.November, 10, 12, 0, 0, 0, time.UTC)
	endTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	testEvent := Event{name: "test_name", startDate: startTime, endDate: endTime}
	storage.AddEvent(testEvent)
	answer := storage.EventList()
	require.Equal(t, storage.Events, answer)
}
