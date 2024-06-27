package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ChatsPostgres struct {
	db *sqlx.DB
}

func NewChatsPostgres(db *sqlx.DB) *ChatsPostgres {
	return &ChatsPostgres{
		db: db,
	}
}

func (r *ChatsPostgres) PostChat(firstUserId int, secondUserId int) (int, error) {
	var chatId int
	query := fmt.Sprintf("INSERT INTO %s (first_user_id, second_user_id) values ($1, $2) RETURNING id", chatsTable)
	row := r.db.QueryRow(query, firstUserId, secondUserId)
	err := row.Scan(&chatId)
	if err != nil {
		return 0, err
	}
	return chatId, nil
}
