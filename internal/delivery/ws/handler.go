package ws

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type WebsocketHandler struct {
	service int
}

func NewWebsocketHandler(service int) *WebsocketHandler {
	return &WebsocketHandler{service: service}
}

func (h *WebsocketHandler) InitRoutes() *fiber.App {
	router := fiber.New()

	router.Get("/ws/chat", websocket.New(chatConnection))

	return router
}
