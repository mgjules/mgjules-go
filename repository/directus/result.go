package directus

type Result[T any] struct {
	Data T `json:"data"`
}
