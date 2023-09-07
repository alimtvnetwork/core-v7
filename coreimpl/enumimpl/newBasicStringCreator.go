package enumimpl

import (
	"fmt"
	"reflect"
	"strings"
)

type newBasicStringCreator struct{}

func (it newBasicStringCreator) Create(
	typeName string,
	actualRangesNames []string,
) *BasicString {
	return it.CreateAliasMapOnly(
		typeName,
		actualRangesNames,
		nil,
	)
}

func (it newBasicStringCreator) CreateDefault(
	firstItem interface{},
	actualRangesNames []string,
) *BasicString {
	typeName := reflect.TypeOf(firstItem).String()

	return it.CreateAliasMapOnly(
		typeName,
		actualRangesNames,
		nil,
	)
}

func (it newBasicStringCreator) CreateUsingSlicePlusAliasMapOptions(
	isIncludeUppercaseLowercase bool, // lowercase, uppercase all
	firstItem interface{},
	actualRangesNames []string,
	aliasingMap map[string]string,
) *BasicString {
	finalAliasMap := it.generateUppercaseLowercaseAliasMap(
		isIncludeUppercaseLowercase,
		actualRangesNames,
		aliasingMap)

	return it.CreateAliasMapOnly(
		reflect.TypeOf(firstItem).String(),
		actualRangesNames,
		finalAliasMap,
	)
}

func (it newBasicStringCreator) CreateUsingMapPlusAliasMapOptions(
	isIncludeUppercaseLowercase bool, // lowercase, uppercase all
	firstItem interface{},
	actualRangesNames []string,
	aliasingMap map[string]string,
) *BasicString {
	finalAliasMap := it.generateUppercaseLowercaseAliasMap(
		isIncludeUppercaseLowercase,
		actualRangesNames,
		aliasingMap)

	return it.CreateAliasMapOnly(
		reflect.TypeOf(firstItem).String(),
		actualRangesNames,
		finalAliasMap,
	)
}

func (it newBasicStringCreator) UsingFirstItemSliceCaseOptions(
	isIncludeUppercaseLowercase bool, // lowercase, uppercase all
	firstItem interface{},
	indexedSliceWithValues []string,
) *BasicString {
	return it.CreateUsingMapPlusAliasMapOptions(
		isIncludeUppercaseLowercase,
		firstItem,
		indexedSliceWithValues,
		nil)
}

// UsingFirstItemSliceAllCases
//
//  Includes both cases upper, lower case unmarshalling
func (it newBasicStringCreator) UsingFirstItemSliceAllCases(
	firstItem interface{},
	indexedSliceWithValues []string,
) *BasicString {
	return it.CreateUsingMapPlusAliasMapOptions(
		true,
		firstItem,
		indexedSliceWithValues,
		nil)
}

func (it newBasicStringCreator) CreateAliasMapOnly(
	typeName string,
	actualRangesNames []string,
	aliasingMap map[string]string,
) *BasicString {
	actualNames := make([]string, len(actualRangesNames))

	min := ""
	max := ""

	index := 0
	for _, name := range actualRangesNames {
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

func (it newBasicStringCreator) CreateUsingStringersSpread(
	typeName string,
	stringerRanges ...fmt.Stringer,
) *BasicString {
	actualNames := make([]string, len(stringerRanges))
	min := ""
	max := ""

	index := 0
	for _, strigner := range stringerRanges {
		name := strigner.String()
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
		nil,
		min, max)
}

func (it newBasicStringCreator) CreateUsingNamesSpread(
	typeName string,
	stringRangesNames ...string,
) *BasicString {
	min := ""
	max := ""

	index := 0
	for _, name := range stringRangesNames {
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
		stringRangesNames,
		nil,
		min, max)
}

func (it newBasicStringCreator) CreateUsingNamesMinMax(
	typeName string,
	stringRangesNames []string,
	min, max string,
) *BasicString {
	return it.CreateUsingAliasMap(
		typeName,
		stringRangesNames,
		nil,
		min, max)
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
	nameWithIndexMap := make(
		map[string]int,
		len(stringRangesNames))
	valueToJsonDoubleQuoteStringBytesHashmap := make(
		map[string][]byte,
		len(stringRangesNames))

	for index, actualVal := range stringRangesNames {
		key := stringRangesNames[index]
		jsonName := toJsonName(key)

		nameWithIndexMap[jsonName] = index
		nameWithIndexMap[key] = index

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
		numberEnumBase:                           enumBase,
		minVal:                                   min,
		maxVal:                                   max,
		jsonDoubleQuoteNameToValueHashMap:        stringsToHashSet(stringRangesNames),
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
	}
}

func (it newBasicStringCreator) generateUppercaseLowercaseAliasMap(
	isIncludeUppercaseLowercase bool,
	names []string,
	aliasingMap map[string]string,
) map[string]string {
	if !isIncludeUppercaseLowercase {
		return aliasingMap
	}

	finalAliasMap := make(
		map[string]string,
		len(names)*3+len(aliasingMap)*3+2)

	for _, valueAsName := range names {
		toUpper := strings.ToUpper(valueAsName)
		toLower := strings.ToLower(valueAsName)
		finalAliasMap[toUpper] = valueAsName
		finalAliasMap[toLower] = valueAsName
		finalAliasMap[valueAsName] = valueAsName
	}

	if len(aliasingMap) == 0 {
		return finalAliasMap
	}

	for keyAsName, valueAsActualName := range aliasingMap {
		toUpper := strings.ToUpper(keyAsName)
		toLower := strings.ToLower(keyAsName)
		finalAliasMap[toUpper] = valueAsActualName
		finalAliasMap[toLower] = valueAsActualName
		finalAliasMap[keyAsName] = valueAsActualName
	}

	return finalAliasMap
}

func (it newBasicStringCreator) sliceNamesToMap(
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
