package service

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"lets-go-chat/internal/domain"
)

type ChatService struct {
	connections map[domain.Chat]*websocket.Conn
	logPrefix   string
}

func newChatService() *ChatService {
	return &ChatService{
		connections: make(map[domain.Chat]*websocket.Conn),
		logPrefix:   "[ChatService] ",
	}
}

func (s ChatService) HandleConnection(currentConnection *websocket.Conn) {
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

func (s ChatService) saveMessage(chat domain.Chat, message domain.Message) error {
	//TODO implement me
	panic("implement me")
}
