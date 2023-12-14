package codestack

type newTraceCollection struct{}

func (it newTraceCollection) Cap(capacity int) *TraceCollection {
	slice := make([]Trace, 0, capacity)

	return &TraceCollection{
		slice,
	}
}

func (it newTraceCollection) Default() *TraceCollection {
	return it.Cap(DefaultStackCount + 5)
}

func (it newTraceCollection) Using(
	isClone bool,
	traces ...Trace,
) *TraceCollection {
	if traces == nil {
		return it.Empty()
	}

	if !isClone {
		return &TraceCollection{
			traces,
		}
	}

	slice := it.Cap(len(traces))

	return slice.Adds(traces...)
}

func (it newTraceCollection) Empty() *TraceCollection {
	return it.Cap(0)
}
