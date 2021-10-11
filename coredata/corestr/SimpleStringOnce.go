package corestr

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/constants/bitsize"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/utilstringinternal"
)

type SimpleStringOnce struct {
	value        string
	isInitialize bool
}

func NewSimpleStringOnceAny(
	isIncludeFieldNames bool,
	value interface{},
	isInitialize bool,
) SimpleStringOnce {
	toString := AnyToString(
		isIncludeFieldNames,
		value)

	return SimpleStringOnce{
		value:        toString,
		isInitialize: isInitialize,
	}
}

func NewSimpleStringOnceInitialized(
	value string,
) SimpleStringOnce {
	return SimpleStringOnce{
		value:        value,
		isInitialize: true,
	}
}

func NewSimpleStringOnceInitializedPtr(
	value string,
) *SimpleStringOnce {
	return &SimpleStringOnce{
		value:        value,
		isInitialize: true,
	}
}

func NewSimpleStringOnce(
	value string,
	isInitialize bool,
) SimpleStringOnce {
	return SimpleStringOnce{
		value:        value,
		isInitialize: isInitialize,
	}
}

func NewSimpleStringOncePtr(
	value string,
	isInitialize bool,
) *SimpleStringOnce {
	return &SimpleStringOnce{
		value:        value,
		isInitialize: isInitialize,
	}
}

func EmptySimpleStringOnce() SimpleStringOnce {
	return SimpleStringOnce{}
}

func (it *SimpleStringOnce) Value() string {
	return it.value
}

func (it *SimpleStringOnce) IsInitialized() bool {
	return it.isInitialize
}

// IsUnInit Not initialized yet
//
// !it.isInitialize
func (it *SimpleStringOnce) IsUnInit() bool {
	return !it.isInitialize
}

// IsInvalid
//
//  !it.isInitialize || it.value == ""
func (it *SimpleStringOnce) IsInvalid() bool {
	return !it.isInitialize || it.value == ""
}

func (it *SimpleStringOnce) ValueBytes() []byte {
	return []byte(it.value)
}

func (it *SimpleStringOnce) ValueBytesPtr() *[]byte {
	allBytes := []byte(it.value)

	return &allBytes
}

func (it *SimpleStringOnce) SetOnUninitializedError(setVal string) error {
	if it.isInitialize {
		return errcore.
			AlreadyInitialized.
			Error("cannot set :"+setVal, it.value)
	}

	it.value = setVal
	it.SetInitialize()

	return nil
}

func (it *SimpleStringOnce) GetPlusSetOnUninitialized(
	setValPlusGetOnUnInit string,
) (valGet string) {
	if it.isInitialize {
		return it.value
	}

	it.value = setValPlusGetOnUnInit
	it.SetInitialize()

	return it.value
}

func (it *SimpleStringOnce) GetPlusSetEmptyOnUninitialized() (valGet string) {
	if it.isInitialize {
		return it.value
	}

	it.value = constants.EmptyString
	it.SetInitialize()

	return it.value
}

func (it *SimpleStringOnce) GetPlusSetOnUninitializedFunc(
	setValPlusGetOnUnInitFunc func() string,
) (valGet string) {
	if it.isInitialize {
		return it.value
	}

	it.value = setValPlusGetOnUnInitFunc()
	it.SetInitialize()

	return it.value
}

func (it *SimpleStringOnce) SetOnUninitialized(setVal string) (isSet bool) {
	if it.isInitialize {
		return false
	}

	it.value = setVal
	it.SetInitialize()

	return true
}

func (it *SimpleStringOnce) SetInitialize() {
	it.isInitialize = true
}

func (it *SimpleStringOnce) SetUnInit() {
	it.isInitialize = false
}

func (it *SimpleStringOnce) ConcatNew(appendingText string) SimpleStringOnce {
	return SimpleStringOnce{
		value:        it.value + appendingText,
		isInitialize: it.isInitialize,
	}
}

func (it *SimpleStringOnce) ConcatNewUsingStrings(
	joiner string,
	appendingTexts ...string,
) SimpleStringOnce {
	slice := append([]string{it.value}, appendingTexts...)

	return SimpleStringOnce{
		value:        strings.Join(slice, joiner),
		isInitialize: it.isInitialize,
	}
}

func (it *SimpleStringOnce) IsEmpty() bool {
	return it.value == ""
}

func (it *SimpleStringOnce) IsWhitespace() bool {
	return utilstringinternal.IsEmptyOrWhitespace(it.value)
}

func (it *SimpleStringOnce) Trim() string {
	return strings.TrimSpace(it.value)
}

func (it *SimpleStringOnce) HasValidNonEmpty() bool {
	return it.isInitialize && !it.IsEmpty()
}

func (it *SimpleStringOnce) HasValidNonWhitespace() bool {
	return it.isInitialize && !it.IsWhitespace()
}

func (it *SimpleStringOnce) IsValueBool() bool {
	if it.value == "" {
		return false
	}

	toBool, err := strconv.ParseBool(it.value)

	if err != nil {
		return false
	}

	return toBool
}

func (it *SimpleStringOnce) ValueInt(defaultInteger int) int {
	toInt, err := strconv.Atoi(it.value)

	if err != nil {
		return defaultInteger
	}

	return toInt
}

func (it *SimpleStringOnce) ValueDefInt() int {
	toInt, err := strconv.Atoi(it.value)

	if err != nil {
		return constants.Zero
	}

	return toInt
}

func (it *SimpleStringOnce) ValueByte(defVal byte) byte {
	toInt, err := strconv.Atoi(it.value)

	if err != nil || toInt > constants.MaxUnit8AsInt {
		return defVal
	}

	return byte(toInt)
}

