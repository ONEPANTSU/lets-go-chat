package http

import (
	"github.com/gofiber/fiber/v2"
	"lets-go-chat/internal/service"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type HTTPHandler struct {
	service *service.Service
	group   fiber.Router
}

func NewHTTPHandler(service *service.Service, group fiber.Router) *HTTPHandler {
	return &HTTPHandler{
		service: service,
		group:   group,
	}
}

func (h *HTTPHandler) InitRoutes() {

	chatGroup := h.group.Group("/chat")
	chatGroup.Get("/:chat_id<guid>", h.getChat)
	chatGroup.Post("/", h.createChat)
	chatGroup.Delete("/:chat_id<guid>", h.deleteChat)
	chatGroup.Get("/:chat_id<guid>/members", h.getMembers)
	chatGroup.Get("/:chat_id<guid>/messages", h.getMessages)

	userGroup := h.group.Group("/user")
	userGroup.Post("/join/:chat_id<guid>", h.joinChat)
	userGroup.Post("/leave/:chat_id<guid>", h.leaveChat)
}
