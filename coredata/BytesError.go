package coredata

import (
	"gitlab.com/evatix-go/core/coreindexes"
)

type BytesError struct {
	toString *string
	Bytes    *[]byte
	Error    error
}

func (bytesError *BytesError) String() string {
	return *bytesError.StringPtr()
}

func (bytesError *BytesError) StringPtr() *string {
	if bytesError.toString != nil {
		return bytesError.toString
	}

	if bytesError.toString == nil && bytesError.HasBytes() {
		jsonString := string(*bytesError.Bytes)
		bytesError.toString = &jsonString
	} else if bytesError.toString == nil {
		emptyStr := ""
		bytesError.toString = &emptyStr
	}

	return bytesError.toString
}

func (bytesError *BytesError) HasError() bool {
	return bytesError.Error != nil
}

func (bytesError *BytesError) IsEmptyError() bool {
	return bytesError.Error == nil
}

func (bytesError *BytesError) HandleError() {
	if bytesError.IsEmptyError() {
		return
	}

	panic(bytesError.Error)
}

func (bytesError *BytesError) HandleErrorWithMsg(msg string) {
	if bytesError.IsEmptyError() {
		return
	}

	if msg != "" {
		panic(msg + bytesError.Error.Error())
	}

	panic(bytesError.Error)
}

func (bytesError *BytesError) HasBytes() bool {
	return !bytesError.IsEmptyOrErrorBytes()
}

// IsEmptyOrErrorBytes len == 0, nil, {} returns as empty true
func (bytesError *BytesError) IsEmptyOrErrorBytes() bool {
	isEmptyFirst := bytesError.HasError() ||
		bytesError.Bytes == nil

	if isEmptyFirst {
		return isEmptyFirst
	}

	length := len(*bytesError.Bytes)

	if length == 0 {
		return true
	}

	if length == 2 {
		// empty json
		return (*bytesError.Bytes)[coreindexes.First] == 123 &&
			(*bytesError.Bytes)[coreindexes.Second] == 125
	}

	return false
}

func (bytesError *BytesError) IsEmpty() bool {
	return bytesError.Bytes == nil || len(*bytesError.Bytes) == 0
}
