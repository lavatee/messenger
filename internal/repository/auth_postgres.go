package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lavatee/messenger"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}
func (r *AuthPostgres) SignUp(username string, name string, passwordHash string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, name, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, username, name, passwordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (r *AuthPostgres) SignIn(username string, passwordHash string) (messenger.User, error) {
	var input messenger.User
	query := fmt.Sprintf("SELECT id from %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&input, query, username, passwordHash)
	return input, err
}
func (r *AuthPostgres) GetUserById(id int) (string, string, error) {
	var input messenger.User
	query := fmt.Sprintf("SELECT username, name from %s WHERE id=$1", usersTable)
	err := r.db.Get(&input, query, id)
	return input.Username, input.Name, err
}
