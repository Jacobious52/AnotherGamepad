package main

import (
	"os"
	"os/signal"

	"github.com/Jacobious52/AnotherGamepad/gorilla"

	"github.com/Jacobious52/AnotherGamepad/http"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infoln("starting server")

	handler := http.NewHandler()
	handler.WebSocketService = gorilla.NewWebSocketService()

	server := http.NewServer()
	server.Addr = "localhost:8080"
	server.Handler = handler

	if err := server.Open(); err != nil {
		os.Exit(1)
	}

	log.Infoln("server started")

	defer server.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
}
