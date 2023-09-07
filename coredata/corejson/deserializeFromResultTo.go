package corejson

import "gitlab.com/auk-go/core/errcore"

type deserializeFromResultTo struct{}

func (it deserializeFromResultTo) String(
	jsonResult *Result,
) (line string, err error) {
	err = Deserialize.Apply(jsonResult, &line)

	return line, err
}

func (it deserializeFromResultTo) Bool(
	jsonResult *Result,
) (isResult bool, err error) {
	err = Deserialize.Apply(jsonResult, &isResult)

	return isResult, err
}

func (it deserializeFromResultTo) Byte(
	jsonResult *Result,
) (byteVal byte, err error) {
	err = Deserialize.Apply(jsonResult, &byteVal)

	return byteVal, err
}

func (it deserializeFromResultTo) ByteMust(
	jsonResult *Result,
) byte {
	result, err := it.Byte(jsonResult)

	if err != nil {
		panic(err)
	}

	return result
}

func (it deserializeFromResultTo) BoolMust(
	jsonResult *Result,
) bool {
	result, err := it.Bool(jsonResult)

	if err != nil {
		panic(err)
	}

	return result
}

func (it deserializeFromResultTo) StringMust(
	jsonResult *Result,
) string {
	result, err := it.String(jsonResult)

	if err != nil {
		panic(err)
	}

	return result
}

func (it deserializeFromResultTo) StringsMust(
	jsonResult *Result,
) (lines []string) {
	err := jsonResult.Deserialize(&lines)

	if err != nil {
		panic(err)
	}

	return lines
}

func (it deserializeFromResultTo) MapAnyItem(
	jsonResult *Result,
) (mapAnyItem map[string]interface{}, err error) {
	err = jsonResult.Deserialize(
		&mapAnyItem)

	return mapAnyItem, err
}

func (it deserializeFromResultTo) MapAnyItemMust(
	jsonResult *Result,
) (mapAnyItem map[string]interface{}) {
	err := jsonResult.Deserialize(
		&mapAnyItem)

	errcore.HandleErr(err)

	return mapAnyItem
}

func (it deserializeFromResultTo) MapStringString(
	jsonResult *Result,
) (mappedItems map[string]string, err error) {
	err = jsonResult.Deserialize(
		&mappedItems)

	return mappedItems, err
}

func (it deserializeFromResultTo) MapStringStringMust(
	jsonResult *Result,
) (mappedItems map[string]string) {
	err := jsonResult.Deserialize(
		&mappedItems)
	errcore.HandleErr(err)

	return mappedItems
}

func (it deserializeFromResultTo) ResultCollection(
	jsonResult *Result,
) (*ResultsCollection, error) {
	empty := NewResultsCollection.
		Empty()
	err := jsonResult.
		Deserialize(empty)

	if err == nil {
		return empty, nil
	}

	// has error
	return nil, err
}

func (it deserializeFromResultTo) ResultCollectionMust(
	jsonResult *Result,
) *ResultsCollection {
	rs, err := it.ResultCollection(jsonResult)
	errcore.HandleErr(err)

	return rs
}

func (it deserializeFromResultTo) ResultsPtrCollection(
	jsonResult *Result,
) (*ResultsPtrCollection, error) {
	empty := NewResultsPtrCollection.
		Empty()
	err := jsonResult.
		Deserialize(empty)

	if err == nil {
		return empty, nil
	}

	// has error
	return nil, err
}

func (it deserializeFromResultTo) ResultsPtrCollectionMust(
	jsonResult *Result,
) *ResultsPtrCollection {
	rs, err := it.ResultsPtrCollection(
		jsonResult)
	errcore.HandleErr(err)

	return rs
}

func (it deserializeFromResultTo) Result(
	jsonResultInput *Result,
) (jsonResult Result, err error) {
	empty := Empty.ResultPtr()
	err = jsonResultInput.
		Deserialize(empty)

	if err == nil {
		return empty.NonPtr(), nil
	}

	// has error
	return jsonResult, errcore.MergeErrors(
		err,
		jsonResult.MeaningfulError())
}

func (it deserializeFromResultTo) ResultMust(
	jsonResultInput *Result,
) (jsonResult Result) {
	jsonResult, err := it.Result(jsonResultInput)
	errcore.MustBeEmpty(err)

	return jsonResult
}

func (it deserializeFromResultTo) ResultPtr(
	jsonResultInput *Result,
) (jsonResultPtr *Result, err error) {
	jsonResult, err := it.Result(jsonResultInput)

	return jsonResult.Ptr(), err
}

func (it deserializeFromResultTo) ResultPtrMust(
	jsonResultInput *Result,
) (jsonResultPtr *Result) {
	jsonResult, err := it.Result(jsonResultInput)
	errcore.HandleErr(err)

	return jsonResult.Ptr()
}

func (it deserializeFromResultTo) MapResults(
	jsonResult *Result,
) (*MapResults, error) {
	empty := NewMapResults.
		Empty()
	err := jsonResult.
		Deserialize(empty)

	if err == nil {
		return empty, nil
	}

	// has error
	return nil, err
}

func (it deserializeFromResultTo) Bytes(
	jsonResult *Result,
) (nextDeserializedBytes []byte, err error) {
	jsonResultOut, err := it.Result(jsonResult)

	if err == nil {
		return jsonResultOut.Bytes, jsonResultOut.MeaningfulError()
	}

	return jsonResultOut.Bytes, err
}

func (it deserializeFromResultTo) BytesMust(
	jsonResult *Result,
) (nextDeserializedBytes []byte) {
	nextDeserializedBytes, err := it.Bytes(jsonResult)
	errcore.HandleErr(err)

	return nextDeserializedBytes
}

func (it deserializeFromResultTo) MapResultsMust(
	jsonResult *Result,
) *MapResults {
	rs, err := it.MapResults(jsonResult)
	errcore.HandleErr(err)

	return rs
}
