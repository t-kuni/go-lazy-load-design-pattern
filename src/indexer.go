package src

type Indexer[In any, Key any, Out any] struct {
	parent    ILoader[In]
	index     map[any]Out
	indexer   func(In) (Key, Out, error)
	isIndexed bool
}

// TODO indexerの引数でOutを受け取る？
func NewIndexer[In any, Key any, Out any](source ILoader[In], indexer func(In) (Key, Out, error)) *Indexer[In, Key, Out] {
	return &Indexer[In, Key, Out]{
		parent:    source,
		index:     make(map[any]Out),
		indexer:   indexer,
		isIndexed: false,
	}
}

func (i *Indexer[In, Key, Out]) Load() ([]In, error) {
	arr, err := i.parent.Load()
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

func (i *Indexer[In, Key, Out]) Get(key Key) (Out, bool, error) {
	if !i.isIndexed {
		_, err := i.Load()
		if err != nil {
			return *new(Out), false, err
		}
	}

	item, ok := i.index[key]
	return item, ok, nil
}
