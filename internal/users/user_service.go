package users

import (
	"log"
	"template-golang/internal/ent"
)

func FindById(id int) (*UserDto, error) {
	user, err := QueryById(id)
	if err != nil {
		log.Fatal("Failed to find user by ID: %w", err)
		return nil, err
	}

	return toUserDto(user), nil
}

// region - Private functions

func toUserDto(user *ent.User) *UserDto {
	return &UserDto{
		Id:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Hash:     user.Hash,
	}
}

// endregion