func (it *SimpleStringOnce) ValueDefByte() byte {
	toInt, err := strconv.Atoi(it.value)

	if err != nil || toInt > constants.MaxUnit8AsInt {
		return constants.Zero
	}

	return byte(toInt)
}

func (it *SimpleStringOnce) ValueFloat64(defVal float64) float64 {
	toFloat, err := strconv.ParseFloat(it.value, bitsize.Of64)

	if err != nil {
		return defVal
	}

	return toFloat
}

func (it *SimpleStringOnce) ValueDefFloat64() float64 {
	return it.ValueFloat64(constants.Zero)
}

// HasSafeNonEmpty
//      it.isInitialize &&
//		!it.IsEmpty()
func (it *SimpleStringOnce) HasSafeNonEmpty() bool {
	return it.isInitialize &&
		!it.IsEmpty()
}

func (it *SimpleStringOnce) Is(val string) bool {
	return it.value == val
}

// IsAnyOf if length of values are 0 then returns true
func (it *SimpleStringOnce) IsAnyOf(values ...string) bool {
	if len(values) == 0 {
		return true
	}

	for _, value := range values {
		if it.value == value {
			return true
		}
	}

	return false
}

func (it *SimpleStringOnce) IsContains(val string) bool {
	return strings.Contains(it.value, val)
}

// IsAnyContains if length of values are 0 then returns true
func (it *SimpleStringOnce) IsAnyContains(values ...string) bool {
	if len(values) == 0 {
		return true
	}

	for _, value := range values {
		if it.IsContains(value) {
			return true
		}
	}

	return false
}

func (it *SimpleStringOnce) IsEqualNonSensitive(val string) bool {
	return strings.EqualFold(it.value, val)
}

func (it *SimpleStringOnce) IsRegexMatches(regexp *regexp.Regexp) bool {
	if regexp == nil {
		return false
	}

	return regexp.MatchString(it.value)
}

func (it *SimpleStringOnce) RegexFindString(
	regexp *regexp.Regexp,
) string {
	if regexp == nil {
		return constants.EmptyString
	}

	return regexp.FindString(it.value)
}

func (it *SimpleStringOnce) RegexFindAllStringsWithFlag(
	regexp *regexp.Regexp,
	n int,
) (foundItems []string, hasAny bool) {
	if regexp == nil {
		return []string{}, false
	}

	items := regexp.FindAllString(
		it.value, n)

	return items, len(items) > 0
}

func (it *SimpleStringOnce) RegexFindAllStrings(
	regexp *regexp.Regexp,
	n int,
) []string {
	if regexp == nil {
		return []string{}
	}

	return regexp.FindAllString(it.value, n)
}

func (it *SimpleStringOnce) Split(
	sep string,
) []string {
	return strings.Split(it.value, sep)
}

func (it *SimpleStringOnce) SplitNonEmpty(
	sep string,
) []string {
	slice := strings.Split(it.value, sep)

	nonEmptySlice := make([]string, 0, len(slice))

	for _, item := range slice {
		if item == constants.EmptyString {
			continue
		}

		nonEmptySlice = append(nonEmptySlice, item)
	}

	return slice
}

func (it *SimpleStringOnce) SplitTrimNonWhitespace(
	sep string,
) []string {
	slice := strings.Split(it.value, sep)

	nonEmptySlice := make([]string, 0, len(slice))

	for _, item := range slice {
		itemTrimmed := strings.TrimSpace(item)
		if itemTrimmed == constants.EmptyString {
			continue
		}

		nonEmptySlice = append(nonEmptySlice, itemTrimmed)
	}

	return slice
}

func (it *SimpleStringOnce) ClonePtr() *SimpleStringOnce {
	if it == nil {
		return nil
	}

	return &SimpleStringOnce{
		value:        it.value,
		isInitialize: it.isInitialize,
	}
}

func (it SimpleStringOnce) Clone() SimpleStringOnce {
	return SimpleStringOnce{
		value:        it.value,
		isInitialize: it.isInitialize,
	}
}

func (it SimpleStringOnce) CloneUsingNewVal(val string) SimpleStringOnce {
	return SimpleStringOnce{
		value:        val,
		isInitialize: it.isInitialize,
	}
}

func (it *SimpleStringOnce) Dispose() {
	if it == nil {
		return
	}

	it.value = constants.EmptyString
	it.isInitialize = true
}

func (it *SimpleStringOnce) String() string {
	if it == nil {
		return constants.EmptyString
	}

	return it.value
}

func (it *SimpleStringOnce) StringPtr() *string {
	if it == nil {
		emptyString := ""
		return &emptyString
	}

	return &it.value
}

func (it *SimpleStringOnce) JsonModel() SimpleStringOnceModel {
	return SimpleStringOnceModel{
		Value:        it.Value(),
		IsInitialize: it.IsInitialized(),
	}
}

func (it *SimpleStringOnce) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *SimpleStringOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *SimpleStringOnce) UnmarshalJSON(
	data []byte,
) error {
	var dataModel SimpleStringOnceModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.value = dataModel.Value
		it.isInitialize = dataModel.IsInitialize
	}

	return err
}

func (it SimpleStringOnce) Json() corejson.Result {
	return corejson.NewFromAny(it)
}

func (it SimpleStringOnce) JsonPtr() *corejson.Result {
	return corejson.NewFromAnyPtr(it)
}

func (it *SimpleStringOnce) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*SimpleStringOnce, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *SimpleStringOnce) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *SimpleStringOnce {
	parsedResult, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return parsedResult
}

func (it *SimpleStringOnce) AsJsoner() corejson.Jsoner {
	return it
}

func (it *SimpleStringOnce) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *SimpleStringOnce) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *SimpleStringOnce) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}
