package enumimpl

import (
	"gitlab.com/evatix-go/core/converters"
)

type newBasicStringCreator struct{}

func (it newBasicStringCreator) CreateUsingMap(
	typeName string,
	actualRangesMap map[string]string,
) *BasicString {
	return it.CreateUsingMapPlusAliasMap(
		typeName,
		actualRangesMap,
		nil,
	)
}

func (it newBasicStringCreator) CreateUsingMapPlusAliasMap(
	typeName string,
	actualRangesMap map[string]string,
	aliasingMap map[string]string,
) *BasicString {
	actualNames := make([]string, len(actualRangesMap))

	min := ""
	max := ""

	index := 0
	for _, name := range actualRangesMap {
		actualNames[index] = name

		if name > max {
			max = name
		}

		if name < min {
			min = name
		}

		index++
	}

	return it.CreateUsingAliasMap(
		typeName,
		actualNames,
		aliasingMap, // aliasing map
		min,
		max,
	)
}

// CreateUsingAliasMap
//
// Length : must match stringRanges and actualRangesAnyType
func (it newBasicStringCreator) CreateUsingAliasMap(
	typeName string,
	stringRangesNames []string,
	aliasingMap map[string]string,
	min, max string,
) *BasicString {
	enumBase := newNumberEnumBase(
		typeName,
		stringRangesNames,
		stringRangesNames,
		min,
		max)

	jsonDoubleQuoteNameToValueHashMap := make(
		map[string]string,
		len(stringRangesNames))
	valueToJsonDoubleQuoteStringBytesHashmap := make(
		map[string][]byte,
		len(stringRangesNames))

	for i, actualVal := range stringRangesNames {
		key := stringRangesNames[i]
		jsonName := toJsonName(key)
		jsonDoubleQuoteNameToValueHashMap[jsonName] = actualVal
		jsonDoubleQuoteNameToValueHashMap[key] = actualVal
		valueToJsonDoubleQuoteStringBytesHashmap[key] = []byte(jsonName)
	}

	if len(aliasingMap) > 0 {
		for aliasName, aliasValue := range aliasingMap {
			aliasJsonName := toJsonName(aliasName)
			jsonDoubleQuoteNameToValueHashMap[aliasName] = aliasValue
			jsonDoubleQuoteNameToValueHashMap[aliasJsonName] = aliasValue
		}
	}

	return &BasicString{
		numberEnumBase: enumBase,
		minVal:         min,
		maxVal:         max,
		jsonDoubleQuoteNameToValueHashMap: *converters.
			StringsToMap(&stringRangesNames),
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
	}
}
