package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"lets-go-chat/internal/domain"
	"strconv"
)

// @Summary Get Chat
// @Tags chat
// @Description Get chat by its uuid
// @Accept json
// @Produce json
// @Param chat_id path string true "Chat ID" format(uuid)
// @Success 200 {object} Response
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /api/chat/{chat_id} [get]
func (h *HTTPHandler) getChat(ctx *fiber.Ctx) error {
	chatID, err := uuid.Parse(ctx.Params("chat_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	chat, err := h.service.GetChat(chatID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(Response{Success: true, Data: chat})
}

// @Summary Create Chat
// @Tags chat
// @Description Create new chat
// @Accept json
// @Produce json
// @Param chat_id path string true "Chat ID" format(uuid)
// @Param user_id query string true "User ID" format(uuid)
// @Success 200 {object} Response
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /api/chat [post]
func (h *HTTPHandler) createChat(ctx *fiber.Ctx) error {
	var chat domain.ChatInDB
	if err := ctx.BodyParser(&chat); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	chatID, err := h.service.CreateChat(chat)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(Response{Success: true, Data: chatID.String()})
}

// @Summary Delete Chat
// @Tags chat
// @Description Delete chat by its uuid
// @Accept json
// @Produce json
// @Param chat_id path string true "Chat ID" format(uuid)
// @Param user_id query string true "User ID" format(uuid)
// @Success 200 {object} Response
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /api/chat/{chat_id} [delete]
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
	return ctx.JSON(Response{Success: true})
}

// @Summary Get Chats Members
// @Tags chat
// @Description Get members uuids from concrete chat
// @Accept json
// @Produce json
// @Param chat_id path string true "Chat ID" format(uuid)
// @Success 200 {object} Response
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /api/chat/{chat_id}/members [get]
func (h *HTTPHandler) getMembers(ctx *fiber.Ctx) error {
	chatID, err := uuid.Parse(ctx.Params("chat_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	members, err := h.service.GetMembers(chatID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(Response{Success: true, Data: members})
}

// @Summary Get Chats Messages
// @Tags chat
// @Description Get messages with offset and limit from concrete chat
// @Accept json
// @Produce json
// @Param chat_id path string true "Chat ID" format(uuid)
// @Param limit query integer true "Limit"
// @Param offset query integer true "Limit"
// @Success 200 {object} Response
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /api/chat/{chat_id}/messages [get]
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
	return ctx.JSON(Response{Success: true, Data: messages})
}
