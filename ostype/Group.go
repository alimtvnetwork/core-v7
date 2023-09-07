package ostype

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface/enuminf"
)

type Group byte

const (
	WindowsGroup Group = iota
	UnixGroup
	AndroidGroup
	InvalidGroup
)

func (it Group) AllNameValues() []string {
	return basicEnumImplOsGroup.AllNameValues()
}

func (it Group) OnlySupportedErr(names ...string) error {
	return basicEnumImplOsGroup.OnlySupportedErr(names...)
}

func (it Group) OnlySupportedMsgErr(message string, names ...string) error {
	return basicEnumImplOsGroup.OnlySupportedMsgErr(message, names...)
}

func (it Group) ValueUInt16() uint16 {
	return uint16(it)
}

func (it Group) IntegerEnumRanges() []int {
	return basicEnumImplOsGroup.IntegerEnumRanges()
}

func (it Group) MinMaxAny() (min, max interface{}) {
	return basicEnumImplOsGroup.MinMaxAny()
}

func (it Group) MinValueString() string {
	return basicEnumImplOsGroup.MinValueString()
}

func (it Group) MaxValueString() string {
	return basicEnumImplOsGroup.MaxValueString()
}

func (it Group) MaxInt() int {
	return basicEnumImplOsGroup.MaxInt()
}

func (it Group) MinInt() int {
	return basicEnumImplOsGroup.MinInt()
}

func (it Group) RangesDynamicMap() map[string]interface{} {
	return basicEnumImplOsGroup.RangesDynamicMap()
}

func (it Group) IsByteValueEqual(value byte) bool {
	return byte(it) == value
}

func (it Group) Format(format string) (compiled string) {
	return basicEnumImplOsGroup.Format(format, it)
}

func (it Group) IsEnumEqual(enum enuminf.BasicEnumer) bool {
	return it.Value() == enum.ValueByte()
}

func (it *Group) IsAnyEnumsEqual(enums ...enuminf.BasicEnumer) bool {
	for _, enum := range enums {
		if it.IsEnumEqual(enum) {
			return true
		}
	}

	return false
}

func (it Group) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Group) IsAnyNamesOf(names ...string) bool {
	for _, name := range names {
		if it.IsNameEqual(name) {
			return true
		}
	}

	return false
}

func (it Group) IsValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Group) IsAnyValuesEqual(anyByteValues ...byte) bool {
	for _, currentVal := range anyByteValues {
		if it.IsValueEqual(currentVal) {
			return true
		}
	}

	return false
}

func (it Group) ValueInt() int {
	return int(it)
}

func (it Group) ValueInt8() int8 {
	return int8(it)
}

func (it Group) ValueInt16() int16 {
	return int16(it)
}

func (it Group) ValueInt32() int32 {
	return int32(it)
}

func (it Group) ValueString() string {
	return it.ToNumberString()
}

func (it Group) Is(another Group) bool {
	return it == another
}

func (it Group) IsWindows() bool {
	return it == WindowsGroup
}

func (it Group) IsUnix() bool {
	return it == UnixGroup
}

func (it Group) IsAndroid() bool {
	return it == AndroidGroup
}

func (it Group) IsInvalidGroup() bool {
	return it == InvalidGroup
}

func (it Group) Byte() byte {
	return byte(it)
}

func (it Group) MarshalJSON() ([]byte, error) {
	return basicEnumImplOsGroup.ToEnumJsonBytes(it.Value())
}

func (it *Group) UnmarshalJSON(data []byte) error {
	valueByte, err := basicEnumImplOsGroup.UnmarshallToValue(
		true,
		data)

	if err == nil {
		*it = Group(valueByte)
	}

	return err
}

func (it Group) Name() string {
	return basicEnumImplOsGroup.ToEnumString(it.Value())
}

func (it Group) NameValue() string {
	return basicEnumImplOsGroup.NameWithValue(it.Value())
}

func (it Group) ToNumberString() string {
	return basicEnumImplOsGroup.ToNumberString(it.Value())
}

func (it Group) RangeNamesCsv() string {
	return basicEnumImplOsGroup.RangeNamesCsv()
}

func (it Group) TypeName() string {
	return basicEnumImplOsGroup.TypeName()
}

func (it Group) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
	return basicEnumImplOsGroup.UnmarshallToValue(true, jsonUnmarshallingValue)
}

func (it Group) MaxByte() byte {
	return basicEnumImplOsGroup.Max()
}

func (it Group) MinByte() byte {
	return basicEnumImplOsGroup.Min()
}

func (it Group) ValueByte() byte {
	return byte(it)
}

func (it Group) RangesByte() []byte {
	return basicEnumImplOsGroup.Ranges()
}

func (it Group) Value() byte {
	return byte(it)
}

func (it Group) String() string {
	return basicEnumImplOsGroup.ToEnumString(it.Value())
}

func (it Group) IsValid() bool {
	return it != InvalidGroup
}

func (it Group) IsInvalid() bool {
	return it == InvalidGroup
}

func (it Group) EnumType() enuminf.EnumTyper {
	return basicEnumImplOsGroup.EnumType()
}

func (it *Group) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return it
}

func (it *Group) AsJsonContractsBinder() corejson.JsonMarshaller {
	return it
}

func (it Group) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}

func (it Group) ToPtr() *Group {
	return &it
}
