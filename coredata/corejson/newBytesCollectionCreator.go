package corejson

type newBytesCollectionCreator struct{}

// UnmarshalUsingBytes
//
//  Aka. alias for DeserializeUsingBytes
//
//  Should be used when ResultsPtrCollection itself is Serialized
//  and save to somewhere and then unmarshal or deserialize
func (it newBytesCollectionCreator) UnmarshalUsingBytes(
	deserializingBytes []byte,
) (*BytesCollection, error) {
	return it.DeserializeUsingBytes(deserializingBytes)
}

// DeserializeUsingBytes
//
//  Should be used when BytesCollection itself is Serialized
//  and save to somewhere and then unmarshal or deserialize
func (it newBytesCollectionCreator) DeserializeUsingBytes(
	deserializingBytes []byte,
) (*BytesCollection, error) {
	empty := it.Empty()

	err := Deserialize.
		UsingBytes(deserializingBytes, empty)

	if err == nil {
		return empty, nil
	}

	return nil, err
}

func (it newBytesCollectionCreator) DeserializeUsingResult(
	jsonResult *Result,
) (*BytesCollection, error) {
	if jsonResult.HasIssuesOrEmpty() {
		return nil, jsonResult.MeaningfulError()
	}

	empty := it.Empty()

	err := Deserialize.
		UsingBytes(jsonResult.SafeBytes(), empty)

	if err == nil {
		return empty, nil
	}

	return nil, err
}

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
