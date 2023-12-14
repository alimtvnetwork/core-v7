package enumimpl

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreimpl/enumimpl/enumtype"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/defaulterr"
	"gitlab.com/auk-go/core/errcore"
)

type BasicString struct {
	numberEnumBase
	nameWithIndexMap                         map[string]int
	jsonDoubleQuoteNameToValueHashMap        map[string]bool   // contains names double quotes to value
	valueToJsonDoubleQuoteStringBytesHashmap map[string][]byte // contains value to string bytes with double quotes
	minVal, maxVal                           string
}

func (it BasicString) IsAnyNamesOf(
	valueName string,
	names ...string,
) bool {
	for _, name := range names {
		if name == valueName {
			return true
		}
	}

	return false
}

func (it BasicString) IsAnyOf(value string, checkingItems ...string) bool {
	if len(checkingItems) == 0 {
		return true
	}

	for _, givenByte := range checkingItems {
		if value == givenByte {
			return true
		}
	}

	return false
}

func (it BasicString) Max() string {
	return it.maxVal
}

func (it BasicString) Min() string {
	return it.minVal
}

func (it BasicString) Ranges() []string {
	return it.actualValueRanges.([]string)
}

func (it BasicString) HasAnyItem() bool {
	return it.Length() > 0
}

func (it BasicString) MaxIndex() int {
	return it.Length() - 1
}

func (it BasicString) GetNameByIndex(index int) string {
	lastIndex := it.Length() - 1

	if lastIndex >= index && index > 0 {
		return it.StringRanges()[index]
	}

	return constants.EmptyString
}

// GetIndexByName
//
//	constants.InvalidValue refers to the invalid index
func (it BasicString) GetIndexByName(name string) int {
	if name == "" {
		return constants.InvalidValue
	}

	lastIndex := it.Length() - 1

	if lastIndex < 0 {
		return constants.InvalidValue
	}

	index, has := it.nameWithIndexMap[name]

	if has {
		return index
	}

	return constants.InvalidValue
}

func (it BasicString) NameWithIndexMap() map[string]int {
	return it.nameWithIndexMap
}

func (it BasicString) RangesIntegers() []int {
	length := it.Length()

	slice := make([]int, length)

	for i := 0; i < length; i++ {
		slice[i] = i
	}

	return slice
}

func (it BasicString) Hashset() map[string]bool {
	return it.jsonDoubleQuoteNameToValueHashMap
}

func (it BasicString) HashsetPtr() *map[string]bool {
	return &it.jsonDoubleQuoteNameToValueHashMap
}

func (it BasicString) GetValueByName(name string) (string, error) {
	_, has := it.jsonDoubleQuoteNameToValueHashMap[name]

	if has {
		return name, nil
	}

	wrapped := fmt.Sprintf(
		constants.SprintDoubleQuoteFormat,
		name)

	_, isFoundByWrapped := it.jsonDoubleQuoteNameToValueHashMap[wrapped]

	if isFoundByWrapped {
		return wrapped, nil
	}

	// has error
	return constants.EmptyString, enumUnmarshallingMappingFailedError(
		it.TypeName(),
		name,
		it.RangeNamesCsv())
}

func (it BasicString) IsValidRange(value string) bool {
	return it.jsonDoubleQuoteNameToValueHashMap[value]
}

func (it BasicString) OnlySupportedErr(
	supportedNames ...string,
) error {
	return OnlySupportedErr(
		it.StringRanges(),
		supportedNames...)
}

func (it BasicString) OnlySupportedMsgErr(
	errMessage string,
	supportedNames ...string,
) error {
	return errcore.ConcatMessageWithErr(
		errMessage,
		it.OnlySupportedErr(supportedNames...))
}

func (it BasicString) AppendPrependJoinValue(
	joiner string,
	appendVal, prependVal string,
) string {
	return it.ToEnumString(prependVal) +
		joiner +
		it.ToEnumString(appendVal)
}

func (it BasicString) AppendPrependJoinNamer(
	joiner string,
	appendVal, prependVal coreinterface.ToNamer,
) string {
	return prependVal.Name() +
		joiner +
		appendVal.Name()
}

// ToEnumJsonBytes used for MarshalJSON from map
func (it BasicString) ToEnumJsonBytes(value string) ([]byte, error) {
	jsonBytes, has := it.valueToJsonDoubleQuoteStringBytesHashmap[value]

	if has {
		return jsonBytes, nil
	}

	return []byte{}, it.notFoundJsonBytesError(value)
}

// UnmarshallToValue Mostly used for UnmarshalJSON
//
// Given bytes string enum value and transpile to exact enum raw value using map
func (it BasicString) UnmarshallToValue(
	isMappedToFirstIfEmpty bool,
	jsonUnmarshallingValue []byte,
) (string, error) {
	if !isMappedToFirstIfEmpty && jsonUnmarshallingValue == nil {
		return constants.EmptyString,
			defaulterr.UnmarshallingFailedDueToNilOrEmpty
	}

	if isMappedToFirstIfEmpty && jsonUnmarshallingValue == nil {
		return it.minVal, nil
	}

	str := string(jsonUnmarshallingValue)
	if isMappedToFirstIfEmpty && (str == "" || str == `""`) {
		return it.minVal, nil
	}

	return it.GetValueByName(str)
}

func (it BasicString) EnumType() enumtype.Variant {
	return enumtype.String
}
