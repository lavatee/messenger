package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable    = "users"
	roomsTable    = "rooms"
	chatsTable    = "chats"
	messagesTable = "messages"
)

type Config struct {
	Port     string
	Host     string
	Username string
	Password string
	DBName   string
	SSLmode  string
}

func NewPostgresDB(c Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.Host, c.Port, c.Username, c.Password, c.DBName, c.SSLmode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
