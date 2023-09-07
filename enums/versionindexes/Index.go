package versionindexes

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface/enuminf"
	"gitlab.com/auk-go/core/defaulterr"
)

type Index byte

const (
	Major   Index = iota
	Minor   Index = 1
	Patch   Index = 2
	Build   Index = 3
	Invalid Index = 4
)

func (it Index) AllNameValues() []string {
	return BasicEnumImpl.AllNameValues()
}

func (it Index) OnlySupportedErr(names ...string) error {
	return BasicEnumImpl.OnlySupportedErr(names...)
}

func (it Index) OnlySupportedMsgErr(message string, names ...string) error {
	return BasicEnumImpl.OnlySupportedMsgErr(message, names...)
}

func (it Index) ValueUInt16() uint16 {
	return uint16(it)
}

func (it Index) IntegerEnumRanges() []int {
	return BasicEnumImpl.IntegerEnumRanges()
}

func (it Index) MinMaxAny() (min, max interface{}) {
	return BasicEnumImpl.MinMaxAny()
}

func (it Index) MinValueString() string {
	return BasicEnumImpl.MinValueString()
}

func (it Index) MaxValueString() string {
	return BasicEnumImpl.MaxValueString()
}

func (it Index) MaxInt() int {
	return BasicEnumImpl.MaxInt()
}

func (it Index) MinInt() int {
	return BasicEnumImpl.MinInt()
}

func (it Index) RangesDynamicMap() map[string]interface{} {
	return BasicEnumImpl.RangesDynamicMap()
}

func (it Index) IsByteValueEqual(value byte) bool {
	return byte(it) == value
}

func (it Index) Format(format string) (compiled string) {
	return BasicEnumImpl.Format(format, it)
}

func (it Index) IsEnumEqual(enum enuminf.BasicEnumer) bool {
	return it.ValueByte() == enum.ValueByte()
}

func (it *Index) IsAnyEnumsEqual(enums ...enuminf.BasicEnumer) bool {
	for _, enum := range enums {
		if it.IsEnumEqual(enum) {
			return true
		}
	}

	return false
}

func (it Index) IsValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Index) IsAnyValuesEqual(anyByteValues ...byte) bool {
	for _, valByte := range anyByteValues {
		if it.IsValueEqual(valByte) {
			return true
		}
	}

	return false
}

func (it Index) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Index) IsAnyNamesOf(names ...string) bool {
	for _, name := range names {
		if it.IsNameEqual(name) {
			return true
		}
	}

	return false
}

func (it Index) ValueInt8() int8 {
	return int8(it)
}

func (it Index) ValueInt16() int16 {
	return int16(it)
}

func (it Index) ValueInt32() int32 {
	return int32(it)
}

func (it Index) ValueString() string {
	return it.ToNumberString()
}

func (it Index) IsValid() bool {
	return it != Invalid
}

func (it Index) IsInvalid() bool {
	return it == Invalid
}

func (it Index) Name() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Index) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it Index) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.ValueByte())
}

func (it Index) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(
		isMappedToDefault,
		jsonUnmarshallingValue)
}

func (it Index) String() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Index) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Index) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it Index) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.ValueByte())
}

func (it *Index) UnmarshalJSON(data []byte) error {
	rawScriptType, err := it.UnmarshallEnumToValue(
		data)

	if err == nil {
		*it = Index(rawScriptType)
	}

	return err
}

func (it Index) Json() corejson.Result {
	return corejson.New(it)
}

func (it Index) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Index) JsonParseSelfInject(jsonResult *corejson.Result) error {
	if jsonResult == nil {
		return defaulterr.UnmarshallingFailedDueToNilOrEmpty
	}

	if jsonResult.HasError() {
		return jsonResult.MeaningfulError()
	}

	v, err := it.UnmarshallEnumToValue(jsonResult.Bytes)

	if err == nil {
		*it = Index(v)
	}

	return err
}

func (it *Index) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it Index) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it *Index) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it *Index) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it Index) ValueByte() byte {
	return byte(it)
}

func (it Index) ValueInt() int {
	return int(it)
}

func (it *Index) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it Index) IsMajor() bool {
	return it == Major
}

func (it Index) IsMinor() bool {
	return it == Minor
}

func (it Index) IsPatch() bool {
	return it == Patch
}

func (it Index) IsBuild() bool {
	return it == Build
}

func (it Index) EnumType() enuminf.EnumTyper {
	return BasicEnumImpl.EnumType()
}

func (it Index) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}

func (it Index) ToPtr() *Index {
	return &it
}
