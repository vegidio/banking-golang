package users

func FetchUserById(id string) UserDto {
	user := UserDto{id, "Vinicius", "vegidio", "vegidio@gmail.com", "123456"}
	return user
}
