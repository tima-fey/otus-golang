package main

import (
	"log"
	"os"
	"sync"
	"time"

	grpcServer "github.com/tima-fey/otus-golang/calendar/internal/api/server"
	"github.com/tima-fey/otus-golang/calendar/internal/helpers"
	local_storage "github.com/tima-fey/otus-golang/calendar/internal/pkg/storage"
	events "github.com/tima-fey/otus-golang/calendar/internal/pkg/types"
	webServer "github.com/tima-fey/otus-golang/calendar/internal/web"
)

func main() {
	storage := new(local_storage.LocalStorage)
	storage.NextID = 0
	storage.Events = make(map[int32]events.Event)
	config := helpers.GetConfig()
	logFile, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	var wg sync.WaitGroup
	go grpcServer.StartGrpcServer(storage, config.Address, config.GrpcPort, logFile, config.LogLevel, &wg)
	go webServer.StartWebServer(config.Address, config.WebPort, logFile, config.LogLevel, &wg)
	time.Sleep((time.Minute * 2))
	wg.Wait()
}
