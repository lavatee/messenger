package messenger

type Chat struct {
	Id             int    `json:"id" db:"id"`
	FirstUserId    int    `json:"first_user_id" db:"first_user_id"`
	FirstUserName  string `json:"first_user_name" db:"first_user_name"`
	SecondUserId   int    `json:"second_user_id" db:"second_user_id"`
	SecondUserName string `json:"second_user_name" db:"second_user_name"`
}
