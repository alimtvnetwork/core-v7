package chmodclasstype

import (
	"gitlab.com/auk-go/core/coreinterface/enuminf"
)

type Variant byte

const (
	Invalid Variant = iota
	All
	Owner
	Group
	Other
	OwnerGroup
	GroupOther
	OwnerOther
)

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

func (it Variant) IsUnInitialized() bool {
	return it == Invalid
}

func (it Variant) IsAll() bool {
	return it == All
}

func (it Variant) IsOwner() bool {
	return it == Owner
}

func (it Variant) IsGroup() bool {
	return it == Group
}

func (it Variant) IsOther() bool {
	return it == Other
}

func (it Variant) IsOwnerGroup() bool {
	return it == OwnerGroup
}

func (it Variant) IsGroupOther() bool {
	return it == GroupOther
}

func (it Variant) IsOwnerOther() bool {
	return it == OwnerOther
}

func (it Variant) IsValid() bool {
	return it != Invalid
}

func (it Variant) IsInvalid() bool {
	return it == Invalid
}

func (it *Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it *Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.ValueByte())
}

func (it Variant) String() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it *Variant) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(true, jsonUnmarshallingValue)
}

func (it *Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.Value())
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	dataConv, err := it.UnmarshallEnumToValue(data)

	if err == nil {
		*it = Variant(dataConv)
	}

	return err
}

func (it Variant) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it Variant) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it *Variant) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it *Variant) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it Variant) ValueByte() byte {
	return byte(it)
}

func (it Variant) Value() byte {
	return byte(it)
}

func (it *Variant) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it Variant) EnumType() enuminf.EnumTyper {
	return BasicEnumImpl.EnumType()
}

func (it Variant) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}
