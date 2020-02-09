package local_storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tima-fey/otus-golang/calendar/internal/pkg/errors"
	events "github.com/tima-fey/otus-golang/calendar/internal/pkg/types"
)

func TestAddEvent(t *testing.T) {
	storage := new(LocalStorage)
	storage.NextID = 0
	storage.Events = make(map[int]events.Event)
	startTime := time.Date(2009, time.November, 10, 10, 0, 0, 0, time.UTC)
	endTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	testEvent := events.Event{Name: "test_name", StartDate: startTime, EndDate: endTime}
	storage.AddEvent(testEvent)
	require.Equal(t, 1, storage.NextID)
	require.Equal(t, testEvent, storage.Events[0])
	startTime2 := time.Date(2009, time.November, 10, 14, 30, 0, 0, time.UTC)
	endTime2 := time.Date(2009, time.November, 10, 15, 30, 0, 0, time.UTC)
	testEvent2 := events.Event{Name: "test_name", StartDate: startTime2, EndDate: endTime2}
	_, err := storage.AddEvent(testEvent2)
	require.Equal(t, &errors.ErrDateBusy{}, err)
}

func TestRemoveEvent(t *testing.T) {
	storage := new(LocalStorage)
	storage.NextID = 0
	storage.Events = make(map[int]events.Event)
	startTime := time.Date(2009, time.November, 10, 22, 0, 0, 0, time.UTC)
	endTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	testEvent := events.Event{Name: "test_name", StartDate: startTime, EndDate: endTime}
	storage.AddEvent(testEvent)
	storage.RemoveEvent(0)
	require.Equal(t, 1, storage.NextID)
	require.Equal(t, 0, len(storage.Events))
	err := storage.RemoveEvent(0)
	require.Equal(t, &errors.ErrNotSuchID{}, err)
}

func TestReplaceEvent(t *testing.T) {
	storage := new(LocalStorage)
	storage.NextID = 0
	storage.Events = make(map[int]events.Event)
	startTime := time.Date(2009, time.November, 10, 12, 0, 0, 0, time.UTC)
	endTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	testEvent := events.Event{Name: "test_name", StartDate: startTime, EndDate: endTime}
	storage.AddEvent(testEvent)
	testEvent2 := events.Event{Name: "test_name2", StartDate: startTime, EndDate: endTime}
	storage.ReplaceEvent(0, testEvent2)
	require.Equal(t, 1, storage.NextID)
	require.Equal(t, testEvent2, storage.Events[0])
	err := storage.ReplaceEvent(1, testEvent2)
	require.Equal(t, &errors.ErrNotSuchID{}, err)
	startTime3 := time.Date(2009, time.November, 11, 12, 0, 0, 0, time.UTC)
	endTime3 := time.Date(2009, time.November, 11, 23, 0, 0, 0, time.UTC)
	testEvent3 := events.Event{Name: "test_name3", StartDate: startTime3, EndDate: endTime3}
	storage.AddEvent(testEvent3)
	startTime4 := time.Date(2009, time.November, 10, 10, 0, 0, 0, time.UTC)
	endTime4 := time.Date(2009, time.November, 10, 22, 0, 0, 0, time.UTC)
	testEvent4 := events.Event{Name: "test_name4", StartDate: startTime4, EndDate: endTime4}
	err = storage.ReplaceEvent(1, testEvent4)
	require.Equal(t, &errors.ErrDateBusy{}, err)
	require.Equal(t, testEvent3.Name, storage.Events[1].Name)
	require.Equal(t, 2, len(storage.Events))
}
func TestIsFree(t *testing.T) {
	storage := new(LocalStorage)
	storage.NextID = 0
	storage.Events = make(map[int]events.Event)
	startTime := time.Date(2009, time.November, 10, 12, 0, 0, 0, time.UTC)
	endTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	testEvent := events.Event{Name: "test_name", StartDate: startTime, EndDate: endTime}
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
	storage.NextID = 0
	storage.Events = make(map[int]events.Event)
	startTime := time.Date(2009, time.November, 10, 12, 0, 0, 0, time.UTC)
	endTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	testEvent := events.Event{Name: "test_name", StartDate: startTime, EndDate: endTime}
	storage.AddEvent(testEvent)
	answer := storage.EventList()
	require.Equal(t, storage.Events, answer)
}
