package domain

import "github.com/google/uuid"

type Chat struct {
	UserID uuid.UUID `json:"user_id"`
	ChatID uuid.UUID `json:"chat_id"`
}
