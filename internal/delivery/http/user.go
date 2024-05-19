package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *HTTPHandler) joinChat(ctx *fiber.Ctx) error {
	chatID, err := uuid.Parse(ctx.Params("chat_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	userID, err := uuid.Parse(ctx.Query("user_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if err = h.service.User.JoinChat(chatID, userID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(map[string]bool{"joined": true})
}

func (h *HTTPHandler) leaveChat(ctx *fiber.Ctx) error {
	chatID, err := uuid.Parse(ctx.Params("chat_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	userID, err := uuid.Parse(ctx.Query("user_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if err = h.service.User.LeaveChat(chatID, userID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(map[string]bool{"left": true})
}
