package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lavatee/messenger"
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
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var firstUserName string
	queryy := fmt.Sprintf("SELECT name FROM %s WHERE id=$1", usersTable)
	roww := tx.QueryRow(queryy, firstUserId)
	err = roww.Scan(&firstUserName)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	var secondUserName string
	queryyy := fmt.Sprintf("SELECT name FROM %s WHERE id=$1", usersTable)
	rowww := tx.QueryRow(queryyy, secondUserId)
	err = rowww.Scan(&secondUserName)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	var chatId int
	query := fmt.Sprintf("INSERT INTO %s (first_user_id, first_user_name, second_user_id, second_user_name) values ($1, $2, $3, $4) RETURNING id", chatsTable)
	row := r.db.QueryRow(query, firstUserId, firstUserName, secondUserId, secondUserName)
	err = row.Scan(&chatId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return chatId, tx.Commit()
}
func (r *ChatsPostgres) GetUserChats(userId int) ([]messenger.Chat, error) {
	var input []messenger.Chat
	query := fmt.Sprintf("SELECT id, first_user_id, first_user_name, second_user_id, second_user_name FROM %s WHERE first_user_id=$1 OR second_user_id=$2", chatsTable)
	err := r.db.Select(&input, query, userId, userId)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (r *ChatsPostgres) DeleteChat(chatId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", chatsTable)
	_, err := r.db.Exec(query, chatId)
	return err
}
