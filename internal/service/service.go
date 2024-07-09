package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/lavatee/messenger"
	"github.com/lavatee/messenger/internal/repository"
)

type Auth interface {
	SignUp(username string, name string, password string) (int, error)
	HashPassword(password string) string
	SignIn(username string, password string) (int, string, string, string, error)
	Refresh(token string) (string, string, error)
	newToken(tokenStruct jwt.MapClaims) (string, error)
	GetUserById(id int) (string, string, error)
	PutUser(username string, name string, id int) error
}
type Messages interface {
	CreateMessage(message messenger.Message) (int, error)
	GetChatMessages(firstUserId int, secondUserId int) ([]messenger.Message, error)
	DeleteMessage(messageId int) error
}
type Chats interface {
	CreateChat(firstUserId int, secondUserId int) (int, error)
	GetUserChats(userId int) ([]messenger.Chat, error)
	DeleteChat(chatId int) error
}
type Rooms interface {
	JoinRoom(userId int) (int, error)
	LeaveRoom(userId int, roomId int) (int, error)
	LeaveMatchMaking(userId int, roomId int) error
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
