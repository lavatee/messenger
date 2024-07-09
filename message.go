package messenger

type Message struct {
	Id     int    `json:"-" db:"id"`
	ChatId int    `json:"chat_id" db:"chat_id"`
	UserId int    `json:"user_id" db:"user_id"`
	Text   string `json:"text" db:"text"`
}
