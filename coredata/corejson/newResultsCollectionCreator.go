package corejson

import (
	"gitlab.com/auk-go/core/constants"
)

type newResultsCollectionCreator struct{}

// UnmarshalUsingBytes
//
//  Aka. alias for DeserializeUsingBytes
//
//  Should be used when ResultsCollection itself is Serialized
//  and save to somewhere and then unmarshal or deserialize
func (it newResultsCollectionCreator) UnmarshalUsingBytes(
	deserializingBytes []byte,
) (*ResultsCollection, error) {
	return it.DeserializeUsingBytes(deserializingBytes)
}

// DeserializeUsingBytes
//
//  Should be used when ResultsCollection itself is Serialized
//  and save to somewhere and then unmarshal or deserialize
func (it newResultsCollectionCreator) DeserializeUsingBytes(
	deserializingBytes []byte,
) (*ResultsCollection, error) {
	empty := it.Empty()

	err := Deserialize.
		UsingBytes(deserializingBytes, empty)

	if err == nil {
		return empty, nil
	}

	return nil, err
}

func (it newResultsCollectionCreator) DeserializeUsingResult(
	jsonResult *Result,
) (*ResultsCollection, error) {
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

func (it newResultsCollectionCreator) Empty() *ResultsCollection {
	return &ResultsCollection{
		Items: []Result{},
	}
}

func (it newResultsCollectionCreator) Default() *ResultsCollection {
	list := make([]Result, 0, constants.Capacity8)

	return &ResultsCollection{
		Items: list,
	}
}

func (it newResultsCollectionCreator) UsingCap(
	capacity int,
) *ResultsCollection {
	list := make([]Result, 0, capacity)

	return &ResultsCollection{
		Items: list,
	}
}

func (it newResultsCollectionCreator) AnyItems(
	anyItems ...interface{},
) *ResultsCollection {
	list := make(
		[]Result,
		0,
		len(anyItems)+constants.Capacity5)

	collection := &ResultsCollection{
		Items: list,
	}

	return collection.AddAnyItems(
		anyItems...)
}

func (it newResultsCollectionCreator) AnyItemsPlusCap(
	addCapacity int,
	anyItems ...interface{},
) *ResultsCollection {
	length := addCapacity

	if len(anyItems) == 0 {
		return it.UsingCap(length)
	}

	additionalCapacity := len(anyItems)
	length += additionalCapacity
	list := it.UsingCap(length)

	return list.
		AddAnyItems(anyItems...)
}

func (it newResultsCollectionCreator) UsingJsonersOption(
	isIgnoreNilOrError bool,
	addCapacity int,
	jsoners ...Jsoner,
) *ResultsCollection {
	length := addCapacity
	if jsoners == nil {
		return it.UsingCap(length)
	}

	actualLength := len(jsoners)
	length += actualLength
	list := it.UsingCap(length)

	return list.
		AddJsoners(
			isIgnoreNilOrError,
			jsoners...)
}

func (it newResultsCollectionCreator) UsingJsonersNonNull(
	addCapacity int,
	jsoners ...Jsoner,
) *ResultsCollection {
	return it.UsingJsonersOption(
		true,
		addCapacity,
		jsoners...)
}

func (it newResultsCollectionCreator) UsingJsoners(
	jsoners ...Jsoner,
) *ResultsCollection {
	return it.UsingJsonersOption(
		true,
		constants.Capacity2,
		jsoners...)
}

func (it newResultsCollectionCreator) UsingResultsPtrPlusCap(
	addCapacity int,
	results ...*Result,
) *ResultsCollection {
	length := addCapacity

	if results == nil {
		return it.UsingCap(length)
	}

	actualLength := len(results)
	length += actualLength
	list := it.UsingCap(length)

	if actualLength == 0 {
		return list
	}

	return list.
		AddNonNilItemsPtr(results...)
}

func (it newResultsCollectionCreator) UsingResultsPtr(
	results ...*Result,
) *ResultsCollection {
	return it.UsingResultsPtrPlusCap(
		constants.Capacity2,
		results...)
}

func (it newResultsCollectionCreator) UsingResultsPlusCap(
	addCapacity int,
	results ...Result,
) *ResultsCollection {
	length := addCapacity

	if results == nil {
		return it.UsingCap(length)
	}

	actualLength := len(results)
	length += actualLength
	list := it.UsingCap(length)

	if actualLength == 0 {
		return list
	}

	return list.
		Adds(results...)
}

func (it newResultsCollectionCreator) UsingResults(
	results ...Result,
) *ResultsCollection {
	return it.UsingResultsPlusCap(
		constants.Capacity2,
		results...)
}

func (it newResultsCollectionCreator) Serializers(
	serializers ...bytesSerializer,
) *ResultsCollection {
	if len(serializers) == 0 {
		return it.Empty()
	}

	collection := it.UsingCap(len(serializers))

	return collection.AddSerializers(
		serializers...)
}

func (it newResultsCollectionCreator) SerializerFunctions(
	serializerFunctions ...func() ([]byte, error),
) *ResultsCollection {
	if len(serializerFunctions) == 0 {
		return it.Empty()
	}

	collection := it.UsingCap(len(serializerFunctions))

	return collection.AddSerializerFunctions(
		serializerFunctions...)
}
