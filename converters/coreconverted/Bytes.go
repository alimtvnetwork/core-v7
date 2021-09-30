package coreconverted

type Bytes struct {
	Values        []byte
	CombinedError error
}

func (it *Bytes) HasError() bool {
	return it.CombinedError != nil
}

func (it *Bytes) Length() int {
	if it == nil || it.Values == nil {
		return 0
	}

	return len(it.Values)
}

func (it *Bytes) IsEmpty() bool {
	return it.Length() == 0
}

func (it *Bytes) HandleWithPanic() {
	if it.CombinedError == nil {
		return
	}

	panic(it.CombinedError)
}
