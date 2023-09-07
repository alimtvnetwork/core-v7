package enumimpl

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreimpl/enumimpl/enumtype"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/defaulterr"
	"gitlab.com/auk-go/core/errcore"
)

type BasicUInt16 struct {
	numberEnumBase
	jsonDoubleQuoteNameToValueHashMap        map[string]uint16 // contains names double quotes to value
	valueToJsonDoubleQuoteStringBytesHashmap map[uint16][]byte // contains value to string bytes with double quotes
	valueNameHashmap                         map[uint16]string // contains name without double quotes
	minVal, maxVal                           uint16
}

func (it BasicUInt16) IsAnyNamesOf(
	value uint16,
	names ...string,
) bool {
	currentName := it.ToEnumString(value)

	for _, name := range names {
		if name == currentName {
			return true
		}
	}

	return false
}

func (it BasicUInt16) IsAnyOf(value uint16, checkingItems ...uint16) bool {
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

func (it BasicUInt16) Max() uint16 {
	return it.maxVal
}

func (it BasicUInt16) Min() uint16 {
	return it.minVal
}

func (it BasicUInt16) GetValueByString(valueString string) uint16 {
	return it.jsonDoubleQuoteNameToValueHashMap[valueString]
}

func (it BasicUInt16) GetValueByName(
	name string,
) (uint16, error) {
	v, has := it.jsonDoubleQuoteNameToValueHashMap[name]

	if has {
		return v, nil
	}

	wrapped := fmt.Sprintf(
		constants.SprintDoubleQuoteFormat,
		name)

	nextVal, isFoundByWrapped := it.jsonDoubleQuoteNameToValueHashMap[wrapped]

	if isFoundByWrapped {
		return nextVal, nil
	}

	// has error
	return constants.Zero, enumUnmarshallingMappingFailedError(
		it.TypeName(),
		name,
		it.RangeNamesCsv())
}

func (it BasicUInt16) GetStringValue(input uint16) string {
	return it.StringRanges()[input]
}

func (it BasicUInt16) ExpectingEnumValueError(
	rawString string,
	expectedEnum interface{},
) error {
	expectedEnumName := it.ToName(expectedEnum)
	expectedValue := it.GetValueByString(expectedEnumName)
	convValue, err := it.GetValueByName(rawString)

	if err != nil {
		return errcore.ExpectingErrorSimpleNoType(
			"Expecting enum: "+expectedEnumName,
			expectedEnumName,
			rawString+err.Error())
	}

	if convValue == expectedValue {
		return nil
	}

	return errcore.ExpectingErrorSimpleNoType(
		"Expecting enum: "+expectedEnumName,
		expectedEnumName,
		rawString+it.RangesInvalidMessage())
}

func (it BasicUInt16) Ranges() []uint16 {
	return it.actualValueRanges.([]uint16)
}

func (it BasicUInt16) Hashmap() map[string]uint16 {
	return it.jsonDoubleQuoteNameToValueHashMap
}

func (it BasicUInt16) HashmapPtr() *map[string]uint16 {
	return &it.jsonDoubleQuoteNameToValueHashMap
}

func (it BasicUInt16) IsValidRange(value uint16) bool {
	return value >= it.minVal && value <= it.maxVal
}

// ToEnumJsonBytes used for MarshalJSON from map
func (it BasicUInt16) ToEnumJsonBytes(value uint16) ([]byte, error) {
	jsonBytes, has := it.valueToJsonDoubleQuoteStringBytesHashmap[value]

	if has {
		return jsonBytes, nil
	}

	return []byte{}, it.notFoundJsonBytesError(value)
}

func (it BasicUInt16) ToEnumString(value uint16) string {
	return it.valueNameHashmap[value]
}

func (it BasicUInt16) AppendPrependJoinValue(
	joiner string,
	appendVal, prependVal uint16,
) string {
	return it.ToEnumString(prependVal) +
		joiner +
		it.ToEnumString(appendVal)
}

func (it BasicUInt16) AppendPrependJoinNamer(
	joiner string,
	appendVal, prependVal coreinterface.ToNamer,
) string {
	return prependVal.Name() +
		joiner +
		appendVal.Name()
}

func (it BasicUInt16) ToNumberString(valueInRawFormat interface{}) string {
	return fmt.Sprintf(constants.SprintValueFormat, valueInRawFormat)
}

// UnmarshallToValue Mostly used for UnmarshalJSON
//
// Given bytes string enum value and transpile to exact enum raw value using map
func (it BasicUInt16) UnmarshallToValue(
	isMappedToFirstIfEmpty bool,
	jsonUnmarshallingValue []byte,
) (uint16, error) {
	if !isMappedToFirstIfEmpty && jsonUnmarshallingValue == nil {
		return constants.Zero,
			defaulterr.UnmarshallingFailedDueToNilOrEmpty
	}

	if isMappedToFirstIfEmpty && jsonUnmarshallingValue == nil {
		return it.minVal, nil
	}

	str := string(jsonUnmarshallingValue)
	if isMappedToFirstIfEmpty &&
		(str == constants.EmptyString || str == constants.DoubleQuotationStartEnd) {
		return it.minVal, nil
	}

	return it.GetValueByName(str)
}

func (it BasicUInt16) EnumType() enumtype.Variant {
	return enumtype.UnsignedInteger16
}
