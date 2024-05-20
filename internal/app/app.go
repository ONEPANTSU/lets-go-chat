package app

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/fiber-swagger"

	_ "lets-go-chat/docs"
	"lets-go-chat/internal/config"
	"lets-go-chat/internal/delivery"
	httpHandler "lets-go-chat/internal/delivery/http"
	wsHandler "lets-go-chat/internal/delivery/ws"
	repositoryPool "lets-go-chat/internal/repository/postgres"
	servicePool "lets-go-chat/internal/service"
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
	server.Get("/swagger/*", fiberSwagger.WrapHandler)
	api := server.Group("/api")

	repository := repositoryPool.NewPostgresRepository(db)
	service := servicePool.NewService(repository)
	handlers := []delivery.Handler{
		wsHandler.NewWebsocketHandler(service, api.Group("/ws")),
		httpHandler.NewHTTPHandler(service, api),
	}
	for _, handler := range handlers {
		handler.InitRoutes()
	}

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
