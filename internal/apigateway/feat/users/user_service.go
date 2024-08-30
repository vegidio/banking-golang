package users

import (
	. "banking/internal/apigateway/shared"
	"banking/internal/apigateway/shared/errors"
)

func RetrieveById(id int) (*UserDto, *ApiError) {
	user, err := QueryById(id)
	if err != nil {
		return nil, errors.NotFound(Params{Detail: err.Error()})
	}

	return ToUserDto(user), nil
}
