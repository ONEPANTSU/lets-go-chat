package repository

import (
	"github.com/google/uuid"
	"lets-go-chat/internal/domain"
)

type Chat interface {
	CreateChat(chat domain.ChatInDB) (uuid.UUID, error)
	DeleteChat(chatID uuid.UUID) error
	JoinChat(chatID, userID uuid.UUID) error
	LeaveChat(chatID, userID uuid.UUID) error
	IsMember(chatID, userID uuid.UUID) bool
	GetChatUsers(chatID uuid.UUID) ([]uuid.UUID, error)
	GetUserChats(userID uuid.UUID) ([]uuid.UUID, error)
	GetChat(chatID uuid.UUID) (*domain.ChatInDB, error)
}

type Message interface {
	CreateMessage(message *domain.MessageInDB) (int, error)
	GetMessagesByChat(chatID uuid.UUID, offset, limit int) (*[]domain.MessageInDB, error)
	DeleteMessage(messageID int) error
}

type Repository struct {
	Chat
	Message
}
