package corejson

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/coredata"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
	"gitlab.com/evatix-go/core/msgtype"
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
	jsonBytes []byte, err error,
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
		Bytes: *jsonBytes,
		Error: nil,
	}
}

func NewFromAny(any interface{}) Result {
	jsonBytes, err := json.Marshal(any)

	if err != nil {
		return Result{
			Bytes: nil,
			Error: msgtype.MarshallingFailed.Error(
				err.Error(),
				reflectinternal.TypeName(any)),
		}
	}

	return Result{
		Bytes: jsonBytes,
		Error: err,
	}
}

func NewFromAnyPtr(any interface{}) *Result {
	jsonBytes, err := json.Marshal(any)

	if err != nil {
		return &Result{
			Bytes: nil,
			Error: msgtype.MarshallingFailed.Error(
				err.Error(),
				reflectinternal.TypeName(any)),
		}
	}

	return &Result{
		Bytes: jsonBytes,
		Error: err,
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

func NewResultsCollectionUsingJsonResults(
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

func EmptyMapResultsUsingCap() *MapResults {
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
