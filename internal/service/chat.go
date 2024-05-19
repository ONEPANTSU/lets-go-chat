package service

import (
	"errors"
	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"lets-go-chat/internal/domain"
	"lets-go-chat/internal/repository"
)

type ChatService struct {
	repository  *repository.Repository
	connections map[domain.Chat]*websocket.Conn
	logPrefix   string
}

func newChatService(repository *repository.Repository) *ChatService {
	return &ChatService{
		repository:  repository,
		connections: make(map[domain.Chat]*websocket.Conn),
		logPrefix:   "[ChatService] ",
	}
}

func (s *ChatService) HandleConnection(currentConnection *websocket.Conn) {
	var currentChat domain.Chat
	var err error
	if currentChat.ChatID, err = uuid.Parse(currentConnection.Params("chat_id")); err != nil {
		logrus.Printf(s.logPrefix + err.Error())
		return
	}
	if currentChat.UserID, err = uuid.Parse(currentConnection.Query("user_id")); err != nil {
		logrus.Printf(s.logPrefix + err.Error())
		return
	}

	s.connections[currentChat] = currentConnection
	logrus.Println(currentChat.ChatID.String() + " connected")
	for {
		var message domain.Message
		if err = currentConnection.ReadJSON(&message); err != nil {
			delete(s.connections, currentChat)
			logrus.Printf(s.logPrefix + err.Error())
			return
		}
		err = s.saveMessage(currentChat.ChatID, currentChat.UserID, message)
		if err != nil {
			logrus.Printf(s.logPrefix + err.Error())
		}

		for chat, connection := range s.connections {
			logrus.Printf(chat.UserID.String() + " --- " + currentChat.UserID.String())
			if chat != currentChat && chat.ChatID == currentChat.ChatID {
				logrus.Printf(s.logPrefix + "sending...")
				if err := connection.WriteJSON(message); err != nil {
					logrus.Printf(s.logPrefix + err.Error())
					return
				}
				logrus.Printf(s.logPrefix + "message was sent")
			}
			logrus.Printf(s.logPrefix + "message was handeled")
		}

	}
}

func (s *ChatService) saveMessage(chatID, userID uuid.UUID, message domain.Message) error {

	if !s.repository.IsMember(chatID, userID) {
		return errors.New("user is not a member of the chat")
	}
	if _, err := s.repository.CreateMessage(&domain.MessageInDB{
		ChatID: chatID,
		UserID: userID,
		Text:   message.Text,
	}); err != nil {
		return err
	}
	return nil
}

func (s *ChatService) CreateChat(chat domain.ChatInDB) (uuid.UUID, error) {
	return s.repository.CreateChat(chat)
}

func (s *ChatService) DeleteChat(userID uuid.UUID, chatID uuid.UUID) error {
	chat, err := s.repository.GetChat(chatID)
	if err != nil || chat.OwnerID != userID {
		return errors.New("current user is not an owner of the chat")
	}
	members, err := s.repository.GetChatUsers(chatID)
	if err != nil {
		return err
	}
	for _, memberID := range members {
		if err = s.repository.LeaveChat(chatID, memberID); err != nil {
			return err
		}
	}
	return s.repository.DeleteChat(chatID)
}

func (s *ChatService) GetChat(chatID uuid.UUID) (*domain.ChatInDB, error) {
	return s.repository.GetChat(chatID)
}

func (s *ChatService) GetMembers(chatID uuid.UUID) ([]uuid.UUID, error) {
	return s.repository.GetChatUsers(chatID)
}

func (s *ChatService) GetMessages(chatID uuid.UUID, limit, offset int) (*[]domain.MessageInDB, error) {
	return s.repository.GetMessagesByChat(chatID, limit, offset)
}
