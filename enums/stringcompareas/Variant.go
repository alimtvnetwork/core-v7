package stringcompareas

import (
	"errors"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreinterface/enuminf"
	"gitlab.com/auk-go/core/errcore"
)

type Variant byte

const (
	Equal Variant = iota
	StartsWith
	EndsWith
	Anywhere
	Contains // alias for Anywhere
	AnyChars // If given search chars is found in the content
	// Regex strings will be cached and
	// compiled using map, mutex
	// will be used to lock,
	// if failed to compile then panic
	Regex
	NotEqual      // invert of Equal
	NotStartsWith // invert of StartsWith
	NotEndsWith   // invert of EndsWith
	NotContains   // invert of Anywhere
	NotAnyChars   // invert of AnyChars
	NotMatchRegex // invert of Regex
	Invalid
)

func (it Variant) Value() byte {
	return byte(it)
}

func (it Variant) IsAnyMethod(methodNames ...string) bool {
	return BasicEnumImpl.IsAnyNamesOf(it.Value(), methodNames...)
}

func (it Variant) AllNameValues() []string {
	return BasicEnumImpl.AllNameValues()
}

func (it Variant) OnlySupportedErr(names ...string) error {
	return BasicEnumImpl.OnlySupportedErr(names...)
}

func (it Variant) OnlySupportedMsgErr(message string, names ...string) error {
	return BasicEnumImpl.OnlySupportedMsgErr(message, names...)
}

func (it Variant) ValueUInt16() uint16 {
	return uint16(it)
}

func (it Variant) IntegerEnumRanges() []int {
	return BasicEnumImpl.IntegerEnumRanges()
}

func (it Variant) MinMaxAny() (min, max interface{}) {
	return BasicEnumImpl.MinMaxAny()
}

func (it Variant) MinValueString() string {
	return BasicEnumImpl.MinValueString()
}

func (it Variant) MaxValueString() string {
	return BasicEnumImpl.MaxValueString()
}

func (it Variant) MaxInt() int {
	return BasicEnumImpl.MaxInt()
}

func (it Variant) MinInt() int {
	return BasicEnumImpl.MinInt()
}

func (it Variant) RangesDynamicMap() map[string]interface{} {
	return BasicEnumImpl.RangesDynamicMap()
}

func (it Variant) IsByteValueEqual(value byte) bool {
	return byte(it) == value
}

func (it Variant) Format(format string) (compiled string) {
	return BasicEnumImpl.Format(format, it)
}

func (it Variant) IsEnumEqual(enum enuminf.BasicEnumer) bool {
	return it.ValueByte() == enum.ValueByte()
}

func (it *Variant) IsAnyEnumsEqual(enums ...enuminf.BasicEnumer) bool {
	for _, enum := range enums {
		if it.IsEnumEqual(enum) {
			return true
		}
	}

	return false
}

func (it Variant) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Variant) IsAnyNamesOf(names ...string) bool {
	for _, name := range names {
		if it.IsNameEqual(name) {
			return true
		}
	}

	return false
}

func (it Variant) IsValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Variant) IsAnyValuesEqual(anyByteValues ...byte) bool {
	for _, currentVal := range anyByteValues {
		if it.IsValueEqual(currentVal) {
			return true
		}
	}

	return false
}

func (it Variant) ValueInt() int {
	return int(it)
}

func (it Variant) ValueInt8() int8 {
	return int8(it)
}

func (it Variant) ValueInt16() int16 {
	return int16(it)
}

func (it Variant) ValueInt32() int32 {
	return int32(it)
}

func (it Variant) ValueString() string {
	return it.ToNumberString()
}

func (it Variant) IsValid() bool {
	return it != Invalid
}

func (it Variant) IsInvalid() bool {
	return it == Invalid
}

func (it Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Variant) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.ValueByte())
}

func (it Variant) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(
		isMappedToDefault,
		jsonUnmarshallingValue)
}

func (it Variant) String() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Variant) Is(compare Variant) bool {
	return it == compare
}

func (it Variant) IsEqual() bool {
	return it == Equal
}

func (it Variant) IsStartsWith() bool {
	return it == StartsWith
}

func (it Variant) IsEndsWith() bool {
	return it == EndsWith
}

func (it Variant) IsAnywhere() bool {
	return it == Anywhere
}

func (it Variant) IsContains() bool {
	return it == Contains
}

func (it Variant) IsAnyChars() bool {
	return it == AnyChars
}

func (it Variant) IsRegex() bool {
	return it == Regex
}

// IsNegativeCondition returns true for any of the cases mentioned in negativeCases
//
//	NotEqual      // invert of Equal
//	NotStartsWith // invert of StartsWith
//	NotEndsWith   // invert of EndsWith
//	NotContains   // invert of Anywhere
//	NotAnyChars   // invert of AnyChars
//	NotMatchRegex // invert of Regex
func (it Variant) IsNegativeCondition() bool {
	for _, negativeCase := range negativeCases {
		if negativeCase == it {
			return true
		}
	}

	return false
}

