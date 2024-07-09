package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lavatee/messenger"
)

type MessagesPostgres struct {
	db *sqlx.DB
}

func NewMessagesPostgres(db *sqlx.DB) *MessagesPostgres {
	return &MessagesPostgres{
		db: db,
	}
}
func (r *MessagesPostgres) CreateMessage(message messenger.Message) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (text, user_id, chat_id) values ($1, $2, $3) RETURNING id", messagesTable)
	row := r.db.QueryRow(query, message.Text, message.UserId, message.ChatId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (r *MessagesPostgres) GetChatMessages(firstUserId int, secondUserId int) ([]messenger.Message, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	var chatId int
	queryy := fmt.Sprintf("SELECT id FROM %s WHERE (first_user_id=$1 AND second_user_id=$2) OR (first_user_id=$3 AND second_user_id=$4)", chatsTable)
	roww := tx.QueryRow(queryy, firstUserId, secondUserId, secondUserId, firstUserId)
	if err = roww.Scan(&chatId); err != nil {
		tx.Rollback()
		return nil, err
	}

	var messages []messenger.Message
	query := fmt.Sprintf("SELECT id, user_id, text FROM %s WHERE chat_id=$1", messagesTable)
	if err = r.db.Select(&messages, query, chatId); err != nil {
		tx.Rollback()
		return nil, err
	}

	return messages, tx.Commit()
}
func (r *MessagesPostgres) DeleteMessage(messageId int) error {
	query := fmt.Sprintf("DELETE from %s WHERE id=$1", messagesTable)
	_, err := r.db.Exec(query, messageId)
	return err
}
