package database

import (
	"lets-go-chat/internal/config"
)

func ConnectToPostgres(cfg *config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open(cfg.DBDriver, cfg.GetConnectionDSN())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
