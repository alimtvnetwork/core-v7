package corejson

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
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
	return EmptyWithErrorPtr(nil)
}

func NewUsingBytes(
	jsonBytes []byte,
) Result {
	return Result{
		Bytes: jsonBytes,
		Error: nil,
	}
}

func NewUsingBytesPtr(
	jsonBytes *[]byte,
) *Result {
	if jsonBytes == nil {
		return &Result{
			Bytes: nil,
			Error: nil,
		}
	}

	return &Result{
		Bytes: *jsonBytes,
		Error: nil,
	}
}

func NewPtr(
	jsonBytes []byte,
	err error,
	typeName string,
) *Result {
	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func New(
	jsonBytes []byte,
	err error,
	typeName string,
) Result {
	return Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func NewPtrUsingBytesPtr(
	jsonBytes *[]byte,
	err error,
	typeName string,
) *Result {
	if err != nil {
		return &Result{
			Bytes:    []byte{},
			Error:    err,
			TypeName: typeName,
		}
	}

	if jsonBytes == nil {
		return &Result{
			Bytes:    []byte{},
			Error:    nil,
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    *jsonBytes,
		Error:    nil,
		TypeName: typeName,
	}
}

func NewFromAny(any interface{}) Result {
	jsonBytes, err := json.Marshal(any)
	typeName := reflectinternal.TypeName(any)

	if err != nil {
		return Result{
			Bytes: jsonBytes,
			Error: errcore.MarshallingFailedType.Error(
				err.Error(),
				typeName),
			TypeName: typeName,
		}
	}

	return Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func NewFromAnyPtr(any interface{}) *Result {
	jsonBytes, err := json.Marshal(any)
	typeName := reflectinternal.TypeName(any)

	if err != nil {
		return &Result{
			Bytes: jsonBytes,
			Error: errcore.MarshallingFailedType.Error(
				err.Error(),
				typeName),
			TypeName: typeName,
		}
	}

	return &Result{
		Bytes:    jsonBytes,
		Error:    err,
		TypeName: typeName,
	}
}

func EmptyResultsCollection() *ResultsCollection {
	list := make([]Result, 0)

	return &ResultsCollection{
		Items: list,
	}
}

func NewResultsCollection(cap int) *ResultsCollection {
	list := make([]Result, 0, cap)

	return &ResultsCollection{
		Items: list,
	}
}

func EmptyResultsPtrCollection() *ResultsPtrCollection {
	list := make([]*Result, 0)

	return &ResultsPtrCollection{
		Items: list,
	}
}

func NewResultsPtrCollection(cap int) *ResultsPtrCollection {
	list := make([]*Result, 0, cap)

	return &ResultsPtrCollection{
		Items: list,
	}
}

func NewBytesCollection(cap int) *BytesCollection {
	list := make([][]byte, 0, cap)

	return &BytesCollection{
		Items: list,
	}
}

func NewBytesCollectionUsing(anyItems ...interface{}) (*BytesCollection, error) {
	length := len(anyItems)
	collection := NewBytesCollection(length)
	err := collection.AddAnyItems(anyItems...)

	return collection, err

}

func EmptyBytesCollection() *BytesCollection {
	return NewBytesCollection(constants.Zero)
}

func NewResultsCollectionUsingJsoners(
	isIgnoreNilOrError bool,
	addCapacity int,
	jsoners ...Jsoner,
) *ResultsCollection {
	length := addCapacity
	if jsoners == nil {
		return NewResultsCollection(length)
	}

	actualLength := len(jsoners)
	length += actualLength
	list := NewResultsCollection(length)

	return list.
		AddJsoners(
			isIgnoreNilOrError,
			jsoners...)
}

func NewResultsCollectionPtrUsingJsonResultsPtr(
	addCapacity int,
	results ...*Result,
) *ResultsPtrCollection {
	length := addCapacity
	if results == nil {
		return NewResultsPtrCollection(length)
	}

	actualLength := len(results)
	length += actualLength
	list := NewResultsPtrCollection(length)

	if actualLength == 0 {
		return list
	}

	return list.
		AddNonNilItemsPtr(results...)
}

func NewResultsCollectionUsingJsonResultsPtr(
	addCapacity int,
	results ...*Result,
) *ResultsCollection {
	length := addCapacity
	if results == nil {
		return NewResultsCollection(length)
	}

	actualLength := len(results)
	length += actualLength
	list := NewResultsCollection(length)

	if actualLength == 0 {
		return list
	}

	return list.
		AddNonNilItemsPtr(results...)
}

func NewResultsCollectionUsingJsonResults(
	addCapacity int,
	results ...Result,
) *ResultsCollection {
	length := addCapacity
	if results == nil {
		return NewResultsCollection(length)
	}

	actualLength := len(results)
	length += actualLength
	list := NewResultsCollection(length)

	if actualLength == 0 {
		return list
	}

	return list.
		Adds(results...)
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
	list := NewResultsCollection(length)

	return list.
		AddsAnys(&anys)
}

func EmptyMapResults() *MapResults {
	return &MapResults{
		Items: map[string]Result{},
	}
}

func NewMapResultsUsingCap(
	addCapacity int,
) *MapResults {
	return &MapResults{
		Items: make(map[string]Result, addCapacity),
	}
}

func NewMapResultsUsingKeyAnys(
	addCapacity int,
	keyAnys ...KeyAny,
) *MapResults {
	length := addCapacity
	if keyAnys == nil {
		return NewMapResultsUsingCap(length)
	}

	additionalCapacity := len(keyAnys)
	length += additionalCapacity
	mapResults := NewMapResultsUsingCap(length)

	return mapResults.
		AddKeyAnys(keyAnys...)
}

func NewMapResultsUsingMap(
	isClone, isDeepClone bool,
	addCapacity int,
	mapResults map[string]Result,
) *MapResults {
	if len(mapResults) == 0 {
		return NewMapResultsUsingCap(
			addCapacity)
	}

	additionalCapacity :=
		len(mapResults) +
			addCapacity

	hasNoChange := additionalCapacity == len(mapResults) &&
		!isClone &&
		!isDeepClone

	if hasNoChange {
		return &MapResults{
			Items: mapResults,
		}
	}

	finalMapResults := NewMapResultsUsingCap(
		additionalCapacity)

	return finalMapResults.AddMapResultsUsingCloneOption(
		isClone,
		isDeepClone,
		mapResults)
}

func NewMapResultsUsingKeyResults(
	addCapacity int,
	keyWithResults ...KeyWithResult,
) *MapResults {
	length := addCapacity
	if keyWithResults == nil {
		return NewMapResultsUsingCap(length)
	}

	additionalCapacity := len(keyWithResults)
	length += additionalCapacity
	mapResults := NewMapResultsUsingCap(length)

	return mapResults.
		AddKeysWithResults(keyWithResults...)
}

func NewMapResultsUsingKeyJsoners(
	addCapacity int,
	keyWithJsoners ...KeyWithJsoner,
) *MapResults {
	length := addCapacity
	if keyWithJsoners == nil {
		return NewMapResultsUsingCap(length)
	}

	additionalCapacity := len(keyWithJsoners)
	length += additionalCapacity
	mapResults := NewMapResultsUsingCap(length)

	return mapResults.
		AddKeysWithJsoners(keyWithJsoners...)
}
