package ws

import (
	"github.com/gofiber/contrib/websocket"
)

func (h *WebsocketHandler) chatConnection(currentConnection *websocket.Conn) {
	h.service.Chat.HandleConnection(currentConnection)
}
