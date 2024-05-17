package config

import "github.com/spf13/viper"

type AppConfig struct {
	HTTPPort      string
	WebsocketPort string
}

func newAppConfig() *AppConfig {
	return &AppConfig{
		HTTPPort:      viper.GetString("app.http_port"),
		WebsocketPort: viper.GetString("app.ws_port"),
	}
}
