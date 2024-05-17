package ws

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/sirupsen/logrus"
	"lets-go-chat/internal/domain"
)

var logPrefix = "Chat Websocket Logs: "

var connections = make(map[domain.Chat]*websocket.Conn)

func (h *WebsocketHandler) chatConnection(currentConnection *websocket.Conn) {
	var currentChat domain.Chat
	if err := currentConnection.ReadJSON(&currentChat); err != nil {
		logrus.Printf(logPrefix + err.Error())
		return
	}
	connections[currentChat] = currentConnection
	logrus.Println(currentChat.ChatID.String() + " connected")
	for {
		var message domain.Message
		if err := currentConnection.ReadJSON(&message); err != nil {
			delete(connections, currentChat)
			logrus.Printf(logPrefix + err.Error())
			return
		}
		for chat, connection := range connections {
			logrus.Printf(chat.UserID.String() + " --- " + currentChat.UserID.String())
			if chat != currentChat && chat.ChatID == currentChat.ChatID {
				logrus.Printf(logPrefix + "sending...")
				if err := connection.WriteJSON(message); err != nil {
					logrus.Printf(logPrefix + err.Error())
					return
				}
				logrus.Printf(logPrefix + "message was sent")
			}
			logrus.Printf(logPrefix + "message was handeled")
		}

	}
}
