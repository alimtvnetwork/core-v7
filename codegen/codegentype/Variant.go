package codegentype

import (
	"strconv"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface/enuminf"
)

type Variant byte

const (
	Simple Variant = iota
	MultipleArranges
)

func (it Variant) IsSimple() bool {
	return it == Simple
}

func (it Variant) IsMultipleArranges() bool {
	return it == MultipleArranges
}

func (it Variant) AllNameValues() []string {
	return basicEnumImpl.AllNameValues()
}

func (it Variant) OnlySupportedErr(
	names ...string,
) error {
	return basicEnumImpl.OnlySupportedErr(names...)
}

func (it Variant) OnlySupportedMsgErr(
	message string,
	names ...string,
) error {
	return basicEnumImpl.OnlySupportedMsgErr(
		message, names...,
	)
}

func (it Variant) ValueUInt16() uint16 {
	return uint16(it)
}

func (it Variant) IntegerEnumRanges() []int {
	return basicEnumImpl.IntegerEnumRanges()
}

func (it Variant) MinMaxAny() (min, max interface{}) {
	return basicEnumImpl.MinMaxAny()
}

func (it Variant) MinValueString() string {
	return basicEnumImpl.MinValueString()
}

func (it Variant) MaxValueString() string {
	return basicEnumImpl.MaxValueString()
}

func (it Variant) MaxInt() int {
	return basicEnumImpl.MaxInt()
}

func (it Variant) MinInt() int {
	return basicEnumImpl.MinInt()
}

func (it Variant) RangesDynamicMap() map[string]interface{} {
	return basicEnumImpl.RangesDynamicMap()
}

func (it Variant) IsValueEqual(value byte) bool {
	return byte(it) == value
}

func (it Variant) Format(format string) (compiled string) {
	return basicEnumImpl.Format(format, it)
}

func (it Variant) IsEnumEqual(enum enuminf.BasicEnumer) bool {
	return it.Value() == enum.ValueByte()
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

func (it Variant) ValueByte() byte {
	return byte(it)
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
	return it != 0
}

func (it Variant) IsInvalid() bool {
	return it == 0
}

func (it Variant) NameValue() string {
	return basicEnumImpl.NameWithValue(it)
}

func (it Variant) ToNumberString() string {
	return strconv.Itoa(it.ValueInt())
}

func (it Variant) Name() string {
	return basicEnumImpl.ToEnumString(it.Value())
}

func (it Variant) UnmarshallToValue(jsonUnmarshallingValue []byte) (byte, error) {
	newEmpty := Variant(0)
	err := corejson.
		Deserialize.
		UsingBytes(
			jsonUnmarshallingValue, &newEmpty,
		)

	if err != nil {
		return 0, err
	}

	return newEmpty.Value(), nil
}

func (it Variant) MarshalJSON() ([]byte, error) {
	return basicEnumImpl.ToEnumJsonBytes(it.Value())
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	newEmpty := Variant(0)
	err := corejson.
		Deserialize.
		UsingBytes(
			data, &newEmpty,
		)

	if err == nil {
		*it = newEmpty
	}

	return err
}

func (it Variant) String() string {
	return basicEnumImpl.ToEnumString(it.Value())
}

func (it Variant) JsonString() string {
	return basicEnumImpl.JsonString(it)
}

func (it Variant) StringRangesPtr() *[]string {
	return basicEnumImpl.StringRangesPtr()
}

func (it Variant) StringRanges() []string {
	return basicEnumImpl.StringRanges()
}

func (it Variant) RangesInvalidMessage() string {
	return basicEnumImpl.RangesInvalidMessage()
}

func (it Variant) RangesInvalidErr() error {
	return basicEnumImpl.RangesInvalidErr()
}

func (it Variant) IsValidRange() bool {
	return basicEnumImpl.IsValidRange(it.Value())
}

func (it Variant) IsInvalidRange() bool {
	return !it.IsValidRange()
}

func (it Variant) Value() byte {
	return byte(it)
}

func (it Variant) StringValue() string {
	return strconv.Itoa(it.ValueInt())
}

func (it Variant) HasIndexInStrings(sliceOfStrings ...string) (val string, isValid bool) {
	if len(sliceOfStrings) == 0 {
		return "", false
	}

	enumVal := it.ValueInt()
	isValid = len(sliceOfStrings)-1 >= enumVal

	if isValid {
		return sliceOfStrings[enumVal], isValid
	}

	return "", false
}

// Add v + n
func (it Variant) Add(n byte) Variant {
	return Variant(it.Value() + n)
}

// Subtract v - n
func (it Variant) Subtract(n byte) Variant {
	return Variant(it.Value() - n)
}

func (it Variant) Is(n Variant) bool {
	return it.Value() == n.Value()
}

// IsBetween val >= start &&  val <= end
func (it Variant) IsBetween(start, end byte) bool {
	val := it.Value()

	return val >= start && val <= end
}

// IsBetweenInt val >= start &&  val <= end
func (it Variant) IsBetweenInt(start, end int) bool {
	val := it.Value()

	return val >= byte(start) && val <= byte(end)
}

func (it Variant) IsEqual(n byte) bool {
	return it.Value() == n
}

// IsGreater v.Value() > n
func (it Variant) IsGreater(n byte) bool {
	return it.Value() > n
}

// IsGreaterEqual v.Value() >= n
func (it Variant) IsGreaterEqual(n byte) bool {
	return it.Value() >= n
}

// IsLess v.Value() < n
func (it Variant) IsLess(n byte) bool {
	return it.Value() < n
}

// IsLessEqual v.Value() <= n
func (it Variant) IsLessEqual(n byte) bool {
	return it.Value() <= n
}

func (it Variant) IsEqualInt(n int) bool {
	return it.Value() == byte(n)
}

// IsGreaterInt v.Value() > n
func (it Variant) IsGreaterInt(n int) bool {
	return it.Value() > byte(n)
}

// IsGreaterEqualInt v.Value() >= n
func (it Variant) IsGreaterEqualInt(n int) bool {
	return it.Value() >= byte(n)
}

// IsLessInt v.Value() < n
func (it Variant) IsLessInt(n int) bool {
	return it.Value() < byte(n)
}

// IsLessEqualInt v.Value() <= n
func (it Variant) IsLessEqualInt(n int) bool {
	return it.Value() <= byte(n)
}

func (it Variant) RangeNamesCsv() string {
	return basicEnumImpl.RangeNamesCsv()
}

func (it Variant) TypeName() string {
	return basicEnumImpl.TypeName()
}

func (it Variant) EnumType() enuminf.EnumTyper {
	return basicEnumImpl.EnumType()
}

func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it Variant) ToPtr() *Variant {
	return &it
}
