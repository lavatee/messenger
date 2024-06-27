package repository

import "github.com/jmoiron/sqlx"

type MessagesPostgres struct {
	db *sqlx.DB
}

func NewMessagesPostgres(db *sqlx.DB) *MessagesPostgres {
	return &MessagesPostgres{
		db: db,
	}
}
