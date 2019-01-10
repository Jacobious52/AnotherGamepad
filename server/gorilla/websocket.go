package gorilla

import (
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	Upgrade websocket.Upgrader
}
