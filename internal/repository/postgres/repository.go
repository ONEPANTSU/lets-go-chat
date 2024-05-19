package postgres

import (
	"github.com/jackc/pgx/v5"
	"lets-go-chat/internal/repository"
)

func NewPostgresRepository(db *pgx.Conn) *repository.Repository {
	return &repository.Repository{
		Chat:    newChatRepository(db),
		Message: newMessageRepository(db),
	}
}
