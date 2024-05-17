package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type DBConfig struct {
	DBDriver string
	DBName   string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPasswd string
	Queries  string
	SSLMode  string
}

func newDBConfig() *DBConfig {
	return &DBConfig{
		DBDriver: viper.GetString("db.driver"),
		DBName:   viper.GetString("db.dbname"),
		DBHost:   viper.GetString("db.host"),
		DBPort:   viper.GetString("db.port"),
		DBUser:   viper.GetString("db.username"),
		Queries:  viper.GetString("db.queries"),
		DBPasswd: os.Getenv("DB_PASSWORD"),
	}
}

func (cfg *DBConfig) GetConnectionDSN() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?%s",
		cfg.DBDriver,
		cfg.DBUser,
		cfg.DBPasswd,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.Queries,
	)
}
