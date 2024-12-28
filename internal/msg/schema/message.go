package schema

type Message[T any] struct {
	Operation string `json:"operation"`
	Data T `json:"data"`
}