package enumimpl

import (
	"math"
	"reflect"
	"strconv"

	"gitlab.com/auk-go/core/constants"
)

type newBasicInt32Creator struct{}

func (it newBasicInt32Creator) CreateUsingMap(
	typeName string,
	actualRangesMap map[int32]string,
) *BasicInt32 {
	return it.CreateUsingMapPlusAliasMap(
		typeName,
		actualRangesMap,
		nil,
	)
}

func (it newBasicInt32Creator) CreateUsingMapPlusAliasMap(
	typeName string,
	actualRangesMap map[int32]string,
	aliasingMap map[string]int32,
) *BasicInt32 {
	var min, max int32

	max = math.MinInt32
	min = math.MaxInt32

	actualValues := make([]int32, len(actualRangesMap))
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
func (it newBasicInt32Creator) CreateUsingAliasMap(
	typeName string,
	actualValueRanges []int32,
	stringRanges []string,
	aliasingMap map[string]int32,
	min, max int32,
) *BasicInt32 {
	enumBase := newNumberEnumBase(
		typeName,
		actualValueRanges,
		stringRanges,
		min,
		max)

	jsonDoubleQuoteNameToValueHashMap := make(map[string]int32, len(actualValueRanges))
	valueToJsonDoubleQuoteStringBytesHashmap := make(map[int32][]byte, len(actualValueRanges))
	valueNameHashmap := make(map[int32]string, len(actualValueRanges))

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

	return &BasicInt32{
		numberEnumBase:                           enumBase,
		minVal:                                   min,
		maxVal:                                   max,
		jsonDoubleQuoteNameToValueHashMap:        jsonDoubleQuoteNameToValueHashMap,
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
		valueNameHashmap:                         valueNameHashmap,
	}
}

func (it newBasicInt32Creator) UsingFirstItemSliceAliasMap(
	firstItem interface{},
	indexedSliceWithValues []string,
	aliasingMap map[string]int32,
) *BasicInt32 {
	return it.UsingTypeSliceAliasMap(
		reflect.TypeOf(firstItem).String(),
		indexedSliceWithValues,
		aliasingMap)
}

func (it newBasicInt32Creator) UsingTypeSliceAliasMap(
	typeName string,
	indexedSliceWithValues []string,
	aliasingMap map[string]int32,
) *BasicInt32 {
	min := constants.Zero
	max := len(indexedSliceWithValues) - 1

	actualValues := make([]int32, max+1)
	for i := range indexedSliceWithValues {
		actualValues[i] = int32(i)
	}

	return it.CreateUsingAliasMap(
		typeName,
		actualValues,
		indexedSliceWithValues,
		aliasingMap, // aliasing map
		int32(min),
		int32(max),
	)
}

func (it newBasicInt32Creator) UsingTypeSlice(
	typeName string,
	indexedSliceWithValues []string,
) *BasicInt32 {
	return it.UsingTypeSliceAliasMap(
		typeName,
		indexedSliceWithValues,
		nil, // aliasingMap
	)
}

func (it newBasicInt32Creator) Default(
	firstItem interface{},
	indexedSliceWithValues []string,
) *BasicInt32 {
	return it.UsingTypeSliceAliasMap(
		reflect.TypeOf(firstItem).String(),
		indexedSliceWithValues,
		nil, // aliasingMap
	)
}

func (it newBasicInt32Creator) DefaultWithAliasMap(
	firstItem interface{},
	indexedSliceWithValues []string,
	aliasingMap map[string]int32,
) *BasicInt32 {
	return it.UsingTypeSliceAliasMap(
		reflect.TypeOf(firstItem).String(),
		indexedSliceWithValues[:],
		aliasingMap, // aliasingMap
	)
}
