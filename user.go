package messenger

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" db:"username"`
	Name     string `json:"name" db:"name"`
	Password string `json:"password" db:"password_hash"`
}
