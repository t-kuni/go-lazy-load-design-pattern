package getter

type Converter[In any, Key any, Out any] struct {
	source    IGetter[Key, In]
	converter func(key Key, v In) (Out, error)
}

func NewConverter[In any, Key any, Out any](source IGetter[Key, In], converter func(key Key, v In) (Out, error)) *Converter[In, Key, Out] {
	return &Converter[In, Key, Out]{
		source:    source,
		converter: converter,
	}
}

func (i *Converter[In, Key, Out]) Get(key Key) (Out, bool, error) {
	v, ok, err := i.source.Get(key)
	if err != nil || !ok {
		return *new(Out), ok, err
	}
	newVal, err := i.converter(key, v)
	if err != nil {
		return *new(Out), false, err
	}
	return newVal, true, nil
}
