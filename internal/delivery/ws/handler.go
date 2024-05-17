package ws

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type WebsocketHandler struct {
	service int
	group   fiber.Router
}

func NewWebsocketHandler(service int, group fiber.Router) *WebsocketHandler {
	return &WebsocketHandler{service: service, group: group}
}

func (h *WebsocketHandler) InitRoutes() {
	h.group.Get("/ws/chat", websocket.New(h.chatConnection))
}
