package coreconverted

type Integers struct {
	Values        []int
	CombinedError error
}

func (it *Integers) HasError() bool {
	return it.CombinedError != nil
}

func (it *Integers) Length() int {
	if it == nil || it.Values == nil {
		return 0
	}

	return len(it.Values)
}

func (it *Integers) IsEmpty() bool {
	return it.Length() == 0
}

func (it *Integers) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *Integers) HasIssuesOrEmpty() bool {
	return it.IsEmpty() || it.CombinedError != nil
}

func (it *Integers) HandleWithPanic() {
	if it.CombinedError == nil {
		return
	}

	panic(it.CombinedError)
}
