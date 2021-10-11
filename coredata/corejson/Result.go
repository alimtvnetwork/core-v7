package corejson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/csvinternal"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

type Result struct {
	jsonString *string
	Bytes      []byte
	Error      error
	TypeName   string
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
	if it.Error == nil && err == nil {
		return true
	}

	if it.Error == nil || err == nil {
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
	var msgVariation errcore.Variation

	if it.IsEmptyJsonBytes() {
		msgVariation = errcore.JsonResultBytesAreNilOrEmpty
	}

	if it.HasError() {
		return errcore.FailedToParse.Error(
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
		return errcore.
			UnMarshallingFailed.
			Error(
				"cannot unmarshal if JsonResult is ni, type",
				reflectinternal.TypeName(any))
	}

	if it.HasError() {
		reference := errcore.Var3NoType(
			"JsonResult Error", it.Error,
			"Source Type", it.TypeName,
			"To Reference Type", reflectinternal.TypeName(any))

		return errcore.
			UnMarshallingFailed.
			Error(
				"cannot unmarshal if JsonResult has already error.",
				reference)
	}

	err := json.Unmarshal(it.Bytes, any)

	if err == nil {
		return nil
	}

	reference := errcore.Var3NoType(
		"Unmarshall Error", err.Error(),
		"Source Type", it.TypeName,
		"To Reference Type", reflectinternal.TypeName(any))

	return errcore.
		UnMarshallingFailed.
		ErrorRefOnly(reference)
}

func (it *Result) UnmarshalIgnoreExistingError(any interface{}) error {
	if it == nil {
		return errcore.
			UnMarshallingFailed.
			Error(
				"cannot unmarshal if JsonResult is nil, type",
				reflectinternal.TypeName(any))
	}

	err := json.Unmarshal(it.Bytes, any)

	if err == nil {
		return nil
	}

	reference := errcore.Var3NoType(
		"Unmarshall Error", err.Error(),
		"Source Type", it.TypeName,
		"To Reference Type", reflectinternal.TypeName(any))

	return errcore.
		UnMarshallingFailed.
		ErrorRefOnly(reference)
}

func (it *Result) UnmarshalResult() (*Result, error) {
	empty := EmptyWithoutErrorPtr()
	err := it.Unmarshal(empty)

	return empty, err
}

//goland:noinspection GoLinterLocal
func (it Result) JsonModel() Result {
	return it
}

//goland:noinspection GoLinterLocal
func (it *Result) JsonModelAny() interface{} {
	return it.JsonModel()
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
	err := jsonResultIn.Unmarshal(
		it)

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

func (it *Result) CloneError() error {
	if it.HasError() {
		return errors.New(it.Error.Error())
	}

	return nil
}

func (it *Result) Ptr() *Result {
	return it
}

func (it Result) NonPtr() Result {
	return it
}

func (it *Result) IsEqualPtr(another *Result) bool {
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

	if it.TypeName != another.TypeName {
		return false
	}

	if it.jsonString != nil && another.jsonString != nil &&
		it.jsonString == another.jsonString {
		return true
	}

	return bytes.Equal(it.Bytes, another.Bytes)
}

func (it *Result) CombineErrorWithRef(references ...string) string {
	if it.IsEmptyError() {
		return ""
	}

	csv := csvinternal.StringsToStringDefault(references...)

	return fmt.Sprintf(
		constants.MessageReferenceWrap,
		it.Error.Error(),
		csv)
}

func (it *Result) CombineErrorWithRefError(references ...string) error {
	if it.IsEmptyError() {
		return nil
	}

	errorString := it.CombineErrorWithRef(
		references...)

	return errors.New(errorString)
}

func (it Result) IsEqual(another Result) bool {
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

func (it *Result) BytesError() *coredata.BytesError {
	if it == nil {
		return nil
	}

	return &coredata.BytesError{
		Bytes: it.Bytes,
		Error: it.Error,
	}
}

func (it *Result) Dispose() {
	if it == nil {
		return
	}

	it.Error = nil
	it.Bytes = nil
	it.TypeName = ""
	it.jsonString = nil
}

func (it Result) CloneIf(isClone, isDeepClone bool) Result {
	if isClone {
		return it.Clone(isDeepClone)
	}

	return it
}

func (it *Result) ClonePtr(isDeepClone bool) *Result {
	if it == nil {
		return nil
	}

	cloned := it.Clone(isDeepClone)

	return &cloned
}

func (it Result) Clone(isDeepClone bool) Result {
	if it.Length() == 0 {
		return New([]byte{}, it.CloneError(), it.TypeName)
	}

	if !isDeepClone || it.Length() == 0 {
		return New(it.Bytes, it.CloneError(), it.TypeName)
	}

	newBytes := make([]byte, it.Length())
	copy(newBytes, it.Bytes)

	return New(newBytes, it.CloneError(), it.TypeName)
}
