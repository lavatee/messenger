package messenger

type Chat struct {
	Id           int `json:"user_id" db:"id"`
	FirstUserId  int `json:"first_user_id" db:"first_user_id"`
	SecondUserId int `json:"second_user_id" db:"second_user_id"`
}
