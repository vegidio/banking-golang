package shared

import "fmt"

type Response[T any] struct {
	Data  T     `json:"data,omitempty"`
	Error error `json:"error,omitempty"`
}

type ApiError struct {
	Status   int    `json:"status"`
	Type     string `json:"type"`
	Title    string `json:"title,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("ApiError[Status=%d, Type=%s, Title=%s, Detail=%s, Instance=%s]",
		e.Status, e.Type, e.Title, e.Detail, e.Instance)
}
