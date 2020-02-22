package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tima-fey/otus-golang/calendar/internal/helpers"
)

type MyHandler struct {
	logLvl string
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print((fmt.Sprintf("Got request %v", r.URL.Path)))
	if r.URL.Path == "/hello" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("hello")
	}
}

func main() {
	config := helpers.GetConfig()
	logFile, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	handler := &MyHandler{logLvl: config.LogLevel}
	server := &http.Server{
		Addr:    fmt.Sprintf("%v:%d", config.Address, config.Port),
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
}
