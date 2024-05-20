package main

import (
	"github.com/sirupsen/logrus"
	"lets-go-chat/internal/app"
	"lets-go-chat/internal/config"
)

// @title           Let's Go Chat API
// @version         1.0
// @description     This is an API for Let's Go Chat Application
// @host localhost:8080
// @BasePath  /
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg := config.NewConfig()
	application := app.NewApp(cfg)
	application.Start()
}
