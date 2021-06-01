package corejson

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/coredata"
)

// NewUsingBytesError Get created with nil.
func NewUsingBytesError(bytesError *coredata.BytesError) Result {
	if bytesError == nil {
		return Result{
			Bytes: nil,
			Error: nil,
		}
	}

	return Result{
		Bytes: bytesError.Bytes,
		Error: bytesError.Error,
	}
}

func EmptyWithError(err error) Result {
	return Result{
		Bytes: nil,
		Error: err,
	}
}

func EmptyWithErrorPtr(err error) *Result {
	return &Result{
		Bytes: nil,
		Error: err,
	}
}

func EmptyWithoutErrorPtr() *Result {
	return EmptyWithErrorPtr(StaticJsonError)
}

func NewUsingBytes(
	jsonBytes *[]byte,
) Result {
	return Result{
		Bytes: jsonBytes,
		Error: nil,
	}
}

func NewUsingBytesPtr(
	jsonBytes *[]byte,
) *Result {
	return &Result{
		Bytes: jsonBytes,
		Error: nil,
	}
}

func NewPtr(
	jsonBytes []byte, err error,
) *Result {
	if err != nil {
		return EmptyWithErrorPtr(err)
	}

	if jsonBytes == nil {
		return EmptyWithoutErrorPtr()
	}

	return &Result{
		Bytes: &jsonBytes,
		Error: nil,
	}
}

func New(
	jsonBytes []byte, err error,
) Result {
	return *NewPtr(jsonBytes, err)
}

func NewPtrUsingBytesPtr(
	jsonBytes *[]byte, err error,
) *Result {
	if err != nil {
		return EmptyWithErrorPtr(err)
	}

	if jsonBytes == nil {
		return EmptyWithoutErrorPtr()
	}

	return &Result{
		Bytes: jsonBytes,
		Error: nil,
	}
}

func NewFromAny(any interface{}) *Result {
	jsonBytes, err := json.Marshal(any)

	if err != nil {
		return &Result{
			Bytes: &[]byte{},
			Error: err,
		}
	}

	return &Result{
		Bytes: &jsonBytes,
		Error: err,
	}
}

func EmptyResultsCollection() *ResultsCollection {
	list := make([]*Result, 0, 0)

	return &ResultsCollection{
		Items: &list,
	}
}

func NewResultsCollection(cap int) *ResultsCollection {
	list := make([]*Result, 0, cap)

	return &ResultsCollection{
		Items: &list,
	}
}

func NewResultsCollectionUsingJsoners(
	addCapacity int,
	jsoners ...Jsoner,
) *ResultsCollection {
	length := addCapacity
	if jsoners == nil {
		return NewResultsCollection(length)
	}

	additionalCapacity := len(jsoners)
	length += additionalCapacity
	list := make([]*Result, 0, length)
	resultsCollection := &ResultsCollection{
		Items: &list,
	}

	return resultsCollection.
		AddJsonerPtr(&jsoners)
}

func NewResultsCollectionUsingJsonResults(
	addCapacity int,
	results ...*Result,
) *ResultsCollection {
	length := addCapacity
	if results == nil {
		return NewResultsCollection(length)
	}

	additionalCapacity := len(results)
	length += additionalCapacity
	list := make([]*Result, 0, length)
	resultsCollection := &ResultsCollection{
		Items: &list,
	}

	return resultsCollection.
		AddNonNilItemsPtr(&results)
}

func NewResultsCollectionUsingAnys(
	addCapacity int,
	anys ...interface{},
) *ResultsCollection {
	length := addCapacity
	if anys == nil {
		return NewResultsCollection(length)
	}

	additionalCapacity := len(anys)
	length += additionalCapacity
	list := make([]*Result, 0, length)
	resultsCollection := &ResultsCollection{
		Items: &list,
	}

	return resultsCollection.
		AddsAnysPtr(&anys)
}
