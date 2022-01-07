package corejson

type emptyCreator struct{}

func (it emptyCreator) Result() Result {
	return Result{}
}

func (it emptyCreator) ResultWithErr(
	typeName string,
	err error,
) Result {
	return Result{
		Error:    err,
		TypeName: typeName,
	}
}

func (it emptyCreator) ResultPtrWithErr(
	typeName string,
	err error,
) *Result {
	return &Result{
		Error:    err,
		TypeName: typeName,
	}
}

func (it emptyCreator) ResultPtr() *Result {
	return &Result{}
}

func (it emptyCreator) BytesCollection() BytesCollection {
	return BytesCollection{}
}

func (it emptyCreator) BytesCollectionPtr() *BytesCollection {
	return &BytesCollection{}
}

func (it emptyCreator) ResultsCollection() *ResultsCollection {
	return &ResultsCollection{
		Items: []Result{},
	}
}

func (it emptyCreator) ResultsPtrCollection() *ResultsPtrCollection {
	return &ResultsPtrCollection{
		Items: []*Result{},
	}
}

func (it emptyCreator) MapResults() *MapResults {
	return &MapResults{
		Items: map[string]Result{},
	}
}
