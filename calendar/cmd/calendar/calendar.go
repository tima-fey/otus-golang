package main

import (
	"fmt"

	"github.com/tima-fey/otus-golang/calendar/internal/calendar"
	local_storage "github.com/tima-fey/otus-golang/calendar/internal/pkg/storage"
	events "github.com/tima-fey/otus-golang/calendar/internal/pkg/types"
)

func main() {
	storage := new(local_storage.LocalStorage)
	storage.NextID = 0
	storage.Events = make(map[int32]events.Event)
	usersCalendar := calendar.Calendar{
		Events: storage,
		User:   "test_user",
		UserID: 1}
	testEvent := events.Event{Name: "testEvent"}
	usersCalendar.Events.AddEvent(testEvent)
	eventFromCalendar, _ := usersCalendar.Events.GetEvent(0)
	fmt.Println(eventFromCalendar.Name)

}
