package service

import (
	"github.com/lavatee/messenger"
	"github.com/lavatee/messenger/internal/repository"
)

type ChatsService struct {
	repo repository.Chats
}

func NewChatsService(repo repository.Chats) *ChatsService {
	return &ChatsService{
		repo: repo,
	}
}
func (s *ChatsService) CreateChat(firstUserId int, secondUserId int) (int, error) {
	return s.repo.PostChat(firstUserId, secondUserId)
}
func (s *ChatsService) GetUserChats(userId int) ([]messenger.Chat, error) {
	return s.repo.GetUserChats(userId)
}
func (s *ChatsService) DeleteChat(chatId int) error {
	return s.repo.DeleteChat(chatId)
}
