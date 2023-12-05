package getter

import "github.com/t-kuni/go-lazy-load-design-pattern/src/loader"

type ArrayIndexer[In any, Key any, Out any] struct {
	source    loader.ILoader[In]
	index     map[any][]Out
	indexer   func(In) (Key, Out, error)
	isIndexed bool
}

func NewArrayIndexer[In any, Key any, Out any](source loader.ILoader[In], indexer func(In) (Key, Out, error)) *ArrayIndexer[In, Key, Out] {
	return &ArrayIndexer[In, Key, Out]{
		source:    source,
		index:     make(map[any][]Out),
		indexer:   indexer,
		isIndexed: false,
	}
}

func (i *ArrayIndexer[In, Key, Out]) Get(key Key) ([]Out, bool, error) {
	if !i.isIndexed {
		_, err := i.load()
		if err != nil {
			return *new([]Out), false, err
		}
	}

	item, ok := i.index[key]
	return item, ok, nil
}

func (i *ArrayIndexer[In, Key, Out]) load() ([]In, error) {
	arr, err := i.source.Load()
	if err != nil {
		return nil, err
	}

	for _, item := range arr {
		key, out, err := i.indexer(item)
		if err != nil {
			return nil, err
		}
		_, ok := i.index[key]
		if !ok {
			i.index[key] = []Out{}
		}
		i.index[key] = append(i.index[key], out)
	}

	i.isIndexed = true
	return arr, nil
}
