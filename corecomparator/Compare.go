package corecomparator

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type Compare byte

const (
	Equal Compare = iota
	LeftGreater
	LeftGreaterEqual
	LeftLess
	LeftLessEqual
	NotEqual
	Inconclusive
)

func (it Compare) Is(other Compare) bool {
	return it == other
}

func (it Compare) IsLess() bool {
	return it == LeftLess
}

func (it Compare) IsLessEqual() bool {
	return it == LeftLess || it == Equal
}

func (it Compare) IsGreater() bool {
	return it == LeftGreater
}

func (it Compare) IsGreaterEqual() bool {
	return it == LeftGreater || it == Equal
}

func (it Compare) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Compare) ToNumberString() string {
	return strconv.Itoa(int(it))
}

func (it Compare) IsDefined() bool {
	return it != Inconclusive
}

func (it Compare) IsValid() bool {
	return it != Inconclusive
}

func (it Compare) IsInvalid() bool {
	return it == Inconclusive
}

func (it Compare) RangeNamesCsv() string {
	return RangeNamesCsv()
}

func (it Compare) IsValueEqual(value byte) bool {
	return byte(it) == value
}

func (it Compare) IsEqual() bool {
	return it == Equal
}

func (it Compare) IsLeftGreater() bool {
	return it == LeftGreater
}

func (it Compare) IsLeftGreaterEqual() bool {
	return it == LeftGreaterEqual
}

func (it Compare) IsLeftLess() bool {
	return it == LeftLess
}

func (it Compare) IsLeftLessEqual() bool {
	return it == LeftLessEqual
}

// IsLeftLessOrLessEqualOrEqual
//
//   it == Equal || it == LeftLess || it == LeftLessEqual
func (it Compare) IsLeftLessOrLessEqualOrEqual() bool {
	return it == Equal || it == LeftLess || it == LeftLessEqual
}

// IsLeftLessEqualLogically
//
//  it == Equal || it == LeftLess || it == LeftLessEqual
func (it Compare) IsLeftLessEqualLogically() bool {
	return it == Equal || it == LeftLess || it == LeftLessEqual
}

// IsLeftGreaterOrGreaterEqualOrEqual
//
//  it == Equal || it == LeftGreater || it == LeftGreaterEqual
func (it Compare) IsLeftGreaterOrGreaterEqualOrEqual() bool {
	return it == Equal || it == LeftGreater || it == LeftGreaterEqual
}

// IsLeftGreaterEqualLogically
//
//   it == Equal || it == LeftGreater || it == LeftGreaterEqual
func (it Compare) IsLeftGreaterEqualLogically() bool {
	return it == Equal || it == LeftGreater || it == LeftGreaterEqual
}

func (it Compare) IsNotEqual() bool {
	return it == NotEqual
}

// IsNotEqualLogically
//
//  return it != Equal
func (it Compare) IsNotEqualLogically() bool {
	return it != Equal
}

// IsDefinedPlus
//
//  return  it != Inconclusive && it == right
func (it Compare) IsDefinedPlus(right Compare) bool {
	return it != Inconclusive && it == right
}

func (it Compare) IsInconclusive() bool {
	return it == Inconclusive
}

func (it Compare) IsNotInconclusive() bool {
	return it != Inconclusive
}

func (it Compare) IsDefinedProperly() bool {
	return it != Inconclusive
}

func (it Compare) IsInconclusiveOrNotEqual() bool {
	return it == Inconclusive || it == NotEqual
}

func (it Compare) IsAnyOf(values ...Compare) bool {
	if len(values) == 0 {
		return true
	}

	for _, value := range values {
		if it == value {
			return true
		}
	}

	return false
}

func (it Compare) NameValue() string {
	return fmt.Sprintf(
		constants.StringWithBracketWrapNumberFormat,
		it.Name(),
		it.Value())
}

func (it Compare) CsvStrings(values ...Compare) []string {
	if len(values) == 0 {
		return []string{}
	}

	slice := make([]string, len(values))

	for i, value := range values {
		slice[i] = value.NameValue()
	}

	return slice
}

func (it Compare) CsvString(values ...Compare) string {
	if len(values) == 0 {
		return ""
	}

	slice := it.CsvStrings(values...)

	return strings.Join(slice, constants.CsvJoiner)
}

func (it Compare) Name() string {
	return it.String()
}

func (it Compare) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.String())
}

func (it *Compare) UnmarshalJSON(data []byte) error {
	if data == nil {
		return errors.New("compare unmarshal json error: data nil")
	}

	name := string(data)

	compare, has := RangesMap[name]

	if has {
		*it = compare

		return nil
	}

	return errors.New(string(data) + " failed to convert to core-compare. Must be any of the values.")
}

func (it Compare) Value() byte {
	return byte(it)
}

func (it Compare) IsCompareEqualLogically(
	expectedCompare Compare,
) bool {
	if it == expectedCompare {
		return true
	}

	if expectedCompare.IsLeftGreaterEqualLogically() {
		return it.IsLeftGreaterEqualLogically()
	}

	if expectedCompare.IsLeftLessEqualLogically() {
		return it.IsLeftLessEqualLogically()
	}

	return false
}

func (it Compare) OnlySupportedErr(
	message string,
	onlySupportedCompares ...Compare,
) error {
	if message == "" {
		return it.OnlySupportedDirectErr(onlySupportedCompares...)
	}

	if it.IsAnyOf(onlySupportedCompares...) {
		return nil
	}

	csv := it.CsvString(onlySupportedCompares...)

	return fmt.Errorf(constants.EnumOnlySupportedWithMessageFormat,
		it,
		it.NameValue(),
		message,
		csv)
}

func (it Compare) OnlySupportedDirectErr(
	onlySupportedCompares ...Compare,
) error {
	if it.IsAnyOf(onlySupportedCompares...) {
		return nil
	}

	csv := it.CsvString(onlySupportedCompares...)

	return fmt.Errorf(constants.EnumOnlySupportedFormat,
		it,
		it.NameValue(),
		csv)
}

func (it Compare) OperatorSymbol() string {
	return CompareOperatorsSymbols[it]
}

func (it Compare) OperatorShortForm() string {
	return CompareOperatorsShotNames[it]
}

func (it Compare) SqlOperatorSymbol() string {
	return SqlCompareOperators[it]
}

func (it Compare) NumberString() string {
	return strconv.Itoa(int(it))
}

func (it Compare) NumberJsonString() string {
	return "\"" + strconv.Itoa(int(it)) + "\""
}

func (it Compare) StringValue() string {
	return string(it)
}

func (it Compare) String() string {
	return CompareNames[it]
}

func (it Compare) IsAnyNamesOf(names ...string) bool {
	for _, name := range names {
		if it.IsNameEqual(name) {
			return true
		}
	}

	return false
}

func (it Compare) ValueByte() byte {
	return byte(it)
}

func (it Compare) ValueInt() int {
	return int(it)
}

func (it Compare) ValueInt8() int8 {
	return int8(it)
}

func (it Compare) ValueInt16() int16 {
	return int16(it)
}

func (it Compare) ValueInt32() int32 {
	return int32(it)
}

func (it Compare) ValueString() string {
	return it.ToNumberString()
}

func (it Compare) Format(format string) (compiled string) {
	panic("Not implemented for compare purposefully : " + format)
}
