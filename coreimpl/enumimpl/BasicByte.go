package enumimpl

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreimpl/enumimpl/enumtype"
	"gitlab.com/auk-go/core/defaulterr"
	"gitlab.com/auk-go/core/errcore"
)

type BasicByte struct {
	numberEnumBase
	jsonDoubleQuoteNameToValueHashMap        map[string]byte // contains names double quotes to value
	valueToJsonDoubleQuoteStringBytesHashmap map[byte][]byte // contains value to string bytes with double quotes
	valueNameHashmap                         map[byte]string // contains name without double quotes
	minVal, maxVal                           byte
}

func (it BasicByte) IsAnyOf(
	value byte,
	givenBytes ...byte,
) bool {
	if len(givenBytes) == 0 {
		return true
	}

	for _, givenByte := range givenBytes {
		if value == givenByte {
			return true
		}
	}

	return false
}

func (it BasicByte) IsAnyNamesOf(
	value byte,
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

func (it BasicByte) Max() byte {
	return it.maxVal
}

func (it BasicByte) Min() byte {
	return it.minVal
}

func (it BasicByte) GetValueByString(
	jsonValueString string,
) byte {
	return it.jsonDoubleQuoteNameToValueHashMap[jsonValueString]
}

func (it BasicByte) GetValueByName(
	name string,
) (byte, error) {
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
	return 0, enumUnmarshallingMappingFailedError(
		it.TypeName(),
		name,
		it.RangeNamesCsv())
}

func (it BasicByte) GetStringValue(
	input byte,
) string {
	return it.StringRanges()[input]
}

func (it BasicByte) ExpectingEnumValueError(
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

func (it BasicByte) Ranges() []byte {
	return it.actualValueRanges.([]byte)
}

func (it BasicByte) Hashmap() map[string]byte {
	return it.jsonDoubleQuoteNameToValueHashMap
}

func (it BasicByte) HashmapPtr() *map[string]byte {
	return &it.jsonDoubleQuoteNameToValueHashMap
}

func (it BasicByte) IsValidRange(
	value byte,
) bool {
	return value >= it.minVal && value <= it.maxVal
}

// ToEnumJsonBytes used for MarshalJSON from map
func (it BasicByte) ToEnumJsonBytes(
	value byte,
) ([]byte, error) {
	jsonBytes, has := it.valueToJsonDoubleQuoteStringBytesHashmap[value]

	if has {
		return jsonBytes, nil
	}

	return []byte{}, it.notFoundJsonBytesError(value)
}

func (it BasicByte) ToEnumString(
	value byte,
) string {
	return it.valueNameHashmap[value]
}

func (it BasicByte) AppendPrependJoinValue(
	joiner string,
	appendVal, prependVal byte,
) string {
	return it.ToEnumString(prependVal) +
		joiner +
		it.ToEnumString(appendVal)
}

func (it BasicByte) AppendPrependJoinNamer(
	joiner string,
	appendVal, prependVal toNamer,
) string {
	return prependVal.Name() +
		joiner +
		appendVal.Name()
}

func (it BasicByte) ToNumberString(
	valueInNumberFormat interface{}, // 1, 2, ... any number (byte / int, ...)
) string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		valueInNumberFormat)
}

func (it BasicByte) JsonMap() map[string]byte {
	return it.jsonDoubleQuoteNameToValueHashMap
}

// UnmarshallToValue Mostly used for UnmarshalJSON
//
// Given bytes string enum value and transpile to exact enum raw value using map
func (it BasicByte) UnmarshallToValue(
	isMappedToFirstIfEmpty bool,
	jsonUnmarshallingValue []byte,
) (byte, error) {
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

func (it BasicByte) EnumType() enumtype.Variant {
	return enumtype.Byte
}

func (it BasicByte) AsBasicByter() BasicByter {
	return &it
}
