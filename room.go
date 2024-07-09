package messenger

type Room struct {
	Id            int `json:"id" db:"id"`
	FirstUserId   int `json:"first_user_id" db:"first_user_id"`
	SecondUserId  int `json:"second_user_id" db:"second_user_id"`
	UsersQuantity int `json:"users_quantity" db:"users_quantity"`
}
