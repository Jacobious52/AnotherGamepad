package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	anothergamepad "github.com/Jacobious52/AnotherGamepad"
	"github.com/Jacobious52/AnotherGamepad/keyboard"

	"github.com/Jacobious52/AnotherGamepad/gorilla"

	"github.com/gorilla/websocket"

	log "github.com/sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

// Handler is the main handler for the server
type Handler struct {
	*httprouter.Router

	WebSocketService *gorilla.WebSocketService
}

func NewHandler() *Handler {
	h := &Handler{
		Router: httprouter.New(),
	}

	h.GET("/", h.index)
	h.GET("/status", h.status)
	h.GET("/ws", h.upgrade)
	h.NotFound = http.FileServer(http.Dir("static"))

	return h
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "static/index.html")
}

func (h *Handler) status(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("webserver running: %v", time.Now())))
}

func (h *Handler) upgrade(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c, err := h.WebSocketService.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warningln("websocket upgrade failed", err)
		return
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()

		if err != nil {
			log.Warningln("websocket read message failed", err)
			break
		}

		switch mt {
		case websocket.TextMessage:

			var event anothergamepad.Event
			err := json.Unmarshal(message, &event)
			if err != nil {
				log.Errorln("failed to decode json message", err)
				return
			}

			log.Infof("key: \"%s\", state: \"%s\"\n", event.Key, event.State())
			keyboard.KeyToggle(event.Key, event.State())

		case websocket.CloseMessage:
			log.Infoln("close message received")
			return
			// TODO(jacobious52): Ping/Pong
		}
	}
}
