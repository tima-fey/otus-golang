package main

import "github.com/tima-fey/otus-golang/calendar/internal/events"

type Calendar struct {
	Events events.Storage
	User   string
	UserID int
}
