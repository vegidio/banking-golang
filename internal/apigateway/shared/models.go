package shared

type Response[T any] struct {
	Data  T     `json:"data,omitempty"`
	Error error `json:"error,omitempty"`
}
