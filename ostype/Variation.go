package ostype

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface/enuminf"
)

type Variation byte

// https://stackoverflow.com/a/50117892 | https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63
// go tool dist list
const (
	Any Variation = iota
	Unknown
	Windows
	Linux
	DarwinOrMacOs
	JavaScript
	FreeBsd
	NetBsd
	OpenBsd
	DragonFly
	Android
	Plan9
	Solaris
	Nacl
	Illumos
	IOs
	Aix
)

func (it Variation) AllNameValues() []string {
	return basicEnumImplOsType.AllNameValues()
}

func (it Variation) OnlySupportedErr(names ...string) error {
	return basicEnumImplOsType.OnlySupportedErr(names...)
}

func (it Variation) OnlySupportedMsgErr(message string, names ...string) error {
	return basicEnumImplOsType.OnlySupportedMsgErr(message, names...)
}

func (it Variation) ValueUInt16() uint16 {
	return uint16(it)
}

func (it Variation) IntegerEnumRanges() []int {
	return basicEnumImplOsType.IntegerEnumRanges()
}

func (it Variation) MinMaxAny() (min, max interface{}) {
	return basicEnumImplOsType.MinMaxAny()
}

func (it Variation) MinValueString() string {
	return basicEnumImplOsType.MinValueString()
}

func (it Variation) MaxValueString() string {
	return basicEnumImplOsType.MaxValueString()
}

func (it Variation) MaxInt() int {
	return basicEnumImplOsType.MaxInt()
}

func (it Variation) MinInt() int {
	return basicEnumImplOsType.MinInt()
}

func (it Variation) RangesDynamicMap() map[string]interface{} {
	return basicEnumImplOsType.RangesDynamicMap()
}

func (it Variation) IsByteValueEqual(value byte) bool {
	return byte(it) == value
}

func (it Variation) Format(format string) (compiled string) {
	return basicEnumImplOsType.Format(format, it)
}

func (it Variation) IsEnumEqual(enum enuminf.BasicEnumer) bool {
	return it.Value() == enum.ValueByte()
}

func (it *Variation) IsAnyEnumsEqual(enums ...enuminf.BasicEnumer) bool {
	for _, enum := range enums {
		if it.IsEnumEqual(enum) {
			return true
		}
	}

	return false
}

func (it Variation) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Variation) IsAnyNamesOf(names ...string) bool {
	for _, name := range names {
		if it.IsNameEqual(name) {
			return true
		}
	}

	return false
}

func (it Variation) IsValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Variation) IsAnyValuesEqual(anyByteValues ...byte) bool {
	for _, currentVal := range anyByteValues {
		if it.IsValueEqual(currentVal) {
			return true
		}
	}

	return false
}

func (it Variation) ValueInt() int {
	return int(it)
}

func (it Variation) ValueInt8() int8 {
	return int8(it)
}

func (it Variation) ValueInt16() int16 {
	return int16(it)
}

func (it Variation) ValueInt32() int32 {
	return int32(it)
}

func (it Variation) ValueString() string {
	return it.ToNumberString()
}

func (it Variation) IsValid() bool {
	return it.Value() != 0
}

func (it Variation) IsInvalid() bool {
	return it.Value() == 0
}

func (it Variation) IsByte(another byte) bool {
	return it == Variation(another)
}

func (it Variation) IsAnyOperatingSystem() bool {
	return Any == it
}

func (it Variation) Is(other Variation) bool {
	return other == it
}

func (it Variation) IsAnyMatch(others ...Variation) bool {
	for _, other := range others {
		if other == it {
			return true
		}
	}

	return false
}

func (it Variation) IsStringsMatchAny(others ...string) bool {
	for _, other := range others {
		otherVariant := GetVariant(other)

		if otherVariant == it {
			return true
		}
	}

	return false
}

// IsPossibleUnixGroup variation != Windows
func (it Variation) IsPossibleUnixGroup() bool {
	return it != Windows
}

func (it Variation) IsLinuxOrMac() bool {
	return it == Linux || it == DarwinOrMacOs
}

func (it Variation) Group() Group {
	if it == Windows {
		return WindowsGroup
	}

	if it == Android {
		return AndroidGroup
	}

	return UnixGroup
}

func (it Variation) IsActualGroupUnix() bool {
	return it.Group().IsUnix()
}

func (it Variation) IsWindows() bool {
	return it == Windows
}

func (it Variation) IsLinux() bool {
	return it == Linux
}

func (it Variation) IsDarwinOrMacOs() bool {
	return it == DarwinOrMacOs
}

func (it Variation) IsJavaScript() bool {
	return it == JavaScript
}

func (it Variation) IsFreeBsd() bool {
	return it == FreeBsd
}

func (it Variation) IsNetBsd() bool {
	return it == NetBsd
}

func (it Variation) IsOpenBsd() bool {
	return it == NetBsd
}

func (it Variation) IsDragonFly() bool {
	return it == DragonFly
}

func (it Variation) MarshalJSON() ([]byte, error) {
	return basicEnumImplOsType.ToEnumJsonBytes(it.Value())
}

func (it *Variation) UnmarshalJSON(data []byte) error {
	valueByte, err := basicEnumImplOsType.UnmarshallToValue(
		true,
		data)

	if err == nil {
		*it = Variation(valueByte)
	}

	return err
}

func (it Variation) Name() string {
	return basicEnumImplOsType.ToEnumString(it.Value())
}

func (it Variation) GoosName() string {
	return basicEnumImplOsType.ToEnumString(it.Value())
}

func (it Variation) NameValue() string {
	return basicEnumImplOsType.NameWithValue(it.Value())
}

func (it Variation) ToNumberString() string {
	return basicEnumImplOsType.ToNumberString(it.Value())
}

func (it Variation) RangeNamesCsv() string {
	return basicEnumImplOsType.RangeNamesCsv()
}

func (it Variation) TypeName() string {
	return basicEnumImplOsType.TypeName()
}

func (it Variation) UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error) {
	return basicEnumImplOsType.UnmarshallToValue(true, jsonUnmarshallingValue)
}

func (it Variation) MaxByte() byte {
	return basicEnumImplOsType.Max()
}

func (it Variation) MinByte() byte {
	return basicEnumImplOsType.Min()
}

func (it Variation) ValueByte() byte {
	return byte(it)
}

func (it Variation) RangesByte() []byte {
	return basicEnumImplOsType.Ranges()
}

func (it Variation) Value() byte {
	return byte(it)
}

func (it Variation) String() string {
	return basicEnumImplOsType.ToEnumString(it.Value())
}

func (it Variation) EnumType() enuminf.EnumTyper {
	return basicEnumImplOsType.EnumType()
}

func (it Variation) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it Variation) AsJsonContractsBinder() corejson.JsonMarshaller {
	return &it
}

func (it Variation) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}

func (it Variation) ToPtr() *Variation {
	return &it
}
