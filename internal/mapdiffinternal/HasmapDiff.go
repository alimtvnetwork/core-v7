package mapdiffinternal

import (
	"fmt"
	"sort"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type HashmapDiff map[string]string

func (it *HashmapDiff) Length() int {
	if it == nil {
		return 0
	}

	return len(*it)
}

func (it HashmapDiff) IsEmpty() bool {
	return it.Length() == 0
}

func (it HashmapDiff) HasAnyItem() bool {
	return it.Length() > 0
}

func (it HashmapDiff) LastIndex() int {
	return it.Length() - 1
}

func (it HashmapDiff) AllKeysSorted() []string {
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

func (it *HashmapDiff) IsRawEqual(
	rightMap map[string]string,
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

	for key, leftValString := range *it {
		rightValString, has := rightMap[key]

		if !has {
			return false
		}

		if it.isNotEqual(
			leftValString,
			rightValString) {
			return false
		}
	}

	return true
}

func (it *HashmapDiff) HashmapDiffUsingRaw(
	rightMap map[string]string,
) HashmapDiff {
	diffMap := it.DiffRaw(
		rightMap)

	if len(diffMap) == 0 {
		return map[string]string{}
	}

	return diffMap
}

func (it *HashmapDiff) DiffRaw(
	rightMap map[string]string,
) map[string]string {
	if it == nil && rightMap == nil {
		return map[string]string{}
	}

	if it == nil && rightMap != nil {
		return rightMap
	}

	if it != nil && rightMap == nil {
		return *it
	}

	length := it.Length() / 3
	diffMap := make(
		map[string]string,
		length)

	for key, leftValString := range *it {
		rightValString, has := rightMap[key]

		if !has {
			diffMap[key] = leftValString

			continue
		}

		if it.isNotEqual(leftValString, rightValString) {
			diffMap[key] = leftValString
		}
	}

	if len(diffMap) == 0 && it.Length() == len(rightMap) {
		return diffMap
	}

	leftMap := *it
	for rightKey, rightStr := range rightMap {
		_, hasDiff := diffMap[rightKey]

		if hasDiff {
			// already added

			continue
		}

		leftValStr, has := leftMap[rightKey]

		if !has {
			diffMap[rightKey] = rightStr

			continue
		}

		if it.isNotEqual(rightStr, leftValStr) {
			diffMap[rightKey] = rightStr
		}
	}

	return diffMap
}

func (it *HashmapDiff) DiffJsonMessage(
	rightMap map[string]string,
) string {
	diffMap := it.HashmapDiffUsingRaw(rightMap)

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

func (it *HashmapDiff) ToStringsSliceOfDiffMap(
	diffMap map[string]string,
) (diffSlice []string) {
	allKeys := HashmapDiff(diffMap).AllKeysSorted()
	slice := make([]string, len(diffMap))

	for index, key := range allKeys {
		val := diffMap[key]

		slice[index] = fmt.Sprintf(
			constants.KeyValQuotationWrapJsonFormat,
			key,
			val)
	}

	return slice
}

func (it *HashmapDiff) ShouldDiffMessage(
	title string,
	rightMap map[string]string,
) string {
	diffMessage := it.DiffJsonMessage(rightMap)

	if diffMessage == "" {
		return ""
	}

	return fmt.Sprintf(
		diffBetweenMapShouldBeMessageFormat,
		title,
		diffMessage)
}

func (it *HashmapDiff) LogShouldDiffMessage(
	title string,
	rightMap map[string]string,
) (diffMessage string) {
	diffMessage = it.ShouldDiffMessage(
		title,
		rightMap)

	if diffMessage == "" {
		return
	}

	fmt.Println(diffMessage)

	return diffMessage
}

func (it *HashmapDiff) isNotEqual(
	left,
	right string,
) bool {
	return left != right
}
