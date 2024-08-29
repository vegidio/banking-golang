package shared

import "fmt"

type Response[T any] struct {
	Data  T     `json:"data,omitempty"`
	Error error `json:"error,omitempty"`
}

type HttpError struct {
	Status   int    `json:"status"`
	Type     string `json:"type"`
	Title    string `json:"title,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("Status %d: %s", e.Status, e.Type)
}
