package app

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"lets-go-chat/internal/config"
	"lets-go-chat/internal/database"
	wsHandler "lets-go-chat/internal/delivery/ws"
	"lets-go-chat/internal/server"
	wsServer "lets-go-chat/internal/server/ws"
)

type App struct {
	server server.Server
	db     *sql.DB
}

func NewApp(cfg *config.Config) *App {
	db, err := database.ConnectToPostgres(cfg.DB)
	if err != nil {
		logrus.Fatal(err)
	}
	handler := wsHandler.NewWebsocketHandler(1)

	return &App{
		server: wsServer.NewWebsocketServer(cfg.App.WebsocketPort, handler.InitRoutes()),
	}
}
