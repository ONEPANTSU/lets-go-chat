package domain

import (
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	UserID uuid.UUID `json:"user_id"`
	ChatID uuid.UUID `json:"chat_id"`
}

type ChatInDB struct {
	ID        uuid.UUID `json:"id" db:"id"`
	OwnerID   uuid.UUID `json:"user_id" db:"owner_id" binding:"required"`
	Name      string    `json:"name" db:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at" binding:"required"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
}
