package main

import (
	"os"

	"github.com/Jacobious52/AnotherGamepad/gorilla"
	"github.com/Jacobious52/AnotherGamepad/http"
	"github.com/Jacobious52/AnotherGamepad/systray"

	log "github.com/sirupsen/logrus"
)

func serveForever() *http.Server {
	log.Infoln("starting server")

	handler := http.NewHandler()
	handler.WebSocketService = gorilla.NewWebSocketService()

	server := http.NewServer()
	server.Addr = ":8080"
	server.Handler = handler

	if err := server.Open(); err != nil {
		os.Exit(1)
	}

	log.Infoln("server started")
	return server
}

func main() {
	log.Infoln("Starting app")
	server := serveForever()
	systray.LaunchTray(server)
}
