package users

import (
	"banking/internal/apigateway/ent"
)

type UserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Hash     string `json:"-"`
}

func ToUserDto(user *ent.User) *UserDto {
	return &UserDto{
		Id:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Hash:     user.Hash,
	}
}
