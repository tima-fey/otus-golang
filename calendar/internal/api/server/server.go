package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes"
	"github.com/tima-fey/otus-golang/calendar/internal/api/scheme"
	"github.com/tima-fey/otus-golang/calendar/internal/calendar"
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
	fmt.Println(e)
	return &scheme.Event{Id: 1}, nil
}
func (c calendarpb) Delete(ctx context.Context, id *scheme.EventId) (*scheme.EventId, error) {
	fmt.Println(id)
	return id, nil
}
func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	scheme.RegisterEventsHandlerServer(grpcServer, &calendarpb{})
	grpcServer.Serve(lis)
}
