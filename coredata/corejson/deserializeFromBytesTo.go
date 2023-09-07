package corejson

import "gitlab.com/auk-go/core/errcore"

type deserializeFromBytesTo struct{}

func (it deserializeFromBytesTo) Strings(
	rawBytes []byte,
) (lines []string, err error) {
	err = Deserialize.
		UsingBytes(rawBytes, &lines)

	return lines, err
}

func (it deserializeFromBytesTo) StringsMust(
	rawBytes []byte,
) (lines []string) {
	err := Deserialize.UsingBytes(
		rawBytes,
		&lines)

	if err != nil {
		panic(err)
	}

	return lines
}

func (it deserializeFromBytesTo) String(
	rawBytes []byte,
) (line string, err error) {
	err = Deserialize.UsingBytes(
		rawBytes,
		&line)

	return line, err
}

func (it deserializeFromBytesTo) Integer(
	rawBytes []byte,
) (integer int, err error) {
	err = Deserialize.UsingBytes(
		rawBytes,
		&integer)

	return integer, err
}

func (it deserializeFromBytesTo) IntegerMust(
	rawBytes []byte,
) (integer int) {
	err := Deserialize.UsingBytes(
		rawBytes,
		&integer)

	errcore.HandleErr(err)

	return integer
}

func (it deserializeFromBytesTo) Integer64(
	rawBytes []byte,
) (integer64 int64, err error) {
	err = Deserialize.UsingBytes(
		rawBytes,
		&integer64)

	return integer64, err
}

func (it deserializeFromBytesTo) Integer64Must(
	rawBytes []byte,
) (integer64 int64) {
	err := Deserialize.UsingBytes(
		rawBytes,
		&integer64)

	errcore.HandleErr(err)

	return integer64
}

func (it deserializeFromBytesTo) Integers(
	rawBytes []byte,
) (integers []int, err error) {
	err = Deserialize.UsingBytes(
		rawBytes,
		&integers)

	return integers, err
}

func (it deserializeFromBytesTo) IntegersMust(
	rawBytes []byte,
) (integers []int) {
	err := Deserialize.UsingBytes(
		rawBytes,
		&integers)

	if err != nil {
		panic(err)
	}

	return integers
}

func (it deserializeFromBytesTo) StringMust(
	rawBytes []byte,
) (line string) {
	err := Deserialize.UsingBytes(rawBytes, &line)

	if err != nil {
		panic(err)
	}

	return line
}

func (it deserializeFromBytesTo) MapAnyItem(
	rawBytes []byte,
) (mapAnyItem map[string]interface{}, err error) {
	err = Deserialize.UsingBytes(rawBytes, &mapAnyItem)

	return mapAnyItem, err
}

func (it deserializeFromBytesTo) MapAnyItemMust(
	rawBytes []byte,
) (mapAnyItem map[string]interface{}) {
	err := Deserialize.UsingBytes(
		rawBytes,
		&mapAnyItem)

	errcore.HandleErr(err)

	return mapAnyItem
}

func (it deserializeFromBytesTo) MapStringString(
	rawBytes []byte,
) (mappedItems map[string]string, err error) {
	err = Deserialize.UsingBytes(
		rawBytes,
		&mappedItems)

	return mappedItems, err
}

func (it deserializeFromBytesTo) MapStringStringMust(
	rawBytes []byte,
) (mappedItems map[string]string) {
	err := Deserialize.UsingBytes(
		rawBytes,
		&mappedItems)

	errcore.HandleErr(err)

	return mappedItems
}

func (it deserializeFromBytesTo) ResultCollection(
	rawBytes []byte,
) (*ResultsCollection, error) {
	empty := NewResultsCollection.
		Empty()
	err := Deserialize.UsingBytes(
		rawBytes,
		empty)

	if err == nil {
		return empty, nil
	}

	// has error
	return nil, err
}

func (it deserializeFromBytesTo) ResultCollectionMust(
	rawBytes []byte,
) *ResultsCollection {
	rs, err := it.ResultCollection(rawBytes)
	errcore.HandleErr(err)

	return rs
}

func (it deserializeFromBytesTo) ResultsPtrCollection(
	rawBytes []byte,
) (*ResultsPtrCollection, error) {
	empty := NewResultsPtrCollection.
		Empty()
	err := Deserialize.UsingBytes(
		rawBytes,
		empty)

	if err == nil {
		return empty, nil
	}

	// has error
	return nil, err
}

func (it deserializeFromBytesTo) ResultsPtrCollectionMust(
	rawBytes []byte,
) *ResultsPtrCollection {
	rs, err := it.ResultsPtrCollection(
		rawBytes)
	errcore.HandleErr(err)

	return rs
}

func (it deserializerLogic) Result(
	rawBytes []byte,
) (jsonResult Result, err error) {
	jsonResult = Empty.Result()
	err = Deserialize.UsingBytes(
		rawBytes, jsonResult)

	if err == nil {
		return jsonResult, jsonResult.MeaningfulError()
	}

	// has error
	return jsonResult, errcore.MergeErrors(
		err,
		jsonResult.MeaningfulError())
}

func (it deserializerLogic) ResultMust(
	rawBytes []byte,
) (jsonResult Result) {
	jsonResult, err := it.Result(rawBytes)
	errcore.MustBeEmpty(err)

	return jsonResult
}

func (it deserializerLogic) ResultPtr(
	rawBytes []byte,
) (jsonResult *Result, err error) {
	jsonResult = Empty.ResultPtr()
	err = Deserialize.UsingBytes(
		rawBytes, jsonResult)

	if err == nil {
		return jsonResult, jsonResult.MeaningfulError()
	}

	// has error
	return nil, errcore.MergeErrors(
		err,
		jsonResult.MeaningfulError())
}

func (it deserializerLogic) ResultPtrMust(
	rawBytes []byte,
) (jsonResult *Result) {
	jsonResult, err := it.ResultPtr(rawBytes)
	errcore.MustBeEmpty(err)

	return jsonResult
}

func (it deserializeFromBytesTo) MapResults(
	rawBytes []byte,
) (*MapResults, error) {
	empty := NewMapResults.
		Empty()
	err := Deserialize.UsingBytes(
		rawBytes,
		empty)

	if err == nil {
		return empty, nil
	}

	// has error
	return nil, err
}

func (it deserializeFromBytesTo) Bytes(
	rawBytes []byte,
) (nextDeserializedBytes []byte, err error) {
	err = Deserialize.UsingBytes(
		rawBytes,
		&nextDeserializedBytes)

	return nextDeserializedBytes, err
}

func (it deserializeFromBytesTo) BytesMust(
	rawBytes []byte,
) (nextDeserializedBytes []byte) {
	err := Deserialize.UsingBytes(
		rawBytes,
		&nextDeserializedBytes)
	errcore.HandleErr(err)

	return nextDeserializedBytes
}

func (it deserializeFromBytesTo) MapResultsMust(
	rawBytes []byte,
) *MapResults {
	rs, err := it.MapResults(rawBytes)
	errcore.HandleErr(err)

	return rs
}

func (it deserializeFromBytesTo) Bool(
	rawBytes []byte,
) (isResult bool, err error) {
	err = Deserialize.UsingBytes(rawBytes, &isResult)

	return isResult, err
}

func (it deserializeFromBytesTo) BoolMust(
	rawBytes []byte,
) (isResult bool) {
	err := Deserialize.UsingBytes(rawBytes, &isResult)
	errcore.HandleErr(err)

	return isResult
}
