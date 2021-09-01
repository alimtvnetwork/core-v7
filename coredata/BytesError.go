package coredata

import (
	"gitlab.com/evatix-go/core/coreindexes"
)

type BytesError struct {
	toString *string
	Bytes    []byte
	Error    error
}

func (it *BytesError) String() string {
	return *it.StringPtr()
}

func (it *BytesError) StringPtr() *string {
	if it.toString != nil {
		return it.toString
	}

	if it.toString == nil && it.HasBytes() {
		jsonString := string(it.Bytes)
		it.toString = &jsonString
	} else if it.toString == nil {
		emptyStr := ""
		it.toString = &emptyStr
	}

	return it.toString
}

func (it *BytesError) HasError() bool {
	return it != nil && it.Error != nil
}

func (it *BytesError) IsEmptyError() bool {
	return it == nil || it.Error == nil
}

func (it *BytesError) HandleError() {
	if it.IsEmptyError() {
		return
	}

	panic(it.Error)
}

func (it *BytesError) HandleErrorWithMsg(msg string) {
	if it.IsEmptyError() {
		return
	}

	if msg != "" {
		panic(msg + it.Error.Error())
	}

	panic(it.Error)
}

func (it *BytesError) HasBytes() bool {
	return !it.IsEmptyOrErrorBytes()
}

// IsEmptyOrErrorBytes len == 0, nil, {} returns as empty true
func (it *BytesError) IsEmptyOrErrorBytes() bool {
	isEmptyFirst := it.HasError() ||
		it.Bytes == nil

	if isEmptyFirst {
		return isEmptyFirst
	}

	length := len(it.Bytes)

	if length == 0 {
		return true
	}

	if length == 2 {
		// empty json
		return (it.Bytes)[coreindexes.First] == 123 &&
			(it.Bytes)[coreindexes.Second] == 125
	}

	return false
}

func (it *BytesError) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Bytes)
}

func (it *BytesError) IsEmpty() bool {
	return it == nil || len(it.Bytes) == 0
}
