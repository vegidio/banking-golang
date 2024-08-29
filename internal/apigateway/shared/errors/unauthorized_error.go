package errors

import (
	"banking/internal/apigateway/shared"
	"net/http"
)

func Unauthorized(params Params) *shared.HttpError {
	if params.Type == "" {
		params.Type = "UNAUTHORIZED"
	}

	return &shared.HttpError{
		Status:   http.StatusUnauthorized,
		Type:     params.Type,
		Title:    params.Title,
		Detail:   params.Detail,
		Instance: params.Instance,
	}
}
