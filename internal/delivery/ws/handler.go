package ws

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"lets-go-chat/internal/service"
)

type WebsocketHandler struct {
	service *service.Service
	group   fiber.Router
}

func NewWebsocketHandler(service *service.Service, group fiber.Router) *WebsocketHandler {
	return &WebsocketHandler{
		service: service,
		group:   group,
	}
}

func (h *WebsocketHandler) InitRoutes() {
	h.group.Get("/chat/:chat_id", websocket.New(h.chatConnection))
}
