package calendar

import (
	"testing"

	"github.com/stretchr/testify/require"
	local_storage "github.com/tima-fey/otus-golang/calendar/internal/pkg/storage"
	events "github.com/tima-fey/otus-golang/calendar/internal/pkg/types"
)

func TestCalendar(t *testing.T) {
	// storage := local_storage.LocalStorage{
	// 	NextID: 0,
	// 	Events: make(map[int]events.Event),
	// }
	storage := new(local_storage.LocalStorage)
	storage.NextID = 0
	storage.Events = make(map[int32]events.Event)
	usersCalendar := Calendar{
		Events: storage,
		User:   "test_user",
		UserID: 1}
	testEvent := events.Event{Name: "testEvent"}
	usersCalendar.Events.AddEvent(testEvent)
	eventFromCalendar, _ := usersCalendar.Events.GetEvent(0)
	require.Equal(t, "testEvent", eventFromCalendar.Name)
}
