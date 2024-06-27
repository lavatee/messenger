package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lavatee/messenger"
)

type Auth interface {
	SignUp(username string, name string, passwordHash string) (int, error)
	SignIn(username string, passwordHash string) (messenger.User, error)
	GetUserById(id int) (string, string, error)
}
type Messages interface {
}
type Chats interface {
	PostChat(firstUserId int, secondUserId int) (int, error)
}
type Rooms interface {
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
