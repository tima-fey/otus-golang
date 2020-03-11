package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/tima-fey/otus-golang/calendar/internal/api/scheme"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := scheme.NewEventsHandlerClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	_, _ = c.Add(ctx, &scheme.Event{Name: "test", StarTtime: ptypes.TimestampNow(), EndTime: ptypes.TimestampNow()})
	answer_to_print, _ := c.Get(ctx, &scheme.EventId{Id: 0})
	fmt.Println(answer_to_print)
}
