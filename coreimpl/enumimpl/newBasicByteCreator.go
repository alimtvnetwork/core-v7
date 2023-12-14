package enumimpl

import (
	"reflect"
	"strconv"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type newBasicByteCreator struct{}

func (it newBasicByteCreator) CreateUsingMap(
	typeName string,
	actualRangesMap map[byte]string,
) *BasicByte {
	return it.CreateUsingMapPlusAliasMap(
		typeName,
		actualRangesMap,
		nil,
	)
}

func (it newBasicByteCreator) CreateUsingMapPlusAliasMap(
	typeName string,
	actualRangesMap map[byte]string,
	aliasingMap map[string]byte,
) *BasicByte {
	var min, max byte

	max = constants.MinUint
	min = 0

	actualValues := make([]byte, len(actualRangesMap))
	actualNames := make([]string, len(actualRangesMap))

	index := 0
	for val, name := range actualRangesMap {
		actualValues[index] = val
		actualNames[index] = name

		if max < val {
			max = val
		}

		index++
	}

	return it.CreateUsingAliasMap(
		typeName,
		actualValues,
		actualNames,
		aliasingMap, // aliasing map
		min,         // zero
		max,
	)
}

func (it newBasicByteCreator) CreateUsingSlicePlusAliasMapOptions(
	isIncludeUppercaseLowercase bool, // lowercase, uppercase all
	firstItem interface{},
	names []string,
	aliasingMap map[string]byte,
) *BasicByte {
	actualRangesMap := it.sliceNamesToMap(names)

	finalAliasMap := it.generateUppercaseLowercaseAliasMap(
		isIncludeUppercaseLowercase,
		actualRangesMap,
		aliasingMap)

	return it.CreateUsingMapPlusAliasMap(
		reflect.TypeOf(firstItem).String(),
		actualRangesMap,
		finalAliasMap,
	)
}

func (it newBasicByteCreator) CreateUsingMapPlusAliasMapOptions(
	isIncludeUppercaseLowercase bool, // lowercase, uppercase all
	firstItem interface{},
	actualRangesMap map[byte]string,
	aliasingMap map[string]byte,
) *BasicByte {
	finalAliasMap := it.generateUppercaseLowercaseAliasMap(
		isIncludeUppercaseLowercase,
		actualRangesMap,
		aliasingMap)

	return it.CreateUsingMapPlusAliasMap(
		reflect.TypeOf(firstItem).String(),
		actualRangesMap,
		finalAliasMap,
	)
}

func (it newBasicByteCreator) Create(
	typeName string,
	actualValueRanges []byte,
	stringRanges []string,
	min, max byte,
) *BasicByte {
	return it.CreateUsingAliasMap(
		typeName,
		actualValueRanges,
		stringRanges,
		nil,
		min,
		max)
}

// CreateUsingAliasMap
//
// Length : must match stringRanges and actualRangesAnyType
func (it newBasicByteCreator) CreateUsingAliasMap(
	typeName string,
	actualValueRanges []byte,
	stringRanges []string,
	aliasingMap map[string]byte,
	min, max byte,
) *BasicByte {
	enumBase := newNumberEnumBase(
		typeName,
		actualValueRanges,
		stringRanges,
		min,
		max)

	jsonDoubleQuoteNameToValueHashMap := make(map[string]byte, len(actualValueRanges))
	valueToJsonDoubleQuoteStringBytesHashmap := make(map[byte][]byte, len(actualValueRanges))
	valueNameHashmap := make(map[byte]string, len(actualValueRanges))

	for i, actualVal := range actualValueRanges {
		key := stringRanges[i]
		indexJson := toJsonName(i)
		indexString := strconv.Itoa(i)
		jsonName := toJsonName(key)
		jsonDoubleQuoteNameToValueHashMap[jsonName] = actualVal
		jsonDoubleQuoteNameToValueHashMap[key] = actualVal
		jsonDoubleQuoteNameToValueHashMap[indexJson] = actualVal
		jsonDoubleQuoteNameToValueHashMap[indexString] = actualVal
		valueToJsonDoubleQuoteStringBytesHashmap[actualVal] = []byte(jsonName)
		valueNameHashmap[actualVal] = key
	}

	if len(aliasingMap) > 0 {
		for aliasName, aliasValue := range aliasingMap {
			aliasJsonName := toJsonName(aliasName)
			jsonDoubleQuoteNameToValueHashMap[aliasName] = aliasValue
			jsonDoubleQuoteNameToValueHashMap[aliasJsonName] = aliasValue
		}
	}

	return &BasicByte{
		numberEnumBase:                           enumBase,
		minVal:                                   min,
		maxVal:                                   max,
		jsonDoubleQuoteNameToValueHashMap:        jsonDoubleQuoteNameToValueHashMap,
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
		valueNameHashmap:                         valueNameHashmap,
	}
}

func (it newBasicByteCreator) UsingFirstItemSliceAliasMap(
	firstItem interface{},
	indexedSliceWithValues []string,
	aliasingMap map[string]byte,
) *BasicByte {
	return it.UsingTypeSliceAliasMap(
		reflect.TypeOf(firstItem).String(),
		indexedSliceWithValues,
		aliasingMap)
}

func (it newBasicByteCreator) UsingFirstItemSliceCaseOptions(
	isIncludeUppercaseLowercase bool, // lowercase, uppercase all
	firstItem interface{},
	indexedSliceWithValues []string,
) *BasicByte {
	return it.CreateUsingSlicePlusAliasMapOptions(
		isIncludeUppercaseLowercase,
		firstItem,
		indexedSliceWithValues,
		nil)
}

// UsingFirstItemSliceAllCases
//
//	Includes both cases upper, lower case unmarshalling
func (it newBasicByteCreator) UsingFirstItemSliceAllCases(
	firstItem interface{},
	indexedSliceWithValues []string,
) *BasicByte {
	return it.CreateUsingSlicePlusAliasMapOptions(
		true,
		firstItem,
		indexedSliceWithValues,
		nil)
}

func (it newBasicByteCreator) UsingTypeSliceAliasMap(
	typeName string,
	indexedSliceWithValues []string,
	aliasingMap map[string]byte,
) *BasicByte {
	min := constants.Zero
	max := len(indexedSliceWithValues) - 1

	actualValues := make([]byte, max+1)
	for i := range indexedSliceWithValues {
		actualValues[i] = byte(i)
	}

	return it.CreateUsingAliasMap(
		typeName,
		actualValues,
		indexedSliceWithValues,
		aliasingMap, // aliasing map
		byte(min),
		byte(max),
	)
}

func (it newBasicByteCreator) UsingTypeSlice(
	typeName string,
	indexedSliceWithValues []string,
) *BasicByte {
	return it.UsingTypeSliceAliasMap(
		typeName,
		indexedSliceWithValues,
		nil, // aliasingMap
	)
}

func (it newBasicByteCreator) Default(
	firstItem interface{},
	indexedSliceWithValues []string,
) *BasicByte {
	return it.UsingTypeSliceAliasMap(
		reflect.TypeOf(firstItem).String(),
		indexedSliceWithValues,
		nil, // aliasingMap
	)
}

func (it newBasicByteCreator) DefaultWithAliasMap(
	firstItem interface{},
	indexedSliceWithValues []string,
	aliasingMap map[string]byte,
) *BasicByte {
	return it.UsingTypeSliceAliasMap(
		reflect.TypeOf(firstItem).String(),
		indexedSliceWithValues[:],
		aliasingMap, // aliasingMap
	)
}

// DefaultAllCases
//
//	includes both lowercase and uppercase parsing.
func (it newBasicByteCreator) DefaultAllCases(
	firstItem interface{},
	indexedSliceWithValues []string,
) *BasicByte {
	return it.CreateUsingSlicePlusAliasMapOptions(
		true,
		firstItem,
		indexedSliceWithValues,
		nil, // aliasingMap
	)
}

// DefaultWithAliasMapAllCases
//
//	includes both lowercase and uppercase parsing.
func (it newBasicByteCreator) DefaultWithAliasMapAllCases(
	firstItem interface{},
	indexedSliceWithValues []string,
	aliasingMap map[string]byte,
) *BasicByte {
	return it.CreateUsingSlicePlusAliasMapOptions(
		true,
		firstItem,
		indexedSliceWithValues[:],
		aliasingMap, // aliasingMap
	)
}

func (it newBasicByteCreator) generateUppercaseLowercaseAliasMap(
	isIncludeUppercaseLowercase bool,
	rangesMap map[byte]string,
	aliasingMap map[string]byte,
) map[string]byte {
	if !isIncludeUppercaseLowercase {
		return aliasingMap
	}

	finalAliasMap := make(
		map[string]byte,
		len(rangesMap)*3+len(aliasingMap)*3+2)

	for keyAsByte, valueAsName := range rangesMap {
		toUpper := strings.ToUpper(valueAsName)
		toLower := strings.ToLower(valueAsName)
		finalAliasMap[toUpper] = keyAsByte
		finalAliasMap[toLower] = keyAsByte
		finalAliasMap[valueAsName] = keyAsByte
	}

	if len(aliasingMap) == 0 {
		return finalAliasMap
	}

	for keyAsName, valueAsByte := range aliasingMap {
		toUpper := strings.ToUpper(keyAsName)
		toLower := strings.ToLower(keyAsName)
		finalAliasMap[toUpper] = valueAsByte
		finalAliasMap[toLower] = valueAsByte
		finalAliasMap[keyAsName] = valueAsByte
	}

	return finalAliasMap
}

func (it newBasicByteCreator) sliceNamesToMap(
	names []string,
) map[byte]string {
	newMap := make(
		map[byte]string,
		len(names))

	for i, name := range names {
		newMap[byte(i)] = name
	}

	return newMap
}
