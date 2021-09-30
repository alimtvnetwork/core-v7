package corejson

import (
	"bytes"
	"encoding/json"
	"errors"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/msgtype"
)

type Result struct {
	jsonString *string
	Bytes      []byte
	Error      error
}

func (it Result) JsonString() string {
	return *it.JsonStringPtr()
}

func (it *Result) JsonStringPtr() *string {
	if it == nil {
		return constants.EmptyStringPtr
	}

	if it.jsonString != nil {
		return it.jsonString
	}

	if it.jsonString == nil && it.HasBytes() {
		jsonString := string(it.Bytes)
		it.jsonString = &jsonString
	} else if it.jsonString == nil {
		emptyStr := ""
		it.jsonString = &emptyStr
	}

	return it.jsonString
}

func (it *Result) Length() int {
	if it == nil || it.Bytes == nil {
		return 0
	}

	return len(it.Bytes)
}

func (it *Result) HasError() bool {
	return it != nil && it.Error != nil
}

func (it *Result) ErrorString() string {
	if it.IsEmptyError() {
		return constants.EmptyString
	}

	return it.Error.Error()
}

func (it *Result) IsErrorEqual(err error) bool {
	if it.IsEmptyError() && err == nil {
		return true
	}

	if it.IsEmptyError() || err == nil {
		return false
	}

	if it.HasError() && it.ErrorString() == err.Error() {
		return true
	}

	return false
}

func (it *Result) String() string {
	if it.HasIssuesOrEmpty() {
		return constants.EmptyString
	}

	return it.JsonString()
}

func (it *Result) ValuesNonPtr() []byte {
	return it.ValueMust()
}

func (it *Result) ValueMust() []byte {
	if it.HasIssuesOrEmpty() {
		return []byte{}
	}

	return it.Bytes
}

// MeaningfulError create error even if results are nil.
func (it *Result) MeaningfulError() error {
	var msgVariation msgtype.Variation

	if it.IsEmptyJsonBytes() {
		msgVariation = msgtype.JsonResultBytesAreNilOrEmpty
	}

	if it.HasError() {
		return msgtype.FailedToParse.Error(
			it.Error.Error(),
			msgVariation.String())
	}

	return nil
}

func (it *Result) IsEmptyError() bool {
	return it == nil || it.Error == nil
}

func (it *Result) HasSafeItems() bool {
	return !it.IsEmptyJsonBytes()
}

func (it *Result) HasIssuesOrEmpty() bool {
	return it.IsEmptyJsonBytes()
}

func (it *Result) HandleError() {
	if it == nil || it.IsEmptyError() {
		return
	}

	panic(it.MeaningfulError())
}

func (it *Result) HandleErrorWithMsg(msg string) {
	if it.IsEmptyError() {
		return
	}

	err := it.MeaningfulError()

	if err != nil && msg != "" {
		panic(msg + err.Error())
	}

	panic(err)
}

func (it *Result) HasBytes() bool {
	return !it.IsEmptyJsonBytes()
}

// IsEmptyJsonBytes len == 0, nil, {} returns as empty true
func (it *Result) IsEmptyJsonBytes() bool {
	if it == nil {
		return true
	}

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
		return (it.Bytes)[coreindexes.First] == constants.CurlyBraceStartChar &&
			(it.Bytes)[coreindexes.Second] == constants.CurlyBraceEndChar
	}

	return false
}

func (it *Result) IsEmptyJson() bool {
	return it == nil || len(it.Bytes) == 0
}

func (it *Result) HasJson() bool {
	return it.HasBytes()
}

func (it *Result) InjectInto(
	injector JsonParseSelfInjector,
) error {
	return injector.JsonParseSelfInject(it)
}

func (it *Result) Unmarshal(any interface{}) error {
	if it == nil {
		return defaulterr.UnMarshallingFailedDueToNilOrEmpty
	}

	if it.HasError() {
		return it.Error
	}

	return json.Unmarshal(it.Bytes, any)
}

func (it *Result) UnmarshalResult() (*Result, error) {
	empty := EmptyWithoutErrorPtr()
	err := it.Unmarshal(empty)

	return empty, err
}

//goland:noinspection GoLinterLocal
func (it *Result) JsonModel() *ResultModel {
	return NewModel(it)
}

//goland:noinspection GoLinterLocal
func (it *Result) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *Result) UnmarshalJSON(data []byte) error {
	var dataModel ResultModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		transpileModelToResult(&dataModel, it)
	}

	return err
}

func (it Result) Json() Result {
	return NewFromAny(it)
}

func (it Result) JsonPtr() *Result {
	return NewFromAnyPtr(it)
}

// ParseInjectUsingJson It will not update the self but creates a new one.
func (it *Result) ParseInjectUsingJson(
	jsonResultIn *Result,
) (*Result, error) {
	if jsonResultIn == nil || jsonResultIn.IsEmptyJsonBytes() {
		return EmptyWithoutErrorPtr(), defaulterr.UnMarshallingFailedDueToNilOrEmpty
	}

	err := json.Unmarshal(
		jsonResultIn.Bytes,
		&it)

	if err != nil {
		return EmptyWithErrorPtr(err), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *Result) ParseInjectUsingJsonMust(
	jsonResultIn *Result,
) *Result {
	result, err := it.ParseInjectUsingJson(
		jsonResultIn)

	if err != nil {
		panic(err)
	}

	return result
}

func (it *Result) AsJsoner() Jsoner {
	return it
}

func (it *Result) JsonParseSelfInject(
	jsonResultIn *Result,
) error {
	_, err := it.ParseInjectUsingJson(jsonResultIn)

	return err
}

func (it *Result) AsJsonParseSelfInjector() JsonParseSelfInjector {
	return it
}

func (it *Result) AsJsonMarshaller() JsonMarshaller {
	return it
}

func (it *Result) CloneError() error {
	if it.HasError() {
		return errors.New(it.Error.Error())
	}

	return nil
}

func (it *Result) IsEqual(another *Result) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it.Length() != another.Length() {
		return false
	}

	if !it.IsErrorEqual(another.Error) {
		return false
	}

	if it.jsonString != nil && another.jsonString != nil &&
		it.jsonString == another.jsonString {
		return true
	}

	return bytes.Equal(it.Bytes, another.Bytes)
}

func (it *Result) Clone(isDeepClone bool) *Result {
	if it == nil {
		return nil
	}

	if it.Length() == 0 {
		return NewPtr([]byte{}, it.CloneError())
	}

	if !isDeepClone || it.Length() == 0 {
		return NewPtr(it.Bytes, it.CloneError())
	}

	newBytes := make([]byte, it.Length())
	copy(newBytes, it.Bytes)

	return NewPtr(newBytes, it.CloneError())
}
