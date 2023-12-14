package enumimpl

import (
	"strconv"
	"strings"
)

func toStringsSliceOfDiffMap(diffMap DynamicMap) (diffSlice []string) {
	isString := diffMap.IsValueString()
	isNumber := !isString
	slice := make([]string, diffMap.Length())
	spaceIndent := strings.Repeat(" ", 2)

	if isNumber {
		sortedKeyValueIntegers := diffMap.SortedKeyValues()
		for index, keyValInteger := range sortedKeyValueIntegers {
			valueString := strconv.Itoa(keyValInteger.ValueInteger)
			line := keyValInteger.WrapKey() + ":" + valueString
			slice[index] = spaceIndent + line
		}

		return slice
	}

	sortedKeysAnyValues := diffMap.SortedKeyAnyValues()

	for index, anyKeyVal := range sortedKeysAnyValues {
		line := anyKeyVal.WrapKey() + ":" +
			anyKeyVal.WrapValue()
		slice[index] = spaceIndent + line
	}

	return slice
}
