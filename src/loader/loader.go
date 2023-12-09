package loader

type Loader[T any] struct {
	loader func() ([]T, error)
}

type LoaderOption struct {
	Cache bool
}

var defaultLoaderOption = &LoaderOption{
	Cache: true,
}

func NewLoader[T any](loader func() ([]T, error), opt *LoaderOption) ILoader[T] {
	l := &Loader[T]{
		loader: loader,
	}

	if opt == nil {
		opt = defaultLoaderOption
	}

	if opt.Cache {
		return NewCacheHolder[T](l)
	}

	return l
}

func (l *Loader[T]) Load() ([]T, error) {
	return l.loader()
}
