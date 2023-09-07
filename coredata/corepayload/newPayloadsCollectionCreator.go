package corepayload

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/errcore"
)

type newPayloadsCollectionCreator struct{}

func (it newPayloadsCollectionCreator) Empty() *PayloadsCollection {
	return &PayloadsCollection{
		Items: []*PayloadWrapper{},
	}
}

func (it newPayloadsCollectionCreator) Deserialize(
	rawBytes []byte,
) (*PayloadsCollection, error) {
	empty := it.Empty()

	err := corejson.
		Deserialize.
		UsingBytes(
			rawBytes, empty)

	if err != nil {
		return nil, err
	}

	return empty, nil
}

func (it newPayloadsCollectionCreator) DeserializeMust(
	rawBytes []byte,
) *PayloadsCollection {
	collection, err := it.Deserialize(rawBytes)
	errcore.HandleErr(err)

	return collection
}

func (it newPayloadsCollectionCreator) DeserializeToMany(
	rawBytes []byte,
) (payloadsSlice []*PayloadsCollection, err error) {
	err = corejson.
		Deserialize.
		UsingBytes(
			rawBytes,
			&payloadsSlice)

	if err != nil {
		return nil, err
	}

	return payloadsSlice, nil
}

func (it newPayloadsCollectionCreator) DeserializeUsingJsonResult(
	jsonResult *corejson.Result,
) (*PayloadsCollection, error) {
	empty := it.Empty()

	err := corejson.
		Deserialize.
		Apply(jsonResult, empty)

	if err != nil {
		return nil, err
	}

	return empty, nil
}

func (it newPayloadsCollectionCreator) UsingCap(
	capacity int,
) *PayloadsCollection {
	return &PayloadsCollection{
		Items: make([]*PayloadWrapper, 0, capacity),
	}
}

func (it newPayloadsCollectionCreator) UsingWrappers(
	payloadsWrappers ...*PayloadWrapper,
) *PayloadsCollection {
	if len(payloadsWrappers) == 0 {
		return it.Empty()
	}

	collection := it.UsingCap(
		len(payloadsWrappers) +
			constants.Capacity3)

	return collection.AddsPtr(payloadsWrappers...)
}
