package corestr

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/constants/bitsize"
	"gitlab.com/evatix-go/core/internal/utilstringinternal"
	"gitlab.com/evatix-go/core/msgtype"
)

type SimpleStringOnce struct {
	value         string
	isInitialized bool
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
		value:         toString,
		isInitialized: isInitialize,
	}
}

func NewSimpleStringOnce(
	value string,
	isInitialize bool,
) SimpleStringOnce {
	return SimpleStringOnce{
		value:         value,
		isInitialized: isInitialize,
	}
}

func NewSimpleStringOncePtr(
	value string,
	isInitialize bool,
) *SimpleStringOnce {
	return &SimpleStringOnce{
		value:         value,
		isInitialized: isInitialize,
	}
}

func EmptySimpleStringOnce() SimpleStringOnce {
	return SimpleStringOnce{}
}

func (it *SimpleStringOnce) Value() string {
	return it.value
}

func (it *SimpleStringOnce) IsInitialized() bool {
	return it.isInitialized
}

// IsUnInit Not initialized yet
func (it *SimpleStringOnce) IsUnInit() bool {
	return !it.isInitialized
}

func (it *SimpleStringOnce) IsInvalid() bool {
	return !it.isInitialized || it.value == ""
}

func (it *SimpleStringOnce) ValueBytes() []byte {
	return []byte(it.value)
}

func (it *SimpleStringOnce) ValueBytesPtr() *[]byte {
	allBytes := []byte(it.value)

	return &allBytes
}

func (it *SimpleStringOnce) SetOnUninitializedError(setVal string) error {
	if it.isInitialized {
		return msgtype.
			AlreadyInitialized.
			Error("cannot set :"+setVal, it.value)
	}

	it.value = setVal
	it.SetInitialized()

	return nil
}

func (it *SimpleStringOnce) GetPlusSetOnUninitialized(
	setValPlusGetOnUnInit string,
) (valGet string) {
	if it.isInitialized {
		return it.value
	}

	it.value = setValPlusGetOnUnInit
	it.SetInitialized()

	return it.value
}

func (it *SimpleStringOnce) GetPlusSetOnUninitializedFunc(
	setValPlusGetOnUnInitFunc func() string,
) (valGet string) {
	if it.isInitialized {
		return it.value
	}

	it.value = setValPlusGetOnUnInitFunc()
	it.SetInitialized()

	return it.value
}

func (it *SimpleStringOnce) SetOnUninitialized(setVal string) (isSet bool) {
	if it.isInitialized {
		return false
	}

	it.value = setVal
	it.SetInitialized()

	return true
}

func (it *SimpleStringOnce) SetInitialized() {
	it.isInitialized = true
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
	return it.isInitialized && !it.IsEmpty()
}

func (it *SimpleStringOnce) HasValidNonWhitespace() bool {
	return it.isInitialized && !it.IsWhitespace()
}

func (it *SimpleStringOnce) ValueBool() bool {
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
//      it.isInitialized &&
//		!it.IsEmpty()
func (it *SimpleStringOnce) HasSafeNonEmpty() bool {
	return it.isInitialized &&
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

func (it *SimpleStringOnce) Clone() *SimpleStringOnce {
	if it == nil {
		return nil
	}

	return &SimpleStringOnce{
		value:         it.value,
		isInitialized: it.isInitialized,
	}
}

func (it *SimpleStringOnce) String() string {
	if it == nil {
		return constants.EmptyString
	}

	return it.value
}

func (it *SimpleStringOnce) FullString() string {
	if it == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		*it)
}
