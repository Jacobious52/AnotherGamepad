package gorilla

import (
	"github.com/gorilla/websocket"
)

type WebSocketService struct {
	websocket.Upgrader
}
