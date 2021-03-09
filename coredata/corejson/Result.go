package corejson

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/coreindexes"
)

type Result struct {
	jsonString *string
	Bytes      *[]byte
	Error      error
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

func (jsonResult *Result) InjectInto(
	injector ParseSelfInjector,
) error {
	return injector.JsonParseSelfInject(jsonResult)
}

func (jsonResult *Result) Unmarshal(any interface{}) error {
	if jsonResult.HasError() {
		return jsonResult.Error
	}

	return json.Unmarshal(*jsonResult.Bytes, any)
}

//goland:noinspection GoLinterLocal
func (jsonResult *Result) JsonModel() *ResultModel {
	return NewModel(jsonResult)
}

//goland:noinspection GoLinterLocal
func (jsonResult *Result) JsonModelAny() interface{} {
	return jsonResult.JsonModel()
}

func (jsonResult *Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonResult.JsonModel())
}

func (jsonResult *Result) UnmarshalJSON(data []byte) error {
	var dataModel ResultModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		transpileModelToResult(&dataModel, jsonResult)
	}

	return err
}

//goland:noinspection GoLinterLocal
func (jsonResult *Result) Json() *Result {
	return jsonResult
}

// It will not update the self but creates a new one.
func (jsonResult *Result) ParseInjectUsingJson(
	jsonResultIn *Result,
) (*Result, error) {
	if jsonResultIn == nil || jsonResultIn.IsEmptyJsonBytes() {
		return EmptyWithoutErrorPtr(), nil
	}

	err := json.Unmarshal(
		*jsonResultIn.Bytes,
		&jsonResult)

	if err != nil {
		return EmptyWithErrorPtr(err), err
	}

	return jsonResult, nil
}

// Panic if error
func (jsonResult *Result) ParseInjectUsingJsonMust(
	jsonResultIn *Result,
) *Result {
	result, err := jsonResult.ParseInjectUsingJson(
		jsonResultIn)

	if err != nil {
		panic(err)
	}

	return result
}

func (jsonResult *Result) AsJsoner() *Jsoner {
	var jsoner Jsoner = jsonResult

	return &jsoner
}

func (jsonResult *Result) JsonParseSelfInject(
	jsonResultIn *Result,
) error {
	_, err := jsonResult.ParseInjectUsingJson(jsonResultIn)

	return err
}

func (jsonResult *Result) AsJsonParseSelfInjector() *ParseSelfInjector {
	var jsonMarshaller ParseSelfInjector = jsonResult

	return &jsonMarshaller
}

func (jsonResult *Result) AsJsonMarshaller() *JsonMarshaller {
	var jsonMarshaller JsonMarshaller = jsonResult

	return &jsonMarshaller
}
