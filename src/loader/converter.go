package loader

type Converter[In any, Out any] struct {
	source    ILoader[In]
	converter func([]In) ([]Out, error)
}

func NewConverter[In any, Out any](source ILoader[In], converter func([]In) ([]Out, error)) ILoader[Out] {
	return &Converter[In, Out]{
		source:    source,
		converter: converter,
	}
}

func (h *Converter[In, Out]) Load() ([]Out, error) {
	arr, err := h.source.Load()
	if err != nil {
		return nil, err
	}
	return h.converter(arr)
}