func (it Variant) IsNotEqual() bool {
	return it == NotEqual
}

func (it Variant) IsNotStartsWith() bool {
	return it == NotStartsWith
}

func (it Variant) IsNotEndsWith() bool {
	return it == NotEndsWith
}

func (it Variant) IsNotContains() bool {
	return it == NotContains
}

func (it Variant) IsNotMatchRegex() bool {
	return it == NotMatchRegex
}

func (it Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.ValueByte())
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	rawScriptType, err := BasicEnumImpl.UnmarshallToValue(
		isMappedToDefault, data)

	if err == nil {
		*it = Variant(rawScriptType)
	}

	return err
}

func (it Variant) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it *Variant) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it *Variant) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it *Variant) ValueByte() byte {
	return byte(*it)
}

func (it *Variant) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

// IsLineCompareFunc for
// Regex case has no use, use regex
// pattern for case sensitive or insensitive search
//
// Functions Mapping:
//  Equal:         isEqualFunc,
//  StartsWith:    isStartsWithFunc,
//  EndsWith:      isEndsWithFunc,
//  Anywhere:      isAnywhereFunc,
//  AnyChars:      isAnyCharsFunc,
//  Contains:      isAnywhereFunc,
//  Regex:         isRegexFunc,
//  NotEqual:      isNotEqualFunc,
//  NotStartsWith: isNotStartsWithFunc,
//  NotEndsWith:   isNotEndsWithFunc,
//  NotContains:   isNotContainsFunc,
//  NotAnyChars:   isNotAnyCharsFunc,
//  NotMatchRegex: isNotMatchRegex,
func (it Variant) IsLineCompareFunc() IsLineCompareFunc {
	return rangesMap[it]
}

func (it Variant) DynamicCompare(
	isDynamicCompareFunc IsDynamicCompareFunc,
	lineNumber int, content string,
) bool {
	return isDynamicCompareFunc(
		lineNumber,
		content,
		it)
}

// IsCompareSuccess
// Regex case has no use,
// use regex pattern for case sensitive or insensitive search
//
// Regex will be cached to map for the syntax,
// if running twice it will not create new but the same one from the map.
// It save the regex into map using mutex lock, so async codes can run.
func (it Variant) IsCompareSuccess(
	isIgnoreCase bool,
	content,
	search string,
) bool {
	return it.IsLineCompareFunc()(
		content,
		search,
		isIgnoreCase)
}

func (it Variant) VerifyMessage(
	isIgnoreCase bool,
	content,
	search string,
) string {
	isMatch := it.IsCompareSuccess(
		isIgnoreCase,
		content,
		search,
	)

	if isMatch {
		return constants.EmptyString
	}

	isIgnoreCaseString := "- {case strict}"

	if isIgnoreCase {
		isIgnoreCaseString = "- {case ignored}"
	}

	if it.IsNegativeCondition() {
		return errcore.ExpectingNotEqualSimpleNoType(
			"CompareMethod \""+it.Name()+"\" - {negative} match failed "+isIgnoreCaseString,
			search,
			content)
	}

	return errcore.ExpectingSimpleNoType(
		"CompareMethod \""+it.Name()+"\" - match failed "+isIgnoreCaseString,
		search,
		content)
}

func (it Variant) VerifyError(
	isIgnoreCase bool,
	content,
	search string,
) error {
	message := it.VerifyMessage(
		isIgnoreCase,
		content,
		search,
	)

	if message == constants.EmptyString {
		return nil
	}

	return errors.New(message)
}

func (it Variant) VerifyMessageCaseSensitive(
	content,
	search string,
) string {
	return it.VerifyMessage(
		false,
		content,
		search,
	)
}

func (it Variant) VerifyErrorCaseSensitive(
	content,
	search string,
) error {
	return it.VerifyError(
		false,
		content,
		search,
	)
}

// IsCompareSuccessCaseSensitive for
// Regex case has no use, use regex
// pattern for case sensitive or insensitive search
func (it *Variant) IsCompareSuccessCaseSensitive(content, search string) bool {
	return it.IsLineCompareFunc()(
		content,
		search,
		false)
}

// IsCompareSuccessNonCaseSensitive for
// Regex case has no use, use regex
// pattern for case sensitive or insensitive search
func (it *Variant) IsCompareSuccessNonCaseSensitive(content, search string) bool {
	return it.IsLineCompareFunc()(
		content,
		search,
		true)
}

func (it Variant) EnumType() enuminf.EnumTyper {
	return BasicEnumImpl.EnumType()
}

func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it Variant) AsStringCompareTyper() enuminf.StringCompareTyper {
	return &it
}

func (it Variant) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}

func (it Variant) ToPtr() *Variant {
	return &it
}
