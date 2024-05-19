package service

import (
	"errors"
	"github.com/google/uuid"
	"lets-go-chat/internal/repository"
)

type UserService struct {
	repository *repository.Repository
	logPrefix  string
}

func newUserService(repository *repository.Repository) *UserService {
	return &UserService{
		repository: repository,
		logPrefix:  "[UserService] ",
	}
}

func (s *UserService) JoinChat(chatID, userID uuid.UUID) error {
	if s.repository.IsMember(chatID, userID) {
		return errors.New("current user already joined the chat")
	}
	return s.repository.JoinChat(chatID, userID)
}

func (s *UserService) LeaveChat(chatID, userID uuid.UUID) error {
	if !s.repository.IsMember(chatID, userID) {
		return errors.New("current user is not a member of the chat")
	}
	chat, err := s.repository.GetChat(chatID)
	if err != nil {
		return err
	}
	if userID == chat.OwnerID {
		return errors.New("Owner cannot leave the chat")
	}
	return s.repository.LeaveChat(chatID, userID)
}
