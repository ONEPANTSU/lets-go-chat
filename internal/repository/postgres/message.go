package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"lets-go-chat/internal/domain"
)

type MessageRepository struct {
	db *pgx.Conn
}

func newMessageRepository(db *pgx.Conn) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (r *MessageRepository) CreateMessage(message *domain.MessageInDB) (int, error) {
	query := "insert into message (user_id, chat_id, text) values ($1, $2, $3) returning id"
	row := r.db.QueryRow(context.Background(), query, message.UserID.String(), message.ChatID.String(), message.Text)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MessageRepository) GetMessagesByChat(chatID uuid.UUID, limit, offset int) (*[]domain.MessageInDB, error) {
	query := "select id, user_id, chat_id, text, created_at from message " +
		"where chat_id = $1 and deleted_at is null " +
		"limit $2 offset $3"
	rows, err := r.db.Query(context.Background(), query, chatID.String(), limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := make([]domain.MessageInDB, 0)

	for rows.Next() {
		var message domain.MessageInDB
		if err := rows.Scan(&message.ID, &message.UserID, &message.ChatID, &message.Text, &message.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return &messages, nil
}

func (r *MessageRepository) DeleteMessage(messageID int) error {
	query := "update message set deleted_at = now() where id = $1"
	if _, err := r.db.Exec(context.Background(), query, messageID); err != nil {
		return err
	}
	return nil
}
