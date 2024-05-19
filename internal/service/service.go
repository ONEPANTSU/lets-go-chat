package service

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
	"lets-go-chat/internal/domain"
	"lets-go-chat/internal/repository"
)

type Chat interface {
	HandleConnection(connection *websocket.Conn)
	saveMessage(chatID, userID uuid.UUID, message domain.Message) error
	CreateChat(chat domain.ChatInDB) (uuid.UUID, error)
	DeleteChat(userID uuid.UUID, chatID uuid.UUID) error
	GetChat(chatID uuid.UUID) (*domain.ChatInDB, error)
	GetMembers(chatID uuid.UUID) ([]uuid.UUID, error)
	GetMessages(chatID uuid.UUID, limit, offset int) (*[]domain.MessageInDB, error)
}

type User interface {
	JoinChat(chatID, userID uuid.UUID) error
	LeaveChat(chatID, userID uuid.UUID) error
}

type Service struct {
	Chat
	User
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Chat: newChatService(repository),
		User: newUserService(repository),
	}
}
