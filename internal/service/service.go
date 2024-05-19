package service

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
	"lets-go-chat/internal/domain"
)

type Chat interface {
	HandleConnection(connection *websocket.Conn)
	saveMessage(chat domain.Chat, message domain.Message) error
}

type User interface {
	CreateNewChat(userID uuid.UUID) (uuid.UUID, error)
	DeleteChat(userID uuid.UUID, chatID uuid.UUID) error
	JoinChat(userID uuid.UUID, chatID uuid.UUID) error
	LeftChat(userID uuid.UUID, chatID uuid.UUID) error
}

type Service struct {
	Chat
	User
}

func NewService() *Service {
	return &Service{
		Chat: newChatService(),
		User: nil,
	}
}
