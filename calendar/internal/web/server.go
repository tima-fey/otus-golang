package webServer

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
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
func StartWebServer(address string, port int, logFile io.Writer, logLvl string, wg *sync.WaitGroup) {
	defer wg.Done()
	log.SetOutput(logFile)
	handler := &MyHandler{logLvl: logLvl}
	server := &http.Server{
		Addr:    fmt.Sprintf("%v:%d", address, port),
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
}
