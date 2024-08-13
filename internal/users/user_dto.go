package users

type UserDto struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Hash     string
}
