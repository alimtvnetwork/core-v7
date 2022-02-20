package enumimpl

import (
	"math"
	"reflect"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
)

type newBasicInt16Creator struct{}


func (it newBasicInt16Creator) CreateUsingMap(
	typeName string,
	actualRangesMap map[int16]string,
) *BasicInt16 {
	return it.CreateUsingMapPlusAliasMap(
		typeName,
		actualRangesMap,
		nil,
	)
}

func (it newBasicInt16Creator) CreateUsingMapPlusAliasMap(
	typeName string,
	actualRangesMap map[int16]string,
	aliasingMap map[string]int16,
) *BasicInt16 {
	var min, max int16

	max = math.MinInt16
	min = math.MaxInt16

	actualValues := make([]int16, len(actualRangesMap))
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
func (it newBasicInt16Creator) CreateUsingAliasMap(
	typeName string,
	actualValueRanges []int16,
	stringRanges []string,
	aliasingMap map[string]int16,
	min, max int16,
) *BasicInt16 {
	enumBase := newNumberEnumBase(
		typeName,
		actualValueRanges,
		stringRanges,
		min,
		max)

	jsonDoubleQuoteNameToValueHashMap := make(map[string]int16, len(actualValueRanges))
	valueToJsonDoubleQuoteStringBytesHashmap := make(map[int16][]byte, len(actualValueRanges))
	valueNameHashmap := make(map[int16]string, len(actualValueRanges))

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

	return &BasicInt16{
		numberEnumBase:                           enumBase,
		minVal:                                   min,
		maxVal:                                   max,
		jsonDoubleQuoteNameToValueHashMap:        jsonDoubleQuoteNameToValueHashMap,
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
		valueNameHashmap:                         valueNameHashmap,
	}
}

func (it newBasicInt16Creator) UsingFirstItemSliceAliasMap(
	firstItem interface{},
	indexedSliceWithValues []string,
	aliasingMap map[string]int16,
) *BasicInt16 {
	return it.UsingTypeSliceAliasMap(
		reflect.TypeOf(firstItem).String(),
		indexedSliceWithValues,
		aliasingMap)
}

func (it newBasicInt16Creator) UsingTypeSliceAliasMap(
	typeName string,
	indexedSliceWithValues []string,
	aliasingMap map[string]int16,
) *BasicInt16 {
	min := constants.Zero
	max := len(indexedSliceWithValues) - 1

	actualValues := make([]int16, max+1)
	for i := range indexedSliceWithValues {
		actualValues[i] = int16(i)
	}

	return it.CreateUsingAliasMap(
		typeName,
		actualValues,
		indexedSliceWithValues,
		aliasingMap, // aliasing map
		int16(min),
		int16(max),
	)
}

func (it newBasicInt16Creator) UsingTypeSlice(
	typeName string,
	indexedSliceWithValues []string,
) *BasicInt16 {
	return it.UsingTypeSliceAliasMap(
		typeName,
		indexedSliceWithValues,
		nil, // aliasingMap
	)
}

func (it newBasicInt16Creator) Default(
	firstItem interface{},
	indexedSliceWithValues []string,
) *BasicInt16 {
	return it.UsingTypeSliceAliasMap(
		reflect.TypeOf(firstItem).String(),
		indexedSliceWithValues,
		nil, // aliasingMap
	)
}

func (it newBasicInt16Creator) DefaultWithAliasMap(
	firstItem interface{},
	indexedSliceWithValues []string,
	aliasingMap map[string]int16,
) *BasicInt16 {
	return it.UsingTypeSliceAliasMap(
		reflect.TypeOf(firstItem).String(),
		indexedSliceWithValues[:],
		aliasingMap, // aliasingMap
	)
}
