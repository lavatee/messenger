package service

import (
	"github.com/lavatee/messenger"
	"github.com/lavatee/messenger/internal/repository"
)

type MessagesService struct {
	repo repository.Messages
}

func NewMessagesService(repo repository.Messages) *MessagesService {
	return &MessagesService{
		repo: repo,
	}
}
func (s *MessagesService) CreateMessage(message messenger.Message) (int, error) {
	return s.repo.CreateMessage(message)
}
func (s *MessagesService) GetChatMessages(firstUserId int, secondUserId int) ([]messenger.Message, error) {
	return s.repo.GetChatMessages(firstUserId, secondUserId)
}
func (s *MessagesService) DeleteMessage(messageId int) error {
	return s.repo.DeleteMessage(messageId)
}
