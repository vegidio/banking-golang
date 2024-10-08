package errors

import (
	. "banking/internal/apigateway/shared"
	"net/http"
)

func Unauthorized(params Params) *ApiError {
	if params.Type == "" {
		params.Type = "UNAUTHORIZED"
	}

	return &ApiError{
		Status:   http.StatusUnauthorized,
		Type:     params.Type,
		Title:    params.Title,
		Detail:   params.Detail,
		Instance: params.Instance,
	}
}
