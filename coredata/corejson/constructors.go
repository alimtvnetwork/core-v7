package corejson

import "encoding/json"

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

func NewUsingBytes(jsonBytes *[]byte) Result {
	return Result{
		Bytes: jsonBytes,
		Error: nil,
	}
}

func NewUsingBytesPtr(jsonBytes *[]byte) *Result {
	return &Result{
		Bytes: jsonBytes,
		Error: nil,
	}
}

func NewPtr(jsonBytes []byte, err error) *Result {
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

func New(jsonBytes []byte, err error) Result {
	return *NewPtr(jsonBytes, err)
}

func NewPtrUsingBytesPtr(jsonBytes *[]byte, err error) *Result {
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
