package loader

type Loader[T any] struct {
	loader func() ([]T, error)
}

func NewLoader[T any](loader func() ([]T, error)) ILoader[T] {
	return &Loader[T]{
		loader: loader,
	}
}

func (l *Loader[T]) Load() ([]T, error) {
	return l.loader()
}
