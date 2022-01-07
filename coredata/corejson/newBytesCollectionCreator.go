package corejson

type newBytesCollectionCreator struct{}

func (it newBytesCollectionCreator) Empty() *BytesCollection {
	return it.UsingCap(0)
}

func (it newBytesCollectionCreator) UsingCap(
	capacity int,
) *BytesCollection {
	list := make([][]byte, 0, capacity)

	return &BytesCollection{
		Items: list,
	}
}

func (it newBytesCollectionCreator) AnyItems(
	anyItems ...interface{},
) (*BytesCollection, error) {
	length := len(anyItems)
	collection := it.UsingCap(length)
	err := collection.AddAnyItems(
		anyItems...)

	return collection, err
}
