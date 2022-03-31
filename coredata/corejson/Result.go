package corejson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/defaulterr"
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

func (it *Result) Map() map[string]string {
	if it == nil {
		return map[string]string{}
	}

	newMap := make(
		map[string]string,
		constants.Capacity3)

	if len(it.Bytes) > 0 {
		newMap[bytesFieldName] = it.JsonString()
	}

	if it.Error != nil {
		newMap[errorFieldName] = it.Error.Error()
	}

	if it.TypeName != "" {
		newMap[typeFieldName] = it.TypeName
	}

	return newMap
}

func (it *Result) DeserializedFieldsToMap() (
	fieldsMap map[string]interface{},
	parsingErr error,
) {
	if it == nil || len(it.Bytes) == 0 {
		return map[string]interface{}{}, nil
	}

	parsingErr = it.Deserialize(fieldsMap)

	return fieldsMap, parsingErr
}

// SafeDeserializedFieldsToMap
//
// Warning:
//  - Swallows the error
func (it *Result) SafeDeserializedFieldsToMap() (
	fieldsMap map[string]interface{},
) {
	fieldsMap, _ = it.DeserializedFieldsToMap()

	return fieldsMap
}

func (it *Result) FieldsNames() (
	fieldsNames []string,
	parsingErr error,
) {
	fieldsMap, parsingErr := it.DeserializedFieldsToMap()

	if len(fieldsMap) == 0 {
		return []string{}, parsingErr
	}

	fieldsNames = make([]string, len(fieldsMap))
	index := 0

	for fieldNameKey := range fieldsMap {
		fieldsNames[index] = fieldNameKey

		index++
	}

	return fieldsNames, parsingErr
}

// SafeFieldsNames
//
// Warning:
//  - Swallows the error
func (it *Result) SafeFieldsNames() (
	fieldsNames []string,
) {
	fieldsNames, _ = it.FieldsNames()

	return fieldsNames
}

func (it *Result) BytesTypeName() string {
	if it == nil {
		return ""
	}

	return it.TypeName
}

func (it *Result) SafeBytesTypeName() string {
	if it.IsEmpty() {
		return ""
	}

	return it.TypeName
}

func (it *Result) SafeString() string {
	return *it.JsonStringPtr()
}

