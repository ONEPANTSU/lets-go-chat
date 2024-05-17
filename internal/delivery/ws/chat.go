package ws

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/sirupsen/logrus"
	"lets-go-chat/internal/domain"
)

var logPrefix = "Chat Websocket Logs:\t"

var connections = make(map[domain.Chat]*websocket.Conn)

func chatConnection(currentConnection *websocket.Conn) {
	var chat domain.Chat
	if err := currentConnection.ReadJSON(&chat); err != nil {
		logrus.Printf(logPrefix + err.Error())
		return
	}
	connections[chat] = currentConnection

	for {
		var message string
		if err := currentConnection.ReadJSON(&message); err != nil {
			delete(connections, chat)
			logrus.Printf(logPrefix + err.Error())
			return
		}
		for _, connection := range connections {
			if connection != currentConnection {
				if err := connection.WriteJSON(message); err != nil {
					logrus.Printf(logPrefix + err.Error())
					return
				}
			}
		}

	}
}
