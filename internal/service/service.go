package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/lavatee/messenger"
	"github.com/lavatee/messenger/internal/repository"
)

type Auth interface {
	SignUp(username string, name string, password string) (int, error)
	HashPassword(password string) string
	SignIn(username string, password string) (string, string, error)
	Refresh(token string) (string, string, error)
	newToken(tokenStruct jwt.MapClaims) (string, error)
	GetUserById(id int) (string, string, error)
}
type Messages interface {
}
type Chats interface {
	CreateChat(firstUserId int, secondUserId int) (int, error)
	GetUserChats(userId int) ([]messenger.Chat, error)
	DeleteChat(chatId int) error
}
type Rooms interface {
}
type Service struct {
	Auth
	Messages
	Chats
	Rooms
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth:     NewAuthService(repos.Auth),
		Messages: NewMessagesService(repos.Messages),
		Chats:    NewChatsService(repos.Chats),
		Rooms:    NewRoomsService(repos.Rooms),
	}
}
