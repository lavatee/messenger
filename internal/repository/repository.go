package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lavatee/messenger"
)

type Auth interface {
	SignUp(username string, name string, passwordHash string) (int, error)
	SignIn(username string, passwordHash string) (messenger.User, error)
	GetUserById(id int) (string, string, error)
	PutUser(username string, name string, id int) error
}
type Messages interface {
	CreateMessage(message messenger.Message) (int, error)
	GetChatMessages(firstUserId int, secondUserId int) ([]messenger.Message, error)
	DeleteMessage(messageId int) error
}
type Chats interface {
	PostChat(firstUserId int, secondUserId int) (int, error)
	GetUserChats(userId int) ([]messenger.Chat, error)
	DeleteChat(chatId int) error
}
type Rooms interface {
	JoinRoom(userId int) (int, error)
	LeaveRoom(userId int, roomId int) (int, error)
	LeaveMatchMaking(userId int, roomId int) error
}
type Repository struct {
	Auth
	Messages
	Chats
	Rooms
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:     NewAuthPostgres(db),
		Messages: NewMessagesPostgres(db),
		Chats:    NewChatsPostgres(db),
		Rooms:    NewRoomsPostgres(db),
	}
}
