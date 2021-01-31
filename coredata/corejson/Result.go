package corejson

import (
	"gitlab.com/evatix-go/core/coreindexes"
)

type Result struct {
	jsonString *string
	Bytes      *[]byte
	Error      error
}

func EmptyJsonResult(err error) Result {
	return Result{
		Bytes: nil,
		Error: err,
	}
}

func EmptyJsonResultPtr(err error) *Result {
	return &Result{
		Bytes: nil,
		Error: err,
	}
}

func EmptyJsonResultWithoutErrorPtr() *Result {
	return EmptyJsonResultPtr(StaticJsonError)
}

func NewJsonResultBytes(jsonBytes *[]byte) Result {
	return Result{
		Bytes: jsonBytes,
		Error: nil,
	}
}

func NewJsonResultBytesPtr(jsonBytes *[]byte) *Result {
	return &Result{
		Bytes: jsonBytes,
		Error: nil,
	}
}

func NewJsonResultPtr(jsonBytes []byte, err error) *Result {
	if err != nil {
		return EmptyJsonResultPtr(err)
	}

	if jsonBytes == nil {
		return EmptyJsonResultWithoutErrorPtr()
	}

	return &Result{
		Bytes: &jsonBytes,
		Error: nil,
	}
}

func NewJsonResult(jsonBytes []byte, err error) Result {
	return *NewJsonResultPtr(jsonBytes, err)
}

func NewJsonResultPtrUsingBytesPtr(jsonBytes *[]byte, err error) *Result {
	if err != nil {
		return EmptyJsonResultPtr(err)
	}

	if jsonBytes == nil {
		return EmptyJsonResultWithoutErrorPtr()
	}

	return &Result{
		Bytes: jsonBytes,
		Error: nil,
	}
}

func (jsonResult *Result) JsonString() string {
	return *jsonResult.JsonStringPtr()
}

func (jsonResult *Result) JsonStringPtr() *string {
	if jsonResult.jsonString != nil {
		return jsonResult.jsonString
	}

	if jsonResult.jsonString == nil && jsonResult.HasBytes() {
		jsonString := string(*jsonResult.Bytes)
		jsonResult.jsonString = &jsonString
	} else if jsonResult.jsonString == nil {
		emptyStr := ""
		jsonResult.jsonString = &emptyStr
	}

	return jsonResult.jsonString
}

func (jsonResult *Result) HasError() bool {
	return jsonResult.Error != nil
}

func (jsonResult *Result) IsEmptyError() bool {
	return jsonResult.Error == nil
}

func (jsonResult *Result) HandleError() {
	if jsonResult.IsEmptyError() {
		return
	}

	panic(jsonResult.Error)
}

func (jsonResult *Result) HandleErrorWithMsg(msg string) {
	if jsonResult.IsEmptyError() {
		return
	}

	if msg != "" {
		panic(msg + jsonResult.Error.Error())
	}

	panic(jsonResult.Error)
}

func (jsonResult *Result) HasBytes() bool {
	return !jsonResult.IsEmptyJsonBytes()
}

// len == 0, nil, {} returns as empty true
func (jsonResult *Result) IsEmptyJsonBytes() bool {
	isEmptyFirst := jsonResult.HasError() ||
		jsonResult.Bytes == nil

	if isEmptyFirst {
		return isEmptyFirst
	}

	length := len(*jsonResult.Bytes)

	if length == 0 {
		return true
	}

	if length == 2 {
		// empty json
		return (*jsonResult.Bytes)[coreindexes.First] == 123 &&
			(*jsonResult.Bytes)[coreindexes.Second] == 125
	}

	return false
}

func (jsonResult *Result) IsEmptyJson() bool {
	return jsonResult.Bytes == nil || len(*jsonResult.Bytes) == 0
}

func (jsonResult *Result) HasJson() bool {
	return jsonResult.HasBytes()
}
