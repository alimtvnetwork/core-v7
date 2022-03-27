package corestr

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/constants/bitsize"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/internal/strutilinternal"
)

type ValidValue struct {
	Value      string
	valueBytes *[]byte
	IsValid    bool
	Message    string
}

func NewValidValueUsingAny(
	isIncludeFieldName bool,
	isValid bool,
	any interface{},
) *ValidValue {
	toString := AnyToString(
		isIncludeFieldName,
		any)

	return &ValidValue{
		Value:   toString,
		IsValid: isValid,
		Message: constants.EmptyString,
	}
}

// NewValidValueUsingAnyAutoValid IsValid to false on nil or Empty string
func NewValidValueUsingAnyAutoValid(
	isIncludeFieldName bool,
	any interface{},
) *ValidValue {
	toString := AnyToString(
		isIncludeFieldName,
		any)

	return &ValidValue{
		Value:   toString,
		IsValid: toString == constants.EmptyString,
		Message: constants.EmptyString,
	}
}

func NewValidValue(value string) *ValidValue {
	return &ValidValue{
		Value:   value,
		IsValid: true,
		Message: constants.EmptyString,
	}
}

func NewValidValueEmpty() *ValidValue {
	return &ValidValue{
		Value:   constants.EmptyString,
		IsValid: true,
		Message: constants.EmptyString,
	}
}

func InvalidValidValueNoMessage() *ValidValue {
	return InvalidValidValue(constants.EmptyString)
}

func InvalidValidValue(message string) *ValidValue {
	return &ValidValue{
		Value:   constants.EmptyString,
		IsValid: false,
		Message: message,
	}
}

func (it *ValidValue) ValueBytesOnce() []byte {
	return *it.ValueBytesOncePtr()
}

func (it *ValidValue) ValueBytesOncePtr() *[]byte {
	if it.valueBytes == nil {
		valueBytes := []byte(it.Value)

		it.valueBytes = &valueBytes
	}

	return it.valueBytes
}

func (it *ValidValue) IsEmpty() bool {
	return it.Value == ""
}

func (it *ValidValue) IsWhitespace() bool {
	return strutilinternal.IsEmptyOrWhitespace(it.Value)
}

func (it *ValidValue) Trim() string {
	return strings.TrimSpace(it.Value)
}

func (it *ValidValue) HasValidNonEmpty() bool {
	return it.IsValid && !it.IsEmpty()
}

func (it *ValidValue) HasValidNonWhitespace() bool {
	return it.IsValid && !it.IsWhitespace()
}

func (it *ValidValue) ValueBool() bool {
	if it.Value == "" {
		return false
	}

	toBool, err := strconv.ParseBool(it.Value)

	if err != nil {
		return false
	}

	return toBool
}

func (it *ValidValue) ValueInt(defaultInteger int) int {
	toInt, err := strconv.Atoi(it.Value)

	if err != nil {
		return defaultInteger
	}

	return toInt
}

func (it *ValidValue) ValueDefInt() int {
	toInt, err := strconv.Atoi(it.Value)

	if err != nil {
		return constants.Zero
	}

	return toInt
}

func (it *ValidValue) ValueByte(defVal byte) byte {
	toInt, err := strconv.Atoi(it.Value)

	if err != nil || toInt > constants.MaxUnit8AsInt {
		return defVal
	}

	return byte(toInt)
}

func (it *ValidValue) ValueDefByte() byte {
	toInt, err := strconv.Atoi(it.Value)

	if err != nil || toInt > constants.MaxUnit8AsInt {
		return constants.Zero
	}

	return byte(toInt)
}

func (it *ValidValue) ValueFloat64(defVal float64) float64 {
	toFloat, err := strconv.ParseFloat(it.Value, bitsize.Of64)

	if err != nil {
		return defVal
	}

	return toFloat
}

func (it *ValidValue) ValueDefFloat64() float64 {
	return it.ValueFloat64(constants.Zero)
}

// HasSafeNonEmpty receiver.IsValid &&
//		!receiver.IsLeftEmpty() &&
//		!receiver.IsMiddleEmpty() &&
//		!receiver.IsRightEmpty()
func (it *ValidValue) HasSafeNonEmpty() bool {
	return it.IsValid &&
		!it.IsEmpty()
}

func (it *ValidValue) Is(val string) bool {
	return it.Value == val
}

// IsAnyOf if length of values are 0 then returns true
func (it *ValidValue) IsAnyOf(values ...string) bool {
	if len(values) == 0 {
		return true
	}

	for _, value := range values {
		if it.Value == value {
			return true
		}
	}

	return false
}

func (it *ValidValue) IsContains(val string) bool {
	return strings.Contains(it.Value, val)
}

// IsAnyContains if length of values are 0 then returns true
func (it *ValidValue) IsAnyContains(values ...string) bool {
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

func (it *ValidValue) IsEqualNonSensitive(val string) bool {
	return strings.EqualFold(it.Value, val)
}

func (it *ValidValue) IsRegexMatches(regexp *regexp.Regexp) bool {
	if regexp == nil {
		return false
	}

	return regexp.MatchString(it.Value)
}

func (it *ValidValue) RegexFindString(
	regexp *regexp.Regexp,
) string {
	if regexp == nil {
		return constants.EmptyString
	}

	return regexp.FindString(it.Value)
}

func (it *ValidValue) RegexFindAllStringsWithFlag(
	regexp *regexp.Regexp,
	n int,
) (foundItems []string, hasAny bool) {
	if regexp == nil {
		return []string{}, false
	}

	items := regexp.FindAllString(
		it.Value, n)

	return items, len(items) > 0
}

func (it *ValidValue) RegexFindAllStrings(
	regexp *regexp.Regexp,
	n int,
) []string {
	if regexp == nil {
		return []string{}
	}

	return regexp.FindAllString(it.Value, n)
}

func (it *ValidValue) Split(
	sep string,
) []string {
	return strings.Split(it.Value, sep)
}

func (it *ValidValue) SplitNonEmpty(
	sep string,
) []string {
	slice := strings.Split(it.Value, sep)

	nonEmptySlice := make([]string, 0, len(slice))

	for _, item := range slice {
		if item == constants.EmptyString {
			continue
		}

		nonEmptySlice = append(nonEmptySlice, item)
	}

	return slice
}

func (it *ValidValue) SplitTrimNonWhitespace(
	sep string,
) []string {
	slice := strings.Split(it.Value, sep)

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

func (it *ValidValue) Clone() *ValidValue {
	if it == nil {
		return nil
	}

	return &ValidValue{
		Value:   it.Value,
		IsValid: it.IsValid,
		Message: it.Message,
	}
}

func (it *ValidValue) String() string {
	if it == nil {
		return constants.EmptyString
	}

	return it.Value
}

func (it *ValidValue) FullString() string {
	if it == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		*it)
}

func (it *ValidValue) Clear() {
	if it == nil {
		return
	}

	it.Value = ""
	it.valueBytes = nil
	it.IsValid = false
	it.Message = ""
}

func (it *ValidValue) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
}

func (it ValidValue) Json() corejson.Result {
	return corejson.New(it)
}

func (it ValidValue) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *ValidValue) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*ValidValue, error) {
	err := jsonResult.Deserialize(it)

	if err == nil {
		return it, err
	}

	// has err
	return nil, err
}

func (it *ValidValue) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *ValidValue) Deserialize(toPtr interface{}) (parsingErr error) {
	return it.JsonPtr().Deserialize(toPtr)
}
