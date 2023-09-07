package corestr

import (
	"encoding/json"
	"math"
	"regexp"
	"strconv"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/constants/bitsize"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreindexes"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/strutilinternal"
	"gitlab.com/auk-go/core/issetter"
)

type SimpleStringOnce struct {
	value        string
	isInitialize bool
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

// Invalidate
//
// Will make initialize to false
func (it *SimpleStringOnce) Invalidate() {
	it.isInitialize = false
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
			AlreadyInitializedType.
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
	return strutilinternal.IsEmptyOrWhitespace(it.value)
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
	return it.Boolean(false)
}

func (it *SimpleStringOnce) SafeValue() string {
	if it.IsInitialized() {
		return it.value
	}

	return constants.EmptyString
}

func (it *SimpleStringOnce) Uint16() (val uint16, isInRange bool) {
	toUint16, isInRange := it.WithinRange(
		true,
		constants.Zero,
		math.MaxUint16)

	return uint16(toUint16), isInRange
}

func (it *SimpleStringOnce) Uint32() (val uint32, isInRange bool) {
	converted, isInRange := it.WithinRange(
		true,
		constants.Zero,
		math.MaxInt)

	return uint32(converted), isInRange
}

func (it *SimpleStringOnce) WithinRangeDefault(
	min, max int,
) (val int, isInRange bool) {
	return it.WithinRange(
		true,
		min,
		max)
}

func (it *SimpleStringOnce) WithinRange(
	isUsageMinMaxBoundary bool,
	min, max int,
) (val int, isInRange bool) {
	toInt, err := strconv.Atoi(it.value)

	if err != nil {
		return constants.Zero, false
	}

	if toInt >= min && toInt <= max {
		return toInt, true
	}

	if !isUsageMinMaxBoundary {
		return toInt, false
	}

	if toInt < min {
		return min, false
	}

	if toInt > max {
		return max, false
	}

	return constants.Zero, false
}

func (it *SimpleStringOnce) Int() int {
	toInt, err := strconv.Atoi(it.value)

	if err != nil {
		return constants.Zero
	}

	return toInt
}

func (it *SimpleStringOnce) Byte() byte {
	toInt, err := strconv.Atoi(it.value)

	if err != nil {
		return constants.Zero
	}

	if toInt >= constants.Zero && toInt <= constants.MaxUnit8AsInt {
		return byte(toInt)
	}

	return constants.Zero
}

func (it *SimpleStringOnce) Int16() int16 {
	toInt, err := strconv.Atoi(it.value)

	if err != nil {
		return constants.Zero
	}

	if toInt >= math.MinInt16 && toInt <= constants.MaxInt16AsInt {
		return int16(toInt)
	}

	return constants.Zero
}

func (it *SimpleStringOnce) Int32() int32 {
	toInt, err := strconv.Atoi(it.value)

	if err != nil {
		return constants.Zero
	}

	if toInt >= math.MinInt32 && toInt <= math.MaxInt32 {
		return int32(toInt)
	}

	return constants.Zero
}

func (it *SimpleStringOnce) BooleanDefault() bool {
	return it.Boolean(true)
}

func (it *SimpleStringOnce) Boolean(isConsiderInit bool) bool {
	if isConsiderInit && it.IsUnInit() {
		return false
	}

	value := it.value

	if value == "yes" || value == "y" || value == "1" || value == "YES" || value == "Y" {
		return true
	}

	parsedBool, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}

	return parsedBool
}

func (it *SimpleStringOnce) IsSetter(isConsiderInit bool) issetter.Value {
	if isConsiderInit && it.IsUnInit() {
		return issetter.False
	}

	value := it.value

	if value == "yes" || value == "y" || value == "1" || value == "YES" || value == "Y" {
		return issetter.True
	}

	parsedBool, err := strconv.ParseBool(value)
	if err != nil {
		return issetter.Uninitialized
	}

	return issetter.GetBool(parsedBool)
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

func (it SimpleStringOnce) NonPtr() SimpleStringOnce {
	return it
}

func (it *SimpleStringOnce) Ptr() *SimpleStringOnce {
	return it
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

func (it *SimpleStringOnce) LinesSimpleSlice() *SimpleSlice {
	lines := strings.Split(it.value, constants.DefaultLine)

	return New.SimpleSlice.Direct(false, lines)
}

func (it *SimpleStringOnce) SimpleSlice(
	sep string,
) *SimpleSlice {
	lines := strings.Split(it.value, sep)

	return New.SimpleSlice.Direct(false, lines)
}

func (it *SimpleStringOnce) Split(
	sep string,
) []string {
	return strings.Split(it.value, sep)
}

func (it *SimpleStringOnce) SplitLeftRight(
	sep string,
) (left, right string) {
	splits := strings.SplitN(
		it.String(),
		sep,
		expectedLeftRightLength)

	length := len(splits)
	first := splits[coreindexes.First]

	if length == expectedLeftRightLength {
		return first, splits[coreindexes.Second]
	}

	return first, constants.EmptyString
}

func (it *SimpleStringOnce) SplitLeftRightTrim(
	sep string,
) (left, right string) {
	splits := strings.SplitN(
		it.String(), sep,
		expectedLeftRightLength)

	length := len(splits)
	first := splits[coreindexes.First]

	if length == expectedLeftRightLength {
		return strings.TrimSpace(first), strings.TrimSpace(splits[coreindexes.Second])
	}

	return strings.TrimSpace(first), constants.EmptyString
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
	jsonBytes []byte,
) error {
	var dataModel SimpleStringOnceModel
	err := corejson.Deserialize.UsingBytes(
		jsonBytes, &dataModel)

	if err == nil {
		it.value = dataModel.Value
		it.isInitialize = dataModel.IsInitialize
	}

	return err
}

func (it SimpleStringOnce) Json() corejson.Result {
	return corejson.New(it)
}

func (it SimpleStringOnce) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
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

func (it *SimpleStringOnce) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
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

func (it *SimpleStringOnce) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *SimpleStringOnce) Deserialize(toPtr interface{}) (parsingErr error) {
	return it.JsonPtr().Deserialize(toPtr)
}
