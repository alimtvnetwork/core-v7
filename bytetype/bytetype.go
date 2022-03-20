package bytetype

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
)

type Variant byte

func (it Variant) AllNameValues() []string {
	return BasicEnumImpl.AllNameValues()
}

func (it Variant) OnlySupportedErr(
	names ...string,
) error {
	return BasicEnumImpl.OnlySupportedErr(names...)
}

func (it Variant) OnlySupportedMsgErr(
	message string,
	names ...string,
) error {
	return BasicEnumImpl.OnlySupportedMsgErr(
		message, names...)
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

func (it Variant) IsValueEqual(value byte) bool {
	return byte(it) == value
}

func (it Variant) Format(format string) (compiled string) {
	return BasicEnumImpl.Format(format, it)
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
	return BasicEnumImpl.NameWithValue(it)
}

func (it Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.Value())
}

func (it Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it Variant) UnmarshallToValue(jsonUnmarshallingValue []byte) (byte, error) {
	newEmpty := Variant(0)
	err := corejson.
		Deserialize.
		UsingBytes(
			jsonUnmarshallingValue, &newEmpty)

	if err != nil {
		return 0, err
	}

	return newEmpty.Value(), nil
}

func (it Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.Value())
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	newEmpty := Variant(0)
	err := corejson.
		Deserialize.
		UsingBytes(
			data, &newEmpty)

	if err == nil {
		*it = newEmpty
	}

	return err
}

func (it Variant) String() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it Variant) JsonString() string {
	return BasicEnumImpl.JsonString(it)
}

func (it Variant) StringRangesPtr() *[]string {
	return BasicEnumImpl.StringRangesPtr()
}

func (it Variant) StringRanges() []string {
	return BasicEnumImpl.StringRanges()
}

func (it Variant) RangesInvalidMessage() string {
	return BasicEnumImpl.RangesInvalidMessage()
}

func (it Variant) RangesInvalidErr() error {
	return BasicEnumImpl.RangesInvalidErr()
}

func (it Variant) IsValidRange() bool {
	return BasicEnumImpl.IsValidRange(it.Value())
}

func (it Variant) IsInvalidRange() bool {
	return !it.IsValidRange()
}

func (it Variant) Value() byte {
	return byte(it)
}

func (it Variant) StringValue() string {
	return string(it)
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
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it Variant) EnumType() enuminf.EnumTyper {
	return BasicEnumImpl.EnumType()
}

func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it Variant) ToPtr() *Variant {
	return &it
}
