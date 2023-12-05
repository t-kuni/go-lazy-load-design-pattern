package src

type CacheHolder[T any] struct {
	arr    []T
	source ILoader[T]
}

func NewCacheHolder[T any](source ILoader[T]) *CacheHolder[T] {
	return &CacheHolder[T]{
		source: source,
		arr:    nil,
	}
}

func (h *CacheHolder[T]) Load() ([]T, error) {
	if h.arr == nil {
		arr, err := h.source.Load()
		if err != nil {
			return nil, err
		}
		h.arr = arr
	}
	return h.arr, nil
}
