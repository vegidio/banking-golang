package errors

import (
	. "banking/internal/apigateway/shared"
	"net/http"
)

func NotFound(params Params) *ApiError {
	if params.Type == "" {
		params.Type = "NOT_FOUND"
	}

	return &ApiError{
		Status:   http.StatusNotFound,
		Type:     params.Type,
		Title:    params.Title,
		Detail:   params.Detail,
		Instance: params.Instance,
	}
}
