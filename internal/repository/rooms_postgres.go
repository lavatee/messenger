package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type RoomsPostgres struct {
	db *sqlx.DB
}

func NewRoomsPostgres(db *sqlx.DB) *RoomsPostgres {
	return &RoomsPostgres{
		db: db,
	}
}
func (r *RoomsPostgres) JoinRoom(userId int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var availableRoom int
	query := fmt.Sprintf("SELECT id FROM %s WHERE users_quantity=$1", roomsTable)
	row := tx.QueryRow(query, 1)
	if err := row.Scan(&availableRoom); err != nil {
		availableRoom = 0
	}
	if availableRoom == 0 {
		var roomId int
		query := fmt.Sprintf("INSERT INTO %s (first_user_id, second_user_id, users_quantity) values ($1, $2, $3) RETURNING id", roomsTable)
		row := tx.QueryRow(query, userId, 0, 1)
		if err := row.Scan(&roomId); err != nil {
			tx.Rollback()
			return 0, err
		}
		return roomId, tx.Commit()
	}
	queryy := fmt.Sprintf("UPDATE %s SET users_quantity = $1, first_user_id = CASE WHEN id = $2 AND first_user_id = $3 THEN $4 ELSE first_user_id END, second_user_id = CASE WHEN id = $5 AND second_user_id = $6 THEN $7 ELSE second_user_id END WHERE id = $8", roomsTable)
	_, err = tx.Exec(queryy, 2, availableRoom, 0, userId, availableRoom, 0, userId, availableRoom)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return availableRoom, tx.Commit()
}
func (r *RoomsPostgres) LeaveRoom(userId int, roomId int) (int, error) {
	currentRoomId, err := r.JoinRoom(userId)
	if err != nil {
		return 0, err
	}
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query := fmt.Sprintf("UPDATE %s SET users_quantity = $1, first_user_id = CASE WHEN id = $2 AND first_user_id = $3 THEN $4 ELSE first_user_id END, second_user_id = CASE WHEN id = $5 AND second_user_id = $6 THEN $7 ELSE second_user_id END WHERE id = $8", roomsTable)
	_, err = tx.Exec(query, 1, roomId, userId, 0, roomId, userId, 0, roomId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	queryy := fmt.Sprintf("DELETE FROM %s WHERE first_user_id=$1 AND second_user_id=$2", roomsTable)
	_, err = tx.Exec(queryy, 0, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	return currentRoomId, tx.Commit()

}
func (r *RoomsPostgres) LeaveMatchMaking(userId int, roomId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("UPDATE %s SET users_quantity = $1, first_user_id = CASE WHEN id = $2 AND first_user_id = $3 THEN $4 ELSE first_user_id END, second_user_id = CASE WHEN id = $5 AND second_user_id = $6 THEN $7 ELSE second_user_id END WHERE id = $8", roomsTable)
	_, err = tx.Exec(query, 1, roomId, userId, 0, roomId, userId, 0, roomId)
	if err != nil {
		tx.Rollback()
		return err
	}
	queryy := fmt.Sprintf("DELETE FROM %s WHERE first_user_id=$1 AND second_user_id=$2", roomsTable)
	_, err = tx.Exec(queryy, 0, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	return tx.Commit()

}
