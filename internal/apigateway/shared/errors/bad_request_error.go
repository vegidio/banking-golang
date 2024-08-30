package errors

import (
	. "banking/internal/apigateway/shared"
	"net/http"
)

func BadRequest(params Params) *ApiError {
	if params.Type == "" {
		params.Type = "BAD_REQUEST"
	}

	return &ApiError{
		Status:   http.StatusBadRequest,
		Type:     params.Type,
		Title:    params.Title,
		Detail:   params.Detail,
		Instance: params.Instance,
	}
}
