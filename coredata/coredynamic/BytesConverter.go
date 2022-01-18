package coredynamic

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
)

type BytesConverter struct {
	rawBytes []byte
}

func NewBytesConverter(rawBytes []byte) *BytesConverter {
	return &BytesConverter{
		rawBytes: rawBytes,
	}
}

func (it BytesConverter) Deserialize(
	deserializePointer interface{},
) error {
	return corejson.
		Deserialize.
		UsingBytes(it.rawBytes, deserializePointer)
}

func (it BytesConverter) DeserializeMust(
	deserializePointer interface{},
) {
	corejson.
		Deserialize.
		UsingBytesMust(it.rawBytes, deserializePointer)
}

func (it BytesConverter) ToBool() (isResult bool, err error) {
	return corejson.
		Deserialize.
		BytesTo.
		Bool(it.rawBytes)
}

func (it BytesConverter) ToBoolMust() (isResult bool) {
	return corejson.
		Deserialize.
		BytesTo.
		BoolMust(it.rawBytes)
}

func (it BytesConverter) ToStrings() (lines []string, err error) {
	return corejson.
		Deserialize.
		BytesTo.
		Strings(it.rawBytes)
}

func (it BytesConverter) ToStringsMust() (lines []string) {
	return corejson.
		Deserialize.
		BytesTo.
		StringsMust(it.rawBytes)
}

func (it BytesConverter) ToInt64() (isResult bool, err error) {
	err = corejson.
		Deserialize.
		UsingBytes(it.rawBytes, &isResult)

	return isResult, err
}

func (it BytesConverter) ToHashmap() (hashmap *corestr.Hashmap, err error) {
	hashmap = corestr.Empty.Hashmap()

	err = corejson.
		Deserialize.
		UsingBytes(it.rawBytes, hashmap)

	if err == nil {
		return hashmap, nil
	}

	return nil, err
}

func (it BytesConverter) ToHashset() (hashset *corestr.Hashset, err error) {
	hashset = corestr.Empty.Hashset()

	err = corejson.
		Deserialize.
		UsingBytes(it.rawBytes, hashset)

	if err == nil {
		return hashset, nil
	}

	return nil, err
}

func (it BytesConverter) ToCollection() (collection *corestr.Collection, err error) {
	collection = corestr.Empty.Collection()

	err = corejson.
		Deserialize.
		UsingBytes(it.rawBytes, collection)

	if err == nil {
		return collection, nil
	}

	return nil, err
}

func (it BytesConverter) ToKeyValCollection() (keyValCollection *KeyValCollection, err error) {
	keyValCollection = EmptyKeyValCollection()

	err = corejson.
		Deserialize.
		UsingBytes(it.rawBytes, keyValCollection)

	if err == nil {
		return keyValCollection, nil
	}

	return nil, err
}

func (it BytesConverter) ToAnyCollection() (anyCollection *AnyCollection, err error) {
	anyCollection = EmptyAnyCollection()

	err = corejson.
		Deserialize.
		UsingBytes(it.rawBytes, anyCollection)

	if err == nil {
		return anyCollection, nil
	}

	return nil, err
}

func (it BytesConverter) ToMapAnyItems() (mapAnyItems *MapAnyItems, err error) {
	mapAnyItems = EmptyMapAnyItems()

	err = corejson.
		Deserialize.
		UsingBytes(it.rawBytes, mapAnyItems)

	if err == nil {
		return mapAnyItems, nil
	}

	return nil, err
}

func (it BytesConverter) ToDynamicCollection() (dynamicCollection *DynamicCollection, err error) {
	dynamicCollection = EmptyDynamicCollection()

	err = corejson.
		Deserialize.
		UsingBytes(it.rawBytes, dynamicCollection)

	if err == nil {
		return dynamicCollection, nil
	}

	return nil, err
}

func (it BytesConverter) ToJsonResultCollection() (
	jsonResultCollection *corejson.ResultsCollection, err error,
) {
	jsonResultCollection = corejson.Empty.ResultsCollection()

	err = corejson.
		Deserialize.
		UsingBytes(it.rawBytes, jsonResultCollection)

	if err == nil {
		return jsonResultCollection, nil
	}

	return nil, err
}

func (it BytesConverter) ToJsonMapResults() (
	jsonMapResults *corejson.MapResults, err error,
) {
	jsonMapResults = corejson.Empty.MapResults()

	err = corejson.
		Deserialize.
		UsingBytes(it.rawBytes, jsonMapResults)

	if err == nil {
		return jsonMapResults, nil
	}

	return nil, err
}
