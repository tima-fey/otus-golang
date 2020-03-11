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

	answer1, _ := c.Add(ctx, &scheme.Event{Name: "test", StarTtime: ptypes.TimestampNow(), EndTime: ptypes.TimestampNow()})
	fmt.Println(answer1)

	answer2, _ := c.Get(ctx, &scheme.EventId{Id: 0})
	fmt.Println(answer2)
	answer3, _ := c.Update(ctx, &scheme.Event{Name: "tost", Id: 0, StarTtime: ptypes.TimestampNow(), EndTime: ptypes.TimestampNow()})
	fmt.Println(answer3)
	answer2, _ = c.Get(ctx, &scheme.EventId{Id: 0})
	fmt.Println(answer2)
	answer4, _ := c.Delete(ctx, &scheme.EventId{Id: 0})
	fmt.Println(answer4)
	answer2, _ = c.Get(ctx, &scheme.EventId{Id: 0})
	fmt.Println(answer2)
}
