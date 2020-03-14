package grpcServer

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"sync"

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
	startDate, err := ptypes.Timestamp(e.StarTtime)
	if err != nil {
		return nil, err
	}
	endDate, err := ptypes.Timestamp(e.EndTime)
	if err != nil {
		return nil, err
	}
	log.Print(fmt.Sprintf("Got request Add %v", events.Event{Name: e.Name, EventID: e.Id, StartDate: startDate, EndDate: endDate}))
	id, err := c.Events.AddEvent(events.Event{Name: e.Name, EventID: e.Id, StartDate: startDate, EndDate: endDate})
	return &scheme.EventId{Id: id}, err
}
func (c calendarpb) Get(ctx context.Context, id *scheme.EventId) (*scheme.Event, error) {
	log.Print(fmt.Sprintf("Got request Get %v", id.Id))
	event, err := c.Events.GetEvent(id.Id)
	if err != nil {
		return nil, err
	}
	startTime, err := ptypes.TimestampProto(event.StartDate)
	if err != nil {
		return nil, err
	}
	endTime, err := ptypes.TimestampProto(event.StartDate)
	return &scheme.Event{Name: event.Name, Id: event.EventID, StarTtime: startTime, EndTime: endTime}, err
}
func (c calendarpb) Update(ctx context.Context, e *scheme.Event) (*scheme.Event, error) {
	startDate, err := ptypes.Timestamp(e.StarTtime)
	if err != nil {
		return nil, err
	}
	endDate, err := ptypes.Timestamp(e.EndTime)
	if err != nil {
		return nil, err
	}
	log.Print(fmt.Sprintf("Got request Update %v", events.Event{Name: e.Name, StartDate: startDate, EndDate: endDate}))
	err = c.Events.ReplaceEvent(e.Id, events.Event{Name: e.Name, StartDate: startDate, EndDate: endDate})
	return e, err
}
func (c calendarpb) Delete(ctx context.Context, id *scheme.EventId) (*scheme.EventId, error) {
	log.Print(fmt.Sprintf("Got request Delete %v", id.Id))
	err := c.Events.RemoveEvent(id.Id)
	return &scheme.EventId{Id: id.Id}, err
}
func StartGrpcServer(storage calendar.Storage, address string, port int, logFile io.Writer, logLvl string, wg *sync.WaitGroup) {
	defer wg.Done()
	log.SetOutput(logFile)
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", address, port))
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	scheme.RegisterEventsHandlerServer(grpcServer, &calendarpb{Events: storage})
	grpcServer.Serve(lis)
}
