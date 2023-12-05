package getter

import "github.com/t-kuni/go-lazy-load-design-pattern/src/loader"

type Indexer[In any, Key any, Out any] struct {
	source    loader.ILoader[In]
	index     map[any]Out
	indexer   func(In) (Key, Out, error)
	isIndexed bool
}

func NewIndexer[In any, Key any, Out any](source loader.ILoader[In], indexer func(In) (Key, Out, error)) *Indexer[In, Key, Out] {
	return &Indexer[In, Key, Out]{
		source:    source,
		index:     make(map[any]Out),
		indexer:   indexer,
		isIndexed: false,
	}
}

func (i *Indexer[In, Key, Out]) Get(key Key) (Out, bool, error) {
	if !i.isIndexed {
		_, err := i.load()
		if err != nil {
			return *new(Out), false, err
		}
	}

	item, ok := i.index[key]
	return item, ok, nil
}

func (i *Indexer[In, Key, Out]) load() ([]In, error) {
	arr, err := i.source.Load()
	if err != nil {
		return nil, err
	}

	for _, item := range arr {
		key, out, err := i.indexer(item)
		if err != nil {
			return nil, err
		}
		i.index[key] = out
	}

	i.isIndexed = true
	return arr, nil
}
