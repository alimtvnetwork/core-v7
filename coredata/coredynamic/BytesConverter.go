package coredynamic

import (
	"errors"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/errcore"
)

type BytesConverter struct {
	rawBytes []byte
}

func NewBytesConverter(
	rawBytes []byte,
) *BytesConverter {
	return &BytesConverter{
		rawBytes: rawBytes,
	}
}

func NewBytesConverterUsingJsonResult(
	jsonResult *corejson.Result,
) (*BytesConverter, error) {
	if jsonResult.HasIssuesOrEmpty() {
		return nil, jsonResult.MeaningfulError()
	}

	return NewBytesConverter(jsonResult.Bytes), nil
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

func (it BytesConverter) SafeCastString() (line string) {
	if len(it.rawBytes) == 0 {
		return ""
	}

	return string(it.rawBytes)
}

func (it BytesConverter) CastString() (line string, err error) {
	if len(it.rawBytes) == 0 {
		return "", errors.New("cast failed for empty bytes")
	}

	return string(it.rawBytes), nil
}

func (it BytesConverter) ToString() (line string, err error) {
	return corejson.
		Deserialize.
		BytesTo.
		String(it.rawBytes)
}

func (it BytesConverter) ToStringMust() (line string) {
	line, err := corejson.
		Deserialize.
		BytesTo.
		String(it.rawBytes)
	errcore.HandleErr(err)

	return line
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

func (it BytesConverter) ToInt64() (integer64 int64, err error) {
	err = corejson.
		Deserialize.
		UsingBytes(it.rawBytes, &integer64)

	return integer64, err
}

func (it BytesConverter) ToInt64Must() (integer64 int64) {
	corejson.
		Deserialize.
		UsingBytesMust(it.rawBytes, &integer64)

	return integer64
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

func (it BytesConverter) ToHashmapMust() (hashmap *corestr.Hashmap) {
	hashmap, err := it.ToHashmap()
	errcore.HandleErr(err)

	return hashmap
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

func (it BytesConverter) ToHashsetMust() (hashset *corestr.Hashset) {
	hashset, err := it.ToHashset()
	errcore.HandleErr(err)

	return hashset
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

func (it BytesConverter) ToCollectionMust() (collection *corestr.Collection) {
	collection, err := it.ToCollection()
	errcore.HandleErr(err)

	return collection
}

func (it BytesConverter) ToSimpleSlice() (simpleSlice *corestr.SimpleSlice, err error) {
	simpleSlice = corestr.Empty.SimpleSlice()

	err = corejson.
		Deserialize.
		UsingBytes(it.rawBytes, simpleSlice)

	if err == nil {
		return simpleSlice, nil
	}

	return nil, err
}

func (it BytesConverter) ToSimpleSliceMust() (simpleSlice *corestr.SimpleSlice) {
	simpleSlice, err := it.ToSimpleSlice()
	errcore.HandleErr(err)

	return simpleSlice
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

func (it BytesConverter) ToBytesCollection() (
	bytesCollection *corejson.BytesCollection, err error,
) {
	bytesCollection = corejson.Empty.BytesCollectionPtr()

	err = corejson.
		Deserialize.
		UsingBytes(it.rawBytes, bytesCollection)

	if err == nil {
		return bytesCollection, nil
	}

	return nil, err
}
