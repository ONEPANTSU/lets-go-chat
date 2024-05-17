package app

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"lets-go-chat/internal/config"
	wsHandler "lets-go-chat/internal/delivery/ws"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	cfg *config.Config
}

func NewApp(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (app App) Start() {
	db, err := pgx.Connect(context.Background(), app.cfg.DB.GetConnectionDSN())
	if err != nil {
		logrus.Fatal(err)
	}

	server := fiber.New()
	handler := wsHandler.NewWebsocketHandler(1, server.Group("/api"))
	handler.InitRoutes()

	go func() {
		if err := server.Listen(":" + app.cfg.App.HTTPPort); err != nil {
			logrus.Fatalf("error occurred while running server: %s", err.Error())
		}
	}()

	logrus.Info("app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Info("app stopping")

	if err := server.Shutdown(); err != nil {
		logrus.Fatalf("error occurred while shutting down server: %s", err.Error())
		return
	}
	if err := db.Close(context.Background()); err != nil {
		logrus.Fatalf("error occurred while closing db connection: %s", err.Error())
		return
	}
}
