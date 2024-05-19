package domain

import (
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	UserID uuid.UUID `json:"user_id"`
	ChatID uuid.UUID `json:"chat_id"`
}

type Message struct {
	Text string `json:"text"`
}

type ChatInDB struct {
	ID        int       `json:"id" db:"id"`
	OwnerID   uuid.UUID `json:"user_id" db:"owner_id" binding:"required"`
	Name      string    `json:"text" db:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at" binding:"required"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
}

type MessageInDB struct {
	ID        int       `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id" binding:"required"`
	ChatID    uuid.UUID `json:"chat_id" db:"chat_id" binding:"required"`
	Text      string    `json:"text" db:"text" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at" binding:"required"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
}
