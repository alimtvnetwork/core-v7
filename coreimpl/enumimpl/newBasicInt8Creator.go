package enumimpl

import (
	"reflect"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
)

type newBasicInt8Creator struct{}

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
	max := len(indexedSliceWithValues) -1

	actualValues := make([]int8, max)
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
