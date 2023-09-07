package enumimpl

import (
	"reflect"
	"strconv"

	"gitlab.com/auk-go/core/constants"
)

type newBasicUInt16Creator struct{}

func (it newBasicUInt16Creator) CreateUsingMap(
	typeName string,
	actualRangesMap map[uint16]string,
) *BasicUInt16 {
	return it.CreateUsingMapPlusAliasMap(
		typeName,
		actualRangesMap,
		nil,
	)
}

func (it newBasicUInt16Creator) CreateUsingMapPlusAliasMap(
	typeName string,
	actualRangesMap map[uint16]string,
	aliasingMap map[string]uint16,
) *BasicUInt16 {
	var min, max uint16
	actualValues := make([]uint16, len(actualRangesMap))
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

// CreateUsingAliasMap
//
// Length : must match stringRanges and actualRangesAnyType
func (it newBasicUInt16Creator) CreateUsingAliasMap(
	typeName string,
	actualValueRanges []uint16,
	stringRanges []string,
	aliasingMap map[string]uint16,
	min, max uint16,
) *BasicUInt16 {
	enumBase := newNumberEnumBase(
		typeName,
		actualValueRanges,
		stringRanges,
		min,
		max)

	jsonDoubleQuoteNameToValueHashMap := make(map[string]uint16, len(actualValueRanges))
	valueToJsonDoubleQuoteStringBytesHashmap := make(map[uint16][]byte, len(actualValueRanges))
	valueNameHashmap := make(map[uint16]string, len(actualValueRanges))

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

	return &BasicUInt16{
		numberEnumBase:                           enumBase,
		minVal:                                   min,
		maxVal:                                   max,
		jsonDoubleQuoteNameToValueHashMap:        jsonDoubleQuoteNameToValueHashMap,
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
		valueNameHashmap:                         valueNameHashmap,
	}
}

func (it newBasicUInt16Creator) UsingFirstItemSliceAliasMap(
	firstItem interface{},
	indexedSliceWithValues []string,
	aliasingMap map[string]uint16,
) *BasicUInt16 {
	return it.UsingTypeSliceAliasMap(
		reflect.TypeOf(firstItem).String(),
		indexedSliceWithValues,
		aliasingMap)
}

func (it newBasicUInt16Creator) UsingTypeSliceAliasMap(
	typeName string,
	indexedSliceWithValues []string,
	aliasingMap map[string]uint16,
) *BasicUInt16 {
	min := constants.Zero
	max := len(indexedSliceWithValues) - 1

	actualValues := make([]uint16, max+1)
	for i := range indexedSliceWithValues {
		actualValues[i] = uint16(i)
	}

	return it.CreateUsingAliasMap(
		typeName,
		actualValues,
		indexedSliceWithValues,
		aliasingMap, // aliasing map
		uint16(min),
		uint16(max),
	)
}

func (it newBasicUInt16Creator) UsingTypeSlice(
	typeName string,
	indexedSliceWithValues []string,
) *BasicUInt16 {
	return it.UsingTypeSliceAliasMap(
		typeName,
		indexedSliceWithValues,
		nil, // aliasingMap
	)
}

func (it newBasicUInt16Creator) Default(
	firstItem interface{},
	indexedSliceWithValues []string,
) *BasicUInt16 {
	return it.UsingTypeSliceAliasMap(
		reflect.TypeOf(firstItem).String(),
		indexedSliceWithValues,
		nil, // aliasingMap
	)
}

func (it newBasicUInt16Creator) DefaultWithAliasMap(
	firstItem interface{},
	indexedSliceWithValues []string,
	aliasingMap map[string]uint16,
) *BasicUInt16 {
	return it.UsingTypeSliceAliasMap(
		reflect.TypeOf(firstItem).String(),
		indexedSliceWithValues[:],
		aliasingMap, // aliasingMap
	)
}
