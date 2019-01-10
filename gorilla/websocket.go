package gorilla

import (
	"github.com/gorilla/websocket"
)

type WebSocketService struct {
	Upgrader websocket.Upgrader
}

func NewWebSocketService() *WebSocketService {
	return &WebSocketService{
		Upgrader: websocket.Upgrader{},
	}
}