func (it *Result) JsonString() string {
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

func (it *Result) PrettyJsonBuffer(prefix, indent string) (*bytes.Buffer, error) {
	var prettyJSON bytes.Buffer

	if it.IsEmpty() {
		return &prettyJSON, nil
	}

	err := json.Indent(
		&prettyJSON,
		it.Bytes,
		prefix,
		indent)

	return &prettyJSON, err
}

func (it *Result) PrettyJsonString() string {
	if it == nil || it.IsEmptyJson() {
		return ""
	}

	prettyJSON, err := it.PrettyJsonBuffer(
		constants.EmptyString,
		constants.Tab)

	if err != nil {
		return ""
	}

	return prettyJSON.String()
}

func (it *Result) PrettyJsonStringOrErrString() string {
	if it == nil {
		return "json result: nil cannot have json string"
	}

	if it.HasError() {
		return it.MeaningfulError().Error()
	}

	return it.PrettyJsonString()
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

func (it Result) String() string {
	if it.IsAnyNull() {
		return constants.EmptyString
	}

	var currentMap map[string]string

	if it.HasError() {
		currentMap = map[string]string{
			"Json":  it.JsonString(),
			"Type":  it.TypeName,
			"Error": it.MeaningfulErrorMessage(),
		}
	} else {
		currentMap = map[string]string{
			"Json": it.JsonString(),
			"Type": it.TypeName,
		}
	}

	toString := fmt.Sprintf(
		constants.SprintValueFormat,
		currentMap)

	currentMap = nil

	return toString
}

func (it *Result) SafeNonIssueBytes() []byte {
	if it.HasIssuesOrEmpty() {
		return []byte{}
	}

	return it.Bytes
}

func (it *Result) SafeBytes() []byte {
	if it.IsAnyNull() {
		return []byte{}
	}

	return it.Bytes
}

func (it *Result) Values() []byte {
	return it.Bytes
}

func (it *Result) SafeValues() []byte {
	if it.IsAnyNull() {
		return []byte{}
	}

	return it.Bytes
}

func (it *Result) SafeValuesPtr() *[]byte {
	if it.HasIssuesOrEmpty() {
		return &[]byte{}
	}

	return &it.Bytes
}

func (it *Result) Raw() ([]byte, error) {
	if it == nil {
		return []byte{}, defaulterr.JsonResultNull
	}

	return it.SafeBytes(), it.MeaningfulError()
}

func (it *Result) RawMust() []byte {
	allBytes, err := it.Raw()
	errcore.HandleErr(err)

	return allBytes
}

func (it *Result) RawString() (jsonString string, err error) {
	return it.JsonString(), it.MeaningfulError()
}

func (it *Result) RawStringMust() (jsonString string) {
	jsonString, err := it.RawString()

	if err != nil {
		panic(err)
	}

	return jsonString
}

func (it *Result) RawErrString() (rawJsonBytes []byte, errorMsg string) {
	return it.Bytes, it.MeaningfulErrorMessage()
}

func (it *Result) RawPrettyString() (jsonString string, err error) {
	return it.PrettyJsonString(), it.MeaningfulError()
}

func (it *Result) MeaningfulErrorMessage() string {
	err := it.MeaningfulError()

	if err == nil {
		return ""
	}

	return err.Error()
}

// MeaningfulError
//
//  create error even if results are nil.
func (it *Result) MeaningfulError() error {
	if it == nil {
		return defaulterr.JsonResultNull
	}

	if it.Error == nil && len(it.Bytes) > 0 {
		// everything is okay

		return nil
	}

	if it.IsEmptyJsonBytes() {
		// error may or may not exist
		errMsg := errcore.BytesAreNilOrEmptyType.String() +
			" Additional: " +
			errcore.ToString(it.Error) + // error may or may not exist
			", type:"

		return errcore.
			FailedToParseType.
			Error(errMsg, it.TypeName)
	}

	// must error and payload may or may not exist
	return errcore.
		FailedToParseType.
		Error(
			errcore.ToString(it.Error)+", type:"+it.TypeName+", payload:",
			it.safeJsonStringInternal())
}

func (it *Result) safeJsonStringInternal() string {
	if it == nil {
		return ""
	}

	var safeJsonString string
	if it != nil && len(it.Bytes) > 0 {
		safeJsonString = string(it.Bytes)
	}

	return safeJsonString
}

func (it *Result) IsEmptyError() bool {
	return it == nil || it.Error == nil
}

// HasSafeItems
//
//  Returns true if
//  Result is not null
//  and has NO error
//  and has non-Empty json (other than length 0 or "{}")
//
// Invert of HasIssuesOrEmpty
func (it *Result) HasSafeItems() bool {
	return !it.HasIssuesOrEmpty()
}

// IsAnyNull
//
//  Returns true
//  if Result is null
//  or Bytes is null
func (it *Result) IsAnyNull() bool {
	return it == nil || it.Bytes == nil
}

// HasIssuesOrEmpty
//
//  Returns true
//  if Result is null
//  or has any error
//  or has empty json (length 0 or "{}")
//
// Result.IsAnyNull() ||
// Result.HasError() ||
// Result.IsEmptyJsonBytes()
func (it *Result) HasIssuesOrEmpty() bool {
	return it == nil || it.Error != nil || it.IsEmptyJsonBytes()
}

func (it *Result) HandleError() {
	if it.HasIssuesOrEmpty() {
		panic(it.MeaningfulError())
	}
}

// MustBeSafe alias for HandleError
func (it *Result) MustBeSafe() {
	if it.HasIssuesOrEmpty() {
		panic(it.MeaningfulError())
	}
}

func (it *Result) HandleErrorWithMsg(msg string) {
	if it.HasIssuesOrEmpty() {
		panic(msg + constants.DefaultLine + it.MeaningfulErrorMessage())
	}
}

// HasBytes
//
// Invert of Result.IsEmptyJsonBytes()
//  Represents has at least valid json data other than length 0 or "{}"
func (it *Result) HasBytes() bool {
	return !it.IsEmptyJsonBytes()
}

// HasJsonBytes
//
// Invert of Result.IsEmptyJsonBytes()
//  Represents has at least valid json data other than length 0 or "{}"
func (it *Result) HasJsonBytes() bool {
	return !it.IsEmptyJsonBytes()
}

// IsEmptyJsonBytes
//
// len == 0, nil, "{}" returns as empty true
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

func (it *Result) IsEmpty() bool {
	return it == nil || len(it.Bytes) == 0
}

func (it Result) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *Result) IsEmptyJson() bool {
	return it.IsEmptyJsonBytes()
}

// HasJson
//
// Invert of Result.IsEmptyJsonBytes()
func (it *Result) HasJson() bool {
	return !it.IsEmptyJsonBytes()
}

func (it *Result) InjectInto(
	injector JsonParseSelfInjector,
) error {
	return injector.JsonParseSelfInject(it)
}

// Deserialize
//
// Same as Unmarshal, just alias
func (it *Result) Deserialize(
	anyPointer interface{},
) error {
	return it.Unmarshal(anyPointer)
}

// DeserializeMust
//
// Same as UnmarshalMust, just alias
func (it *Result) DeserializeMust(
	anyPointer interface{},
) {
	err := it.Unmarshal(anyPointer)

	if err != nil {
		panic(err)
	}
}

func (it *Result) UnmarshalMust(
	anyPointer interface{},
) {
	err := it.Unmarshal(anyPointer)

	if err != nil {
		panic(err)
	}
}

// Unmarshal
//
//  deserializes current safe bytes to given pointer
func (it *Result) Unmarshal(
	anyPointer interface{},
) error {
	if it == nil {
		return errcore.
			UnMarshallingFailedType.
			Error(
				"cannot unmarshal null json result, to pointer type",
				reflectinternal.TypeName(anyPointer))
	}

	if it.HasError() {
		compiledMessage := errcore.MessageVarMap(
			"json unmarshal failed with existing error",
			map[string]interface{}{
				"err":     it.Error,
				"src":     it.TypeName,
				"dst":     reflectinternal.TypeName(anyPointer),
				"payload": it.safeJsonStringInternal(),
			})

		return errcore.
			UnMarshallingFailedType.
			ErrorNoRefs(compiledMessage)
	}

	err := json.Unmarshal(
		it.Bytes,
		anyPointer)

	if err == nil {
		return nil
	}

	// unmarshal caught error
	compiledMessage := errcore.MessageVarMap(
		"json unmarshal failed",
		map[string]interface{}{
			"err":     it.Error,
			"src":     it.TypeName,
			"dst":     reflectinternal.TypeName(anyPointer),
			"payload": it.safeJsonStringInternal(),
		})

	return errcore.
		UnMarshallingFailedType.
		ErrorNoRefs(compiledMessage)
}

// SerializeSkipExistingIssues
//
// Ignores and returns nil if HasIssuesOrEmpty satisfied
func (it *Result) SerializeSkipExistingIssues() (
	[]byte, error,
) {
	if it.HasIssuesOrEmpty() {
		return nil, nil
	}

	return it.serializeInternal()
}

func (it *Result) serializeInternal() (
	[]byte, error,
) {
	rawBytes, err := json.Marshal(it)

	if err == nil {
		return rawBytes, nil
	}

	// has error
	reference := errcore.VarTwoNoType(
		"marshal or serialize Error", err.Error(),
		"src", it.TypeName,
	)

	return nil, errcore.
		Serialize.
		ErrorRefOnly(reference)
}

func (it *Result) Serialize() ([]byte, error) {
	if it == nil {
		return nil, errcore.
			Serialize.
			ErrorNoRefs("cannot marshal if JsonResult is null")
	}

	if it.Error != nil {
		return []byte{}, it.MeaningfulError()
	}

	return it.serializeInternal()
}

func (it *Result) SerializeMust() []byte {
	rs, err := it.Serialize()
	errcore.MustBeEmpty(err)

	return rs
}

// UnmarshalSkipExistingIssues
//
// Ignores and returns nil if HasIssuesOrEmpty satisfied
func (it *Result) UnmarshalSkipExistingIssues(
	toPointer interface{},
) error {
	if it.HasIssuesOrEmpty() {
		return nil
	}

	err := json.Unmarshal(it.Bytes, toPointer)

	if err == nil {
		return nil
	}

	// unmarshal caught error
	compiledMessage := errcore.MessageVarMap(
		"json unmarshal failed",
		map[string]interface{}{
			"err":     err,
			"src":     it.TypeName,
			"dst":     reflectinternal.TypeName(toPointer),
			"payload": it.safeJsonStringInternal(),
		})

	return errcore.
		UnMarshallingFailedType.
		ErrorNoRefs(compiledMessage)
}

func (it *Result) UnmarshalResult() (*Result, error) {
	empty := Empty.ResultPtr()
	err := it.Unmarshal(empty)

	return empty, err
}

//goland:noinspection GoLinterLocal
func (it *Result) JsonModel() Result {
	if it == nil {
		return Result{
			Error: defaulterr.JsonResultNull,
		}
	}

	return *it
}

//goland:noinspection GoLinterLocal
func (it *Result) JsonModelAny() interface{} {
	return it.JsonModel()
}

// Json
//
//  creates json result of self
func (it Result) Json() Result {
	return NewResult.Any(it)
}

// JsonPtr
//
//  creates json result of self
func (it Result) JsonPtr() *Result {
	return NewResult.AnyPtr(it)
}

// ParseInjectUsingJson It will not update the self but creates a new one.
func (it *Result) ParseInjectUsingJson(
	jsonResultIn *Result,
) (*Result, error) {
	err := jsonResultIn.Unmarshal(
		it)

	if err != nil {
		return Empty.ResultPtrWithErr(it.TypeName, err), err
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

func (it *Result) CloneError() error {
	if it.HasError() {
		return errors.New(it.Error.Error())
	}

	return nil
}

func (it Result) Ptr() *Result {
	return &it
}

func (it *Result) NonPtr() Result {
	if it == nil {
		return Result{
			Error: errors.New("nil json result"),
		}
	}

	return *it
}

func (it Result) ToPtr() *Result {
	return &it
}

func (it Result) ToNonPtr() Result {
	return it
}

func (it *Result) IsEqualPtr(another *Result) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it == another {
		return true
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

func (it *Result) CombineErrorWithRefString(references ...string) string {
	if it.IsEmptyError() {
		return ""
	}

	csv := csvinternal.StringsToStringDefault(references...)

	return fmt.Sprintf(
		constants.MessageReferenceWrapFormat,
		it.Error.Error(),
		csv)
}

func (it *Result) CombineErrorWithRefError(references ...string) error {
	if it.IsEmptyError() {
		return nil
	}

	errorString := it.CombineErrorWithRefString(
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
	it.TypeName = constants.EmptyString
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
		return NewResult.Create(
			[]byte{},
			it.CloneError(),
			it.TypeName)
	}

	if !isDeepClone || it.Length() == 0 {
		return NewResult.Create(
			it.Bytes,
			it.CloneError(),
			it.TypeName)
	}

	newBytes := make([]byte, it.Length())
	copy(newBytes, it.Bytes)

	return NewResult.Create(
		newBytes,
		it.CloneError(),
		it.TypeName)
}

func (it Result) AsJsonContractsBinder() JsonContractsBinder {
	return &it
}

func (it Result) AsJsoner() Jsoner {
	return &it
}

func (it Result) JsonParseSelfInject(
	jsonResultIn *Result,
) error {
	_, err := it.ParseInjectUsingJson(jsonResultIn)

	return err
}

func (it Result) AsJsonParseSelfInjector() JsonParseSelfInjector {
	return &it
}
