package mapdiffinternal

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type MapStringAnyDiff map[string]interface{}

func (it *MapStringAnyDiff) Length() int {
	if it == nil {
		return 0
	}

	return len(*it)
}

func (it MapStringAnyDiff) IsEmpty() bool {
	return it.Length() == 0
}

func (it MapStringAnyDiff) HasAnyItem() bool {
	return it.Length() > 0
}

func (it MapStringAnyDiff) LastIndex() int {
	return it.Length() - 1
}

func (it MapStringAnyDiff) AllKeysSorted() []string {
	if it.IsEmpty() {
		return []string{}
	}

	allKeys := make(
		[]string,
		it.Length())

	index := 0
	for key := range it {
		allKeys[index] = key
		index++
	}

	sort.Strings(allKeys)

	return allKeys
}

func (it MapStringAnyDiff) Raw() map[string]interface{} {
	if it == nil {
		return map[string]interface{}{}
	}

	return it
}

func (it *MapStringAnyDiff) HasAnyChanges(
	isRegardlessType bool,
	rightMap map[string]interface{},
) bool {
	return !it.IsRawEqual(
		isRegardlessType,
		rightMap)
}

func (it *MapStringAnyDiff) IsRawEqual(
	isRegardlessType bool,
	rightMap map[string]interface{},
) bool {
	if it == nil && rightMap == nil {
		return true
	}

	if it == nil || rightMap == nil {
		return false
	}

	if it.Length() != len(rightMap) {
		return false
	}

	for key, leftValInf := range *it {
		rightValInf, has := rightMap[key]

		if !has {
			return false
		}

		if it.isNotEqual(
			isRegardlessType,
			leftValInf,
			rightValInf) {
			return false
		}
	}

	return true
}

func (it *MapStringAnyDiff) HashmapDiffUsingRaw(
	isRegardlessType bool,
	rightMap map[string]interface{},
) MapStringAnyDiff {
	diffMap := it.DiffRaw(
		isRegardlessType,
		rightMap)

	if len(diffMap) == 0 {
		return map[string]interface{}{}
	}

	return diffMap
}

func (it *MapStringAnyDiff) DiffRaw(
	isRegardlessType bool,
	rightMap map[string]interface{},
) map[string]interface{} {
	if it == nil && rightMap == nil {
		return map[string]interface{}{}
	}

	if it == nil && rightMap != nil {
		return rightMap
	}

	if it != nil && rightMap == nil {
		return *it
	}

	length := it.Length() / 3
	diffMap := make(
		map[string]interface{},
		length)

	for key, leftValInf := range *it {
		rightValInf, has := rightMap[key]

		if !has {
			diffMap[key] = leftValInf

			continue
		}

		if it.isNotEqual(
			isRegardlessType,
			leftValInf,
			rightValInf) {
			diffMap[key] = leftValInf
		}
	}

	if len(diffMap) == 0 && it.Length() == len(rightMap) {
		return diffMap
	}

	leftMap := *it
	for rightKey, rightAnyVal := range rightMap {
		_, hasDiff := diffMap[rightKey]

		if hasDiff {
			// already added

			continue
		}

		leftVal, has := leftMap[rightKey]

		if !has {
			diffMap[rightKey] = rightAnyVal

			continue
		}

		if it.isNotEqual(
			isRegardlessType,
			rightAnyVal,
			leftVal) {
			diffMap[rightKey] = rightAnyVal
		}
	}

	return diffMap
}

func (it *MapStringAnyDiff) DiffJsonMessage(
	isRegardlessType bool,
	rightMap map[string]interface{},
) string {
	diffMap := it.HashmapDiffUsingRaw(
		isRegardlessType, rightMap)

	if diffMap.Length() == 0 {
		return ""
	}

	slice := it.ToStringsSliceOfDiffMap(diffMap)
	compiledString := strings.Join(
		slice,
		constants.CommaUnixNewLine)

	return fmt.Sprintf(
		curlyWrapFormat,
		compiledString)
}

func (it *MapStringAnyDiff) ToStringsSliceOfDiffMap(
	diffMap map[string]interface{},
) (diffSlice []string) {
	allKeys := MapStringAnyDiff(diffMap).AllKeysSorted()
	slice := make([]string, len(diffMap))

	for index, key := range allKeys {
		val := diffMap[key]
		if isStringType(val) {
			slice[index] = fmt.Sprintf(
				constants.KeyValQuotationWrapJsonFormat,
				key,
				val)

			continue
		}

		// not string
		slice[index] = fmt.Sprintf(
			constants.KeyStringValAnyWrapJsonFormat,
			key,
			val)
	}

	return slice
}

func (it *MapStringAnyDiff) ShouldDiffMessage(
	isRegardlessType bool,
	title string,
	rightMap map[string]interface{},
) string {
	diffMessage := it.DiffJsonMessage(
		isRegardlessType,
		rightMap)

	if diffMessage == "" {
		return ""
	}

	return fmt.Sprintf(
		diffBetweenMapShouldBeMessageFormat,
		title,
		diffMessage)
}

func (it *MapStringAnyDiff) LogShouldDiffMessage(
	isRegardlessType bool,
	title string,
	rightMap map[string]interface{},
) (diffMessage string) {
	diffMessage = it.ShouldDiffMessage(
		isRegardlessType,
		title,
		rightMap)

	if diffMessage == "" {
		return
	}

	fmt.Println(diffMessage)

	return diffMessage
}

func (it *MapStringAnyDiff) isNotEqual(
	isRegardlessType bool,
	left,
	right interface{},
) bool {
	if isRegardlessType {
		leftString := fmt.Sprintf(
			constants.SprintPropertyNameValueFormat,
			left)
		rightString := fmt.Sprintf(
			constants.SprintPropertyNameValueFormat,
			right)

		return leftString != rightString
	}

	return !reflect.DeepEqual(left, right)
}
