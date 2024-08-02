package helper

func New[T any](v T) *T {
	return &v
}
