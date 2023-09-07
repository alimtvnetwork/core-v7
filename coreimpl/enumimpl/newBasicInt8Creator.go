package enumimpl

import (
	"math"
	"reflect"
	"strconv"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type newBasicInt8Creator struct{}

func (it newBasicInt8Creator) CreateUsingMap(
	typeName string,
	actualRangesMap map[int8]string,
) *BasicInt8 {
	return it.CreateUsingMapPlusAliasMap(
		typeName,
		actualRangesMap,
		nil,
	)
}

func (it newBasicInt8Creator) CreateUsingMapPlusAliasMap(
	typeName string,
	actualRangesMap map[int8]string,
	aliasingMap map[string]int8,
) *BasicInt8 {
	var min, max int8

	max = math.MinInt8
	min = math.MaxInt8

	actualValues := make([]int8, len(actualRangesMap))
	actualNames := make([]string, len(actualRangesMap))

	index := 0
	for val, name := range actualRangesMap {
		actualValues[index] = val
		actualNames[index] = name

		if max < val {
			max = val
		}

		if min > val {
			min = val
		}

		index++
	}

	return it.CreateUsingAliasMap(
		typeName,
		actualValues,
		actualNames,
		aliasingMap, // aliasing map
		min,
		max,
	)
}

// CreateUsingAliasMap
//
// Length : must match stringRanges and actualRangesAnyType
func (it newBasicInt8Creator) CreateUsingAliasMap(
	typeName string,
	actualValueRanges []int8,
	stringRanges []string,
	aliasingMap map[string]int8,
	min, max int8,
) *BasicInt8 {
	enumBase := newNumberEnumBase(
		typeName,
		actualValueRanges,
		stringRanges,
		min,
		max)

	jsonDoubleQuoteNameToValueHashMap := make(map[string]int8, len(actualValueRanges))
	valueToJsonDoubleQuoteStringBytesHashmap := make(map[int8][]byte, len(actualValueRanges))
	valueNameHashmap := make(map[int8]string, len(actualValueRanges))

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

	return &BasicInt8{
		numberEnumBase:                           enumBase,
		minVal:                                   min,
		maxVal:                                   max,
		jsonDoubleQuoteNameToValueHashMap:        jsonDoubleQuoteNameToValueHashMap,
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
		valueNameHashmap:                         valueNameHashmap,
	}
}

func (it newBasicInt8Creator) CreateUsingSlicePlusAliasMapOptions(
	isIncludeUppercaseLowercase bool, // lowercase, uppercase all
	firstItem interface{},
	names []string,
	aliasingMap map[string]int8,
) *BasicInt8 {
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

func (it newBasicInt8Creator) CreateUsingMapPlusAliasMapOptions(
	isIncludeUppercaseLowercase bool, // lowercase, uppercase all
	firstItem interface{},
	actualRangesMap map[int8]string,
	aliasingMap map[string]int8,
) *BasicInt8 {
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

func (it newBasicInt8Creator) UsingFirstItemSliceAliasMap(
	firstItem interface{},
	indexedSliceWithValues []string,
	aliasingMap map[string]int8,
) *BasicInt8 {
	return it.UsingTypeSliceAliasMap(
		reflect.TypeOf(firstItem).String(),
		indexedSliceWithValues,
		aliasingMap)
}

func (it newBasicInt8Creator) UsingTypeSliceAliasMap(
	typeName string,
	indexedSliceWithValues []string,
	aliasingMap map[string]int8,
) *BasicInt8 {
	min := constants.Zero
	max := len(indexedSliceWithValues) - 1

	actualValues := make([]int8, max+1)
	for i := range indexedSliceWithValues {
		actualValues[i] = int8(i)
	}

	return it.CreateUsingAliasMap(
		typeName,
		actualValues,
		indexedSliceWithValues,
		aliasingMap, // aliasing map
		int8(min),
		int8(max),
	)
}

func (it newBasicInt8Creator) UsingTypeSlice(
	typeName string,
	indexedSliceWithValues []string,
) *BasicInt8 {
	return it.UsingTypeSliceAliasMap(
		typeName,
		indexedSliceWithValues,
		nil, // aliasingMap
	)
}

func (it newBasicInt8Creator) Default(
	firstItem interface{},
	indexedSliceWithValues []string,
) *BasicInt8 {
	return it.UsingTypeSliceAliasMap(
		reflect.TypeOf(firstItem).String(),
		indexedSliceWithValues,
		nil, // aliasingMap
	)
}

func (it newBasicInt8Creator) DefaultWithAliasMap(
	firstItem interface{},
	indexedSliceWithValues []string,
	aliasingMap map[string]int8,
) *BasicInt8 {
	return it.UsingTypeSliceAliasMap(
		reflect.TypeOf(firstItem).String(),
		indexedSliceWithValues[:],
		aliasingMap, // aliasingMap
	)
}

// DefaultAllCases
//
//  includes both lowercase and uppercase parsing.
func (it newBasicInt8Creator) DefaultAllCases(
	firstItem interface{},
	indexedSliceWithValues []string,
) *BasicInt8 {
	return it.CreateUsingSlicePlusAliasMapOptions(
		true,
		firstItem,
		indexedSliceWithValues,
		nil, // aliasingMap
	)
}

// DefaultWithAliasMapAllCases
//
//  includes both lowercase and uppercase parsing.
func (it newBasicInt8Creator) DefaultWithAliasMapAllCases(
	firstItem interface{},
	indexedSliceWithValues []string,
	aliasingMap map[string]int8,
) *BasicInt8 {
	return it.CreateUsingSlicePlusAliasMapOptions(
		true,
		firstItem,
		indexedSliceWithValues[:],
		aliasingMap, // aliasingMap
	)
}

func (it newBasicInt8Creator) generateUppercaseLowercaseAliasMap(
	isIncludeUppercaseLowercase bool,
	rangesMap map[int8]string,
	aliasingMap map[string]int8,
) map[string]int8 {
	if !isIncludeUppercaseLowercase {
		return aliasingMap
	}

	finalAliasMap := make(
		map[string]int8,
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

func (it newBasicInt8Creator) sliceNamesToMap(
	names []string,
) map[int8]string {
	newMap := make(
		map[int8]string,
		len(names))

	for i, name := range names {
		newMap[int8(i)] = name
	}

	return newMap
}
