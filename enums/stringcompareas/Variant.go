package stringcompareas

import (
	"errors"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/errcore"
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
)

func (it *Variant) Name() string {
	return basicEnumImpl.ToEnumString(it.ValueByte())
}

func (it *Variant) TypeName() string {
	return basicEnumImpl.TypeName()
}

func (it *Variant) ToNumberString() string {
	return basicEnumImpl.ToNumberString(it.ValueByte())
}

func (it *Variant) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return basicEnumImpl.UnmarshallToValue(
		isMappedToDefault,
		jsonUnmarshallingValue)
}

func (it *Variant) String() string {
	return basicEnumImpl.ToEnumString(it.ValueByte())
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
	return basicEnumImpl.ToEnumJsonBytes(it.ValueByte()), nil
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	rawScriptType, err := basicEnumImpl.UnmarshallToValue(
		isMappedToDefault, data)

	if err == nil {
		*it = Variant(rawScriptType)
	}

	return err
}

func (it *Variant) RangeNamesCsv() string {
	return basicEnumImpl.RangeNamesCsv()
}

func (it *Variant) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return it
}

func (it *Variant) MaxByte() byte {
	return basicEnumImpl.Max()
}

func (it *Variant) MinByte() byte {
	return basicEnumImpl.Min()
}

func (it *Variant) ValueByte() byte {
	return byte(*it)
}

func (it *Variant) RangesByte() []byte {
	return basicEnumImpl.Ranges()
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

func (it *Variant) AsBasicByteEnumContractsBinder() coreinterface.BasicByteEnumContractsBinder {
	return it
}
