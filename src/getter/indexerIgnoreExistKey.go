package getter

import "github.com/t-kuni/go-lazy-load-design-pattern/src/loader"

type IndexerIgnoreExistKey[In any, Key any, Out any] struct {
	source       loader.ILoader[In]
	index        map[any]Out
	provideKey   func(In) (Key, error)
	provideValue func(In) (Out, error)
	isIndexed    bool
}

func NewIndexerIgnoreExistKey[In any, Key any, Out any](source loader.ILoader[In], provideKey func(In) (Key, error), provideValue func(In) (Out, error)) IGetter[Key, Out] {
	return &IndexerIgnoreExistKey[In, Key, Out]{
		source:       source,
		index:        make(map[any]Out),
		provideKey:   provideKey,
		provideValue: provideValue,
		isIndexed:    false,
	}
}

func (i *IndexerIgnoreExistKey[In, Key, Out]) Get(key Key) (Out, bool, error) {
	if !i.isIndexed {
		_, err := i.load()
		if err != nil {
			return *new(Out), false, err
		}
	}

	item, ok := i.index[key]
	return item, ok, nil
}

func (i *IndexerIgnoreExistKey[In, Key, Out]) load() ([]In, error) {
	arr, err := i.source.Load()
	if err != nil {
		return nil, err
	}

	for _, item := range arr {
		key, err := i.provideKey(item)
		if err != nil {
			return nil, err
		}
		_, ok := i.index[key]
		if ok {
			continue
		}
		val, err := i.provideValue(item)
		if err != nil {
			return nil, err
		}
		i.index[key] = val
	}

	i.isIndexed = true
	return arr, nil
}
