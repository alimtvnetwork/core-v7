package corejson

import (
	"gitlab.com/evatix-go/core/constants"
)

type newResultsPtrCollectionCreator struct{}

// UnmarshalUsingBytes
//
//  Aka. alias for DeserializeUsingBytes
//
//  Should be used when ResultsPtrCollection itself is Serialized
//  and save to somewhere and then unmarshal or deserialize
func (it newResultsPtrCollectionCreator) UnmarshalUsingBytes(
	deserializingBytes []byte,
) (*ResultsPtrCollection, error) {
	return it.DeserializeUsingBytes(deserializingBytes)
}

// DeserializeUsingBytes
//
//  Should be used when ResultsPtrCollection itself is Serialized
//  and save to somewhere and then unmarshal or deserialize
func (it newResultsPtrCollectionCreator) DeserializeUsingBytes(
	deserializingBytes []byte,
) (*ResultsPtrCollection, error) {
	empty := it.Empty()

	err := Deserialize.
		UsingBytes(deserializingBytes, empty)

	if err == nil {
		return empty, nil
	}

	return nil, err
}

func (it newResultsPtrCollectionCreator) DeserializeUsingResult(
	jsonResult *Result,
) (*ResultsPtrCollection, error) {
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

func (it newResultsPtrCollectionCreator) Empty() *ResultsPtrCollection {
	list := make([]*Result, 0)

	return &ResultsPtrCollection{
		Items: list,
	}
}

func (it newResultsPtrCollectionCreator) UsingCap(
	capacity int,
) *ResultsPtrCollection {
	list := make([]*Result, 0, capacity)

	return &ResultsPtrCollection{
		Items: list,
	}
}

func (it newResultsPtrCollectionCreator) Default() *ResultsPtrCollection {
	list := make([]*Result, 0, constants.Capacity8)

	return &ResultsPtrCollection{
		Items: list,
	}
}

func (it newResultsPtrCollectionCreator) AnyItemsPlusCap(
	capacity int,
	anyItems ...interface{},
) *ResultsPtrCollection {
	length := capacity + len(anyItems)

	if length == 0 || len(anyItems) == 0 {
		return it.UsingCap(length)
	}

	collection := it.UsingCap(length)

	return collection.AddAnyItems(
		anyItems...)
}

func (it newResultsPtrCollectionCreator) AnyItems(
	anyItems ...interface{},
) *ResultsPtrCollection {
	return it.AnyItemsPlusCap(
		0,
		anyItems...)
}

func (it newResultsPtrCollectionCreator) UsingResultsPlusCap(
	addCapacity int,
	results ...*Result,
) *ResultsPtrCollection {
	length := addCapacity + len(results)

	if length == 0 || len(results) == 0 {
		return it.UsingCap(length)
	}

	list := it.UsingCap(length)

	return list.
		AddNonNilItemsPtr(results...)
}

func (it newResultsPtrCollectionCreator) UsingResults(
	results ...*Result,
) *ResultsPtrCollection {
	return it.UsingResultsPlusCap(
		0,
		results...)
}

func (it newResultsPtrCollectionCreator) JsonersPlusCap(
	isIgnoreNilOrErr bool,
	capacity int,
	jsoners ...Jsoner,
) *ResultsPtrCollection {
	length := capacity + len(jsoners)

	if length == 0 || len(jsoners) == 0 {
		return it.UsingCap(length)
	}

	collection := it.UsingCap(length)

	return collection.AddJsoners(
		isIgnoreNilOrErr,
		jsoners...)
}

func (it newResultsPtrCollectionCreator) Jsoners(
	jsoners ...Jsoner,
) *ResultsPtrCollection {
	return it.JsonersPlusCap(
		true,
		0,
		jsoners...)
}

func (it newResultsPtrCollectionCreator) Serializers(
	serializers ...bytesSerializer,
) *ResultsPtrCollection {
	if len(serializers) == 0 {
		return it.Empty()
	}

	collection := it.UsingCap(len(serializers))

	for _, serializer := range serializers {
		collection.AddSerializer(serializer)
	}

	return collection
}
