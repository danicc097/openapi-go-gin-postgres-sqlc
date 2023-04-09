package pointers

func New[T any](v T) *T {
	return &v
}
