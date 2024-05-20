package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary Join Chat
// @Tags user
// @Description Join to concrete chat by its uuid
// @Accept json
// @Produce json
// @Param chat_id path string true "Chat ID" format(uuid)
// @Param user_id query string true "User ID" format(uuid)
// @Success 200 {object} Response "Returns true if the user is successfully joined"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /api/user/join/{chat_id} [post]
func (h *HTTPHandler) joinChat(ctx *fiber.Ctx) error {
	chatID, err := uuid.Parse(ctx.Params("chat_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	userID, err := uuid.Parse(ctx.Query("user_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err = h.service.User.JoinChat(chatID, userID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(Response{Success: true})
}

// @Summary Leave Chat
// @Tags user
// @Description Leave concrete chat by its uuid
// @Accept json
// @Produce json
// @Param chat_id path string true "Chat ID" format(uuid)
// @Param user_id query string true "User ID" format(uuid)
// @Success 200 {object} Response "Returns true if the user is successfully left"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /api/user/leave/{chat_id} [post]
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
	return ctx.JSON(Response{Success: true})
}
