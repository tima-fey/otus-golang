package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes"
	"github.com/tima-fey/otus-golang/calendar/internal/api/scheme"
	"github.com/tima-fey/otus-golang/calendar/internal/calendar"
	local_storage "github.com/tima-fey/otus-golang/calendar/internal/pkg/storage"
	events "github.com/tima-fey/otus-golang/calendar/internal/pkg/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type calendarpb struct {
	Events calendar.Storage
}

func (c calendarpb) Add(ctx context.Context, e *scheme.Event) (*scheme.EventId, error) {
	startDate, _ := ptypes.Timestamp(e.StarTtime)
	endDate, _ := ptypes.Timestamp(e.EndTime)
	id, err := c.Events.AddEvent(events.Event{Name: e.Name, EventID: e.Id, StartDate: startDate, EndDate: endDate})
	return &scheme.EventId{Id: id}, err
}
func (c calendarpb) Get(ctx context.Context, id *scheme.EventId) (*scheme.Event, error) {
	event, err := c.Events.GetEvent(id.Id)
	startTime, _ := ptypes.TimestampProto(event.StartDate)
	endTime, _ := ptypes.TimestampProto(event.StartDate)
	return &scheme.Event{Name: event.Name, Id: event.EventID, StarTtime: startTime, EndTime: endTime}, err
}
func (c calendarpb) Update(ctx context.Context, e *scheme.Event) (*scheme.Event, error) {
	startDate, _ := ptypes.Timestamp(e.StarTtime)
	endDate, _ := ptypes.Timestamp(e.EndTime)
	err := c.Events.ReplaceEvent(e.Id, events.Event{Name: e.Name, StartDate: startDate, EndDate: endDate})
	fmt.Println(e)
	return e, err
}
func (c calendarpb) Delete(ctx context.Context, id *scheme.EventId) (*scheme.EventId, error) {
	err := c.Events.RemoveEvent(id.Id)
	return &scheme.EventId{Id: id.Id}, err
}
func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	storage := new(local_storage.LocalStorage)
	storage.NextID = 0
	storage.Events = make(map[int32]events.Event)
	scheme.RegisterEventsHandlerServer(grpcServer, &calendarpb{Events: storage})
	grpcServer.Serve(lis)
}
