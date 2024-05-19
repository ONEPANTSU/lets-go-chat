package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"lets-go-chat/internal/domain"
)

type ChatRepository struct {
	db *pgx.Conn
}

func newChatRepository(db *pgx.Conn) *ChatRepository {
	return &ChatRepository{
		db: db,
	}
}

func (r *ChatRepository) CreateChat(chat domain.ChatInDB) (uuid.UUID, error) {
	query := "insert into chat (name, owner_id) values ($1, $2) returning id"
	row := r.db.QueryRow(context.Background(), query, chat.Name, chat.OwnerID.String())
	var chatID uuid.UUID
	if err := row.Scan(&chatID); err != nil {
		return uuid.Nil, err
	}
	return chatID, nil
}

func (r *ChatRepository) DeleteChat(chatID uuid.UUID) error {
	query := "update chat set deleted_at = now() where id = $1"
	if _, err := r.db.Exec(context.Background(), query, chatID.String()); err != nil {
		return err
	}
	return nil
}

func (r *ChatRepository) JoinChat(chatID, userID uuid.UUID) error {
	query := "insert into user_chat (user_id, chat_id) values ($1, $2)"
	if _, err := r.db.Exec(context.Background(), query, userID.String(), chatID.String()); err != nil {
		return err
	}
	return nil
}

func (r *ChatRepository) LeaveChat(chatID, userID uuid.UUID) error {
	query := "delete from user_chat where user_id = $1 and chat_id = $2"
	if _, err := r.db.Exec(context.Background(), query, userID.String(), chatID.String()); err != nil {
		return err
	}
	return nil
}

func (r *ChatRepository) IsMember(chatID, userID uuid.UUID) bool {
	query := "select count(*) from user_chat where user_id = $1 and chat_id = $2"
	row := r.db.QueryRow(context.Background(), query, userID.String(), chatID.String())
	var count int
	if err := row.Scan(&count); err != nil {
		return false
	}
	return count == 1
}

func (r *ChatRepository) GetChatUsers(chatID uuid.UUID) ([]uuid.UUID, error) {
	query := "select user_id from user_chat WHERE chat_id = $1"
	rows, err := r.db.Query(context.Background(), query, chatID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]uuid.UUID, 0)

	for rows.Next() {
		var userID uuid.UUID
		if err := rows.Scan(&userID); err != nil {
			return nil, err
		}
		users = append(users, userID)
	}
	return users, nil
}

func (r *ChatRepository) GetUserChats(userID uuid.UUID) ([]uuid.UUID, error) {
	query := "select chat_id from user_chat WHERE user_id = $1 and deleted_at is null"
	rows, err := r.db.Query(context.Background(), query, userID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	chats := make([]uuid.UUID, 0)

	for rows.Next() {
		var chatID uuid.UUID
		if err := rows.Scan(&chatID); err != nil {
			return nil, err
		}
		chats = append(chats, chatID)
	}
	return chats, nil
}

func (r *ChatRepository) GetChat(chatID uuid.UUID) (*domain.ChatInDB, error) {
	query := "select id, name, owner_id, created_at from chat where id = $1"
	row := r.db.QueryRow(context.Background(), query, chatID.String())
	var chat domain.ChatInDB
	if err := row.Scan(&chat.ID, &chat.Name, &chat.OwnerID, &chat.CreatedAt); err != nil {
		return nil, err
	}
	return &chat, nil
}
