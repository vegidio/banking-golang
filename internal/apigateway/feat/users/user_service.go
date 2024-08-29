package users

import (
	"log"
)

func RetrieveById(id int) (*UserDto, error) {
	user, err := QueryById(id)
	if err != nil {
		log.Fatal("Failed to find user by ID: %w", err)
		return nil, err
	}

	return ToUserDto(user), nil
}
