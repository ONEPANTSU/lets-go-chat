package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"lets-go-chat/internal/domain"
	"strconv"
)

func (h *HTTPHandler) getChat(ctx *fiber.Ctx) error {
	chatID, err := uuid.Parse(ctx.Params("chat_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	chat, err := h.service.GetChat(chatID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(chat)
}

func (h *HTTPHandler) createChat(ctx *fiber.Ctx) error {
	var chat domain.ChatInDB
	if err := ctx.BodyParser(&chat); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	chatID, err := h.service.CreateChat(chat)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(map[string]string{"chat_id": chatID.String()})
}

func (h *HTTPHandler) deleteChat(ctx *fiber.Ctx) error {
	chatID, err := uuid.Parse(ctx.Params("chat_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	userID, err := uuid.Parse(ctx.Query("user_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = h.service.DeleteChat(chatID, userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(map[string]bool{"deleted": true})
}

func (h *HTTPHandler) getMembers(ctx *fiber.Ctx) error {
	chatID, err := uuid.Parse(ctx.Params("chat_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	members, err := h.service.GetMembers(chatID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(members)
}

func (h *HTTPHandler) getMessages(ctx *fiber.Ctx) error {
	chatID, err := uuid.Parse(ctx.Params("chat_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}
	messages, err := h.service.GetMessages(chatID, limit, offset)
	if err != nil {
		return err
	}
	return ctx.JSON(messages)
}
