package gorilla

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketService struct {
	Upgrader websocket.Upgrader
}

func NewWebSocketService() *WebSocketService {
	u := websocket.Upgrader{}
	u.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	return &WebSocketService{
		Upgrader: u,
	}
}
