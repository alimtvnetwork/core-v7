package corestr

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
)

type JsonResult struct {
	jsonString *string
	Bytes      *[]byte
	Error      error
}

func (jsonResult *JsonResult) JsonString() string {
	return *jsonResult.JsonStringPtr()
}

func (jsonResult *JsonResult) JsonStringPtr() *string {
	if jsonResult.jsonString == nil {
		if jsonResult.HasBytes() {
			jsonString := string(*jsonResult.Bytes)
			jsonResult.jsonString = &jsonString
		} else {
			jsonResult.jsonString = constants.EmptyStringPtr
		}
	}

	return jsonResult.jsonString
}

func EmptyJsonResult(err error) JsonResult {
	return JsonResult{
		Bytes: nil,
		Error: err,
	}
}

func EmptyJsonResultPtr(err error) *JsonResult {
	return &JsonResult{
		Bytes: nil,
		Error: err,
	}
}

func EmptyJsonResultWithoutErrorPtr() *JsonResult {
	return EmptyJsonResultPtr(StaticJsonError)
}

func NewJsonResultBytes(jsonBytes *[]byte) JsonResult {
	return JsonResult{
		Bytes: jsonBytes,
		Error: nil,
	}
}

func NewJsonResultBytesPtr(jsonBytes *[]byte) *JsonResult {
	return &JsonResult{
		Bytes: jsonBytes,
		Error: nil,
	}
}

func NewJsonResultPtr(jsonBytes []byte, err error) *JsonResult {
	if err != nil {
		return EmptyJsonResultPtr(err)
	}

	if jsonBytes == nil {
		return EmptyJsonResultWithoutErrorPtr()
	}

	return &JsonResult{
		Bytes: &jsonBytes,
		Error: nil,
	}
}

func NewJsonResult(jsonBytes []byte, err error) JsonResult {
	return *NewJsonResultPtr(jsonBytes, err)
}

func (jsonResult *JsonResult) HasError() bool {
	return jsonResult.Error != nil
}

func (jsonResult *JsonResult) IsEmptyError() bool {
	return jsonResult.Error == nil
}

func (jsonResult *JsonResult) HandleError() {
	if jsonResult.IsEmptyError() {
		return
	}

	panic(jsonResult.Error)
}

func (jsonResult *JsonResult) HandleErrorWithMsg(msg string) {
	if jsonResult.IsEmptyError() {
		return
	}

	if msg != "" {
		panic(msg + jsonResult.Error.Error())
	}

	panic(jsonResult.Error)
}

func (jsonResult *JsonResult) HasBytes() bool {
	return !jsonResult.IsBytesEmpty()
}

func (jsonResult *JsonResult) IsBytesEmpty() bool {
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
		return (*jsonResult.Bytes)[coreindexes.First] == 123 &&
			(*jsonResult.Bytes)[coreindexes.Second] == 125
	}

	return false
}

func (jsonResult *JsonResult) IsEmptyJson() bool {
	return jsonResult.Bytes == nil || len(*jsonResult.Bytes) == 0
}

func (jsonResult *JsonResult) HasJson() bool {
	return jsonResult.HasBytes()
}
