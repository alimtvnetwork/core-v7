package enumimpl

import (
	"reflect"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
)

type newBasicByteCreator struct{}

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
